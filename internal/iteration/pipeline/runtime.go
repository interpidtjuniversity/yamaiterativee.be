package pipeline

import (
	"container/list"
	"fmt"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/stage"
	"yama.io/yamaIterativeE/internal/iteration/step"
	"yama.io/yamaIterativeE/internal/util/guc"
)

type RuntimePipelineStatus int

const (
	Init RuntimePipelineStatus = iota
	Running
	Finish
	Canceled
)

var e Executor

func init() {
	e := Executor{}
	e.Init(guc.Config{
		MaxWorkers: 10,
	})
}

/** business mapping from db.IterationAction to RuntimePipeline */
type RuntimePipeline struct {
	ID          int64
	PipeLineId  int64
	NodeNum     int
	StagesIndex []int
	StageDAG    [][]int64
	StageLayout [][]int64
	Buckets     []*stage.RuntimeStage
	Status      RuntimePipelineStatus
	ExecPath    string
	Channel     chan {}interface
}

// before this, insure already have a IterationAction
func FromIterationAction(action db.IterationAction, pipeline db.Pipeline) *RuntimePipeline {
	r := &RuntimePipeline{ID: action.ID, PipeLineId: action.PipeLineId, StageDAG: pipeline.StageDAG, StageLayout: pipeline.StageLayout, NodeNum: len(pipeline.StageDAG)}
	r.Buckets = make([]*stage.RuntimeStage, GetDAGMaxParallel(pipeline.StageDAG))
	r.StagesIndex = make([]int, r.NodeNum)
	for i := 0; i < r.NodeNum; i++ {
		// let 0,1,2,3,4,5,6,7....
		r.StagesIndex[i] = i
	}

	r.BuildRuntimeState()
	return r
}

func (rp *RuntimePipeline) BuildRuntimeState() {
	m := make(map[int]int)
	height := len(rp.StageDAG)
	width := len(rp.StageDAG[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rp.StageDAG[i][j] == 1 {
				m[j]++
			}
		}
	}
	for _, v := range rp.StagesIndex {
		if m[v] == 0 {
			// 1.query stage
			// 2.pre build stage_exec(RuntimeStage)
			// 3.store the RuntimeStage in a empty slot
			stageId := dagMapLayout(rp.StageLayout, v+1)
			s, _ := db.GetStageById(stageId)
			runtimeStage := stage.FromStage(rp, s)
			for i := 0; i<len(rp.Buckets); i++{
				if rp.Buckets[i] == nil {
					rp.Buckets[i] = runtimeStage
				}
			}

		}
	}
}

func GetDAGMaxParallel(dag [][]int64) (res int) {
	m := make(map[int]int)
	var queue []int
	height := len(dag)
	width := len(dag[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if dag[i][j] == 1 {
				m[j]++
			}
		}
	}
	var update = func(pattern string) {
		switch pattern {
		case "init":
			for i := 0; i < height; i++ {
				if _, ok := m[i]; !ok {
					queue = append(queue, i)
				}
			}
			break

		case "update":
			for i := 0; i < height; i++ {
				if _, ok := m[i]; ok&&m[i]==0 {
					queue = append(queue, i)
					delete(m, i)
				}
			}
		}
	}
	update("init")

	var max = func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	length := len(queue)
	for length > 0 {
		res = max(res, length)
		for _, v := range queue {
			for node, edge := range dag[v] {
				if edge == 1 {
					m[node]--
				}
			}
		}
		queue = queue[length:]
		update("update")
		length = len(queue)
	}

	return
}


type Executor struct {
	guc.BaseTaskPool
	actions *list.List
}

func (e *Executor) Init(config guc.Config) {
	e.BaseTaskPool.Init(config)
	e.actions = list.New()
}

func (e *Executor) Reg(bean interface{}) error  {
	runtime,ok := bean.(RuntimePipeline)
	if ok{
		e.Lock.Lock()
		defer e.Lock.Unlock()
		runtime.Status = Init
		e.actions.PushBack(&runtime)
		return nil
	}
	return RegError{Args: map[string]interface{}{"bean": bean}}
}

func (e *Executor) UnReg(bean interface{}) error {
	action,ok := bean.(RuntimePipeline)
	if ok{
		var p *RuntimePipeline
		var element *list.Element
		for i:=e.actions.Front(); i!=nil; i=i.Next() {
			element=i
			p,_ = (i.Value).(*RuntimePipeline)
			if p.ID == action.ID {
				switch p.Status {
				case Running:
					p.Status=Canceled
					break
				default:
					return UnRegError{Args: map[string]interface{}{"bean": bean}}
				}
				break
			}
		}
		// handle the runtime
		/**
		   1. cancel the runtime
		   2. git commit refresh the runtime
		   3. runtime successful executed
		*/
		if p == nil {
			return UnRegError{Args: map[string]interface{}{"bean": bean}}
		}

		// remove
		e.Lock.Lock()
		e.actions.Remove(element)
		e.Lock.Unlock()
		// handle
		err := Handle(p)

		return err
	}
	return RegError{Args: map[string]interface{}{"bean": bean}}
}

func (e *Executor) Start() {
	e.BaseTaskPool.Start()
	go func() {
		e.Lock.Lock()
		var status RuntimePipelineStatus
		var p *RuntimePipeline
		for i := e.actions.Front(); i!=nil; i=i.Next(){
			p = (i.Value).(*RuntimePipeline)
			status = p.Status
			if status == Init {
				// schedule it!
				e.schedule(p)
				// then
				p.Status = Running
			}
		}
		e.Lock.Unlock()

	}()


}

func (e *Executor) ShutDown()  {
	e.BaseTaskPool.ShutDown()
	/**
		close two goroutine and channel
	*/
}

func (e *Executor) schedule(rp *RuntimePipeline) {
	for i := 0; i < len(rp.Buckets); i++ {
		if rp.Buckets[i] != nil {
			stageExec := db.StageExec{ActId: rp.ID, StageId: rp.Buckets[i].StageId, ExecPath: rp.ExecPath}
			stageExecId, _ := db.InsertStageExec(stageExec)
			rp.Buckets[i].Id = stageExecId
			for _,s := range rp.Buckets[i].Steps {
				s.StageExecId = stageExecId
				e.sendTask(s)
			}
		}
	}
}

func (e *Executor) sendTask(s *step.RuntimeStep) {
	task := interface{}(*s).(guc.Task)
	e.GetWorker().Channel <- &task
}

/**
	1. if a Task is running just wait it down and clear it's log
    2. write db
*/
func Handle(rp *RuntimePipeline) error{
	return nil
}

type RegError struct {
	Args map[string]interface{}
}

func (err RegError) Error() string {
	return fmt.Sprintf("pipeline regist error, bean is not type of IterationAction: %v", err.Args)
}

type UnRegError struct {
	Args map[string]interface{}
}

func (err UnRegError) Error() string {
	return fmt.Sprintf("pipeline unRegist error, bean: %v", err.Args)
}

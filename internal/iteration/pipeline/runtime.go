package pipeline

import (
	"bytes"
	"container/list"
	"fmt"
	"time"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/stage"
	"yama.io/yamaIterativeE/internal/iteration/step"
	"yama.io/yamaIterativeE/internal/util/guc"
)

type RuntimePipelineState int

const (
	// need start
	Init RuntimePipelineState = iota
	// need advance
	Running

	Finish
	Canceled
	Error
	Unknown
)
func (rps RuntimePipelineState)ToString() string {
	switch rps {
	case Init:
		return "Init"
	case Running:
		return "Running"
	case Finish:
		return "Finish"
	case Canceled:
		return "Canceled"
	case Error:
		return "Error"
	default:
		return "Unknown"
	}
	return ""
}

func (rps RuntimePipelineState) FromString(s string) RuntimePipelineState {
	switch s {
	case "Init":
		return Init
	case "Running":
		return Running
	case "Finish":
		return Finish
	case "Canceled":
		return Canceled
	case "Error":
		return Error
	default:
		return Unknown
	}
}

var e *Executor

func init() {
	e = &Executor{}
	e.Init(guc.Config{
		MaxWorkers: 10,
	})
	e.Start()
}

/** business mapping from db.IterationAction to RuntimePipeline */
type RuntimePipeline struct {
	ID          int64
	PipeLineId  int64
	NodeNum     int
	StagesIndex []int                         //changeable
	StageDAG    [][]int64                     //changeable
	StageLayout [][]int64
	Buckets     []*stage.RuntimeStage //changeablewe
	Status      RuntimePipelineState  //changeable
	ExecPath    string
	Success     chan step.Message
	Failure     chan step.Message
	//if buckets is build up by BuildRuntimeStage
	IsBuildUp   bool
	// pipeline args
	Args        map[string]interface{}
}

// before this, insure already have a IterationAction
func FromIterationAction(action db.IterationAction, pipeline db.Pipeline) *RuntimePipeline {
	r := &RuntimePipeline{ID: action.Id, PipeLineId: action.PipeLineId, StageDAG: pipeline.StageDAG, StageLayout: pipeline.StageLayout, NodeNum: len(pipeline.StageDAG), ExecPath: action.ExecPath}
	//r.Buckets = make([]*stage.RuntimeStage, GetDAGMaxParallel(pipeline.StageDAG))
	r.Buckets = make([]*stage.RuntimeStage, len(pipeline.StageDAG))
	r.Success = e.Success
	r.Failure = e.Failure
	r.StagesIndex = make([]int, r.NodeNum)
	for i := 0; i < r.NodeNum; i++ {
		// let 0,1,2,3,4,5,6,7....
		r.StagesIndex[i] = i
	}

	r.BuildRuntimeStage()
	return r
}

func FromStage(pipeline *RuntimePipeline, s *db.Stage, stageIndex int) *stage.RuntimeStage {
	rs := &stage.RuntimeStage{IterationActionId: pipeline.ID, StageId: s.ID, StageIndex: stageIndex, PipelineId: pipeline.PipeLineId, TaskNum: len(s.Steps), Success: pipeline.Success, Failure: pipeline.Failure,ExecPath: pipeline.ExecPath, State: stage.Init}
	steps, _ := db.BranchQueryStepsByIds(s.Steps)
	var runtimeSteps []*step.RuntimeStep
	for _,s := range steps {
		runtimeSteps = append(runtimeSteps, FromStep(rs, s))
	}
	rs.Steps = runtimeSteps
	return rs
}

func FromStep(se *stage.RuntimeStage, sp *db.Step) *step.RuntimeStep {
	rs := &step.RuntimeStep{
		StageExecId: se.Id,
		IterationActionId: se.IterationActionId,
		StepId: sp.ID,
		StageId: se.StageId,
		PipelineId: se.PipelineId,
		StageIndex: se.StageIndex,
		Type: sp.Type,
		ExecPath: se.ExecPath, LogPath: step.FormatLogPath(), Command: sp.Command, Args: sp.Args, SuccessChannel: se.Success, FailureChannel: se.Failure, State: step.Init}
	return rs
}

func (rp *RuntimePipeline) BuildRuntimeStage() {
	if !rp.IsBuildUp {
		for _,v := range rp.StagesIndex {
			stageId := dagMapLayout(rp.StageLayout, v+1)
			s, _:= db.GetStageById(stageId)
			runtimeStage := FromStage(rp, s, v)
			rp.Buckets[runtimeStage.StageIndex] = runtimeStage
		}
		rp.IsBuildUp = true
	}
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
		if m[v] == 0 && rp.Buckets[v].State == stage.Init{
			rp.Buckets[v].State = stage.Ready
		}
	}
}

var InDegree = func(dag [][]int64) (map[int]int, int, int) {
	m := make(map[int]int)
	height := len(dag)
	width := len(dag[0])

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if dag[i][j] == 1 {
				m[j]++
			}
		}
	}
	return m, height, width
}

func GetDAGMaxParallel(dag [][]int64) (res int) {
	m, height, _ := InDegree(dag)

	var queue []int
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
	Success chan step.Message
	Failure chan step.Message
}

func (e *Executor) Init(config guc.Config) {
	e.BaseTaskPool.Init(config)
	e.actions = list.New()
	e.Success = make(chan step.Message)
	e.Failure = make(chan step.Message)
}

func (e *Executor) Reg(bean interface{}) error  {
	runtime,ok := bean.(*RuntimePipeline)
	if ok{
		e.Lock.Lock()
		runtime.Status = Init
		e.actions.PushBack(runtime)
		e.Lock.Unlock()
		return nil
	}
	return RegError{Args: map[string]interface{}{"bean": bean}}
}

func (e *Executor) UnReg(bean interface{}) error {
	action,ok := bean.(*RuntimePipeline)
	if ok{
		var p *RuntimePipeline
		var element *list.Element
		for i:=e.actions.Front(); i!=nil; i=i.Next() {
			element=i
			p,_ = (i.Value).(*RuntimePipeline)
			if p.ID == action.ID {
				switch p.Status {
				case Init:
					p.Status=Canceled
					break
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
	// scan goroutine
	go func() {
		for {
			e.Lock.Lock()
			var status RuntimePipelineState
			var p *RuntimePipeline
			var removeList []*list.Element
			for i := e.actions.Front(); i != nil; i = i.Next() {
				p = (i.Value).(*RuntimePipeline)
				status = p.Status
				if status == Init {
					// schedule it!
					e.schedule(p)
					// then
					p.Status = Running
					Handle(p)
				} else if status == Finish {
					removeList = append(removeList, i)
					Handle(p)
				}
			}
			// remove
			for i := 0; i < len(removeList); i++ {
				e.actions.Remove(removeList[i])
			}
			e.Lock.Unlock()
			time.Sleep(time.Duration(5)*time.Second)
		}
	}()

	// monitor goroutine
	go func() {
		for {
			select {
			case message, ok := <-e.Success:
				if !ok {
					return
				}
				go func() {
					e.Lock.Lock()
					for i:=e.actions.Front(); i!=nil; i=i.Next() {
						pipeline := (i.Value).(*RuntimePipeline)

						if pipeline.ID == message.IterationActionId {
							stageRuntime := pipeline.Buckets[message.StageIndex]
							stageRuntime.Lock.Lock()
							stageRuntime.TaskNum--
							if stageRuntime.TaskNum == 0 {
								// update runtimeStage state
								pipeline.Buckets[message.StageIndex].State = stage.Finish
								// delete node index
								for k,v := range pipeline.StagesIndex {
									if v == message.StageIndex {
										pipeline.StagesIndex = append(pipeline.StagesIndex[0:k], pipeline.StagesIndex[k+1:]...)
										break
									}
								}
								// delete edges
								for i:=0; i<len(pipeline.StageDAG[message.StageIndex]); i++ {
									pipeline.StageDAG[message.StageIndex][i] = 0
								}
								// judge if finished
								if len(pipeline.StagesIndex) == 0 {
									pipeline.Status = Finish
									continue
								}
								pipeline.BuildRuntimeStage()
								e.schedule(pipeline)
							}
							stageRuntime.Lock.Unlock()
						}
					}
					e.Lock.Unlock()
				}()
			}
		}
	}()

	go func() {
		for  {
			select {
			case message, ok := <- e.Failure:
				if !ok{
					return
				}
				go func() {
					e.Lock.Lock()
					for i:=e.actions.Front(); i!=nil; i=i.Next() {
						pipeline := (i.Value).(*RuntimePipeline)

						if pipeline.ID == message.IterationActionId {
							//1. update RuntimeStage state
							stageRuntime := pipeline.Buckets[message.StageIndex]
							stageRuntime.Lock.Lock()
							stageRuntime.State = stage.Failure
							stageRuntime.Lock.Unlock()
							//2. judge if RuntimePipeline needs failure
							m,_,_ := InDegree(pipeline.StageDAG)
							var inDegrees []byte
							var failures []byte
							for _,v := range pipeline.StagesIndex {
								if m[v] == 0 {
									inDegrees = append(inDegrees, byte(v))
								}
							}
							for i:=0; i<len(pipeline.Buckets); i++ {
								if pipeline.Buckets[i].State == stage.Failure {
									failures = append(failures, byte(pipeline.Buckets[i].StageIndex))
								}
							}
							if bytes.Equal(inDegrees, failures) {
								pipeline.Status = Error
							}
							//3. TODO

						}
					}
					e.Lock.Unlock()
				}()
			}
		}
	}()


}

func (e *Executor) ShutDown()  {
	e.BaseTaskPool.ShutDown()
	/**
		close two goroutine and channel
	*/
}

func (e *Executor) schedule(rp *RuntimePipeline) {
	if rp.Status == Canceled || rp.Status == Error {
		return
	}
	for i := 0; i < len(rp.Buckets); i++ {
		//stage is ready to execute
		if rp.Buckets[i].State == stage.Ready {
			stageExec := db.StageExec{ActId: rp.ID, StageId: rp.Buckets[i].StageId, ExecPath: rp.ExecPath, State: stage.Running.ToString()}
			_, _ = db.InsertStageExec(&stageExec)
			rp.Buckets[i].Id = stageExec.Id
			rp.Buckets[i].State = stage.Running
			for _,s := range rp.Buckets[i].Steps {
				s.StageExecId = stageExec.Id
				e.sendTask(s)
			}
		}
	}
}

func (e *Executor) sendTask(s *step.RuntimeStep) {
	task := (interface{}(s)).(guc.Task)
	e.GetWorker().Channel <- &task
}

/**
	1. update RuntimePipeline state and RuntimeStage state
    2. RuntimeStep state will be update earliest
*/
func Handle(rp *RuntimePipeline) error{
	db.UpdateIterationAction(rp.ID, rp.Status.ToString())
	for i := 0; i < len(rp.Buckets); i++ {
		if rp.Buckets[i].State.CanUpdate() {
			db.UpdateStageExecState(rp.Buckets[i].Id, rp.Buckets[i].State.ToString())
		}
	}
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

package pipeline

import (
	"container/list"
	"fmt"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/iteration/stage"
	"yama.io/yamaIterativeE/internal/util/guc"
)

type RuntimePipelineStatus int

const (
	Init RuntimePipelineStatus = iota
	Running
	Finish
	Canceled
)

/** business mapping from db.IterationAction to RuntimePipeline */
type RuntimePipeline struct {
	ID         int64
	PipeLineId int64
	Stages     []int64
	StageDAG   [][]int64
	Buckets    []*stage.RuntimeStage
	Status     RuntimePipelineStatus
}

func FromIterationAction(action db.IterationAction, pipeline db.Pipeline) RuntimePipeline {
	r := &RuntimePipeline{ID: action.ID, PipeLineId: action.PipeLineId, Stages: pipeline.Stages, StageDAG: pipeline.StageDAG}
	r.Buckets = make([]*stage.RuntimeStage, GetDAGMaxParallel(pipeline.StageDAG))
	return *r
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
	guc.TaskPool
	actions list.List
	Channel chan *RuntimePipeline
}

func (e *Executor) Reg(bean interface{}) error  {
	runtime,ok := bean.(RuntimePipeline)
	if ok{
		e.Lock.Lock()
		defer e.Lock.Unlock()
		e.actions.PushBack(&runtime)
		return nil
	}
	return RegError{Args: map[string]interface{}{"bean": bean}}
}

func (e *Executor) UnReg(bean interface{}, origin, target RuntimePipelineStatus) error {
	action,ok := bean.(RuntimePipeline)
	if ok{
		var p *RuntimePipeline
		var element *list.Element
		for i:=e.actions.Front(); i!=nil; i=i.Next() {
			element=i
			p,_ = (i.Value).(*RuntimePipeline)
			if p.ID == action.ID && p.Status == origin{
				// update
				p.Status = target
				break
			}
		}
		// handle the runtime
		/**
		   1. cancel the runtime
		   2. git commit refresh the runtime
		   3. runtime successful executed
		*/
		if p == nil || p.Status != target{
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
	e.TaskPool.Start()
	go func() {
		e.Lock.Lock()
		var status RuntimePipelineStatus
		var p *RuntimePipeline
		for i := e.actions.Front(); i!=nil; i=i.Next(){
			p = (i.Value).(*RuntimePipeline)
			status = p.Status
			if status == Init {
				// schedule it!

				// then
				p.Status = Running
			}
		}
		e.Lock.Unlock()

	}()


}

func (e *Executor) ShutDown()  {
	e.TaskPool.ShutDown()
	/**
		close two goroutine and channel
	*/
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

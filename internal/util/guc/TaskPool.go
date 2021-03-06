package guc

import (
	"math"
	"sync"
)

type TaskPool interface {
	//reg a task and wait is execute
	Reg(bean interface{}) error
	//unReg a task and cancel it's executing if in running, if init just remove if
	UnReg(actionId int64) error
	ShutDown()
	Start()
}

type Config struct {
	MaxWorkers int
}

type Worker struct {
	Channel chan *Task
	Close   chan bool
	Load    int
}

type BaseTaskPool struct {
	Workers []*Worker
	Lock    sync.Mutex
}

func (tp *BaseTaskPool) Init(config Config) {
	tp.Workers = make([]*Worker, config.MaxWorkers)
	for i := 0; i < config.MaxWorkers; i++ {
		tp.Workers[i] = &Worker{Channel: make(chan *Task), Close:make(chan bool)}
	}
}

func (tp *BaseTaskPool) GetWorker() *Worker{
	min := math.MaxInt32
	var index int
	for i := 0; i < len(tp.Workers); i++ {
		nums := tp.Workers[i].Load
		if nums < min {
			min = nums
			index = i
		}
	}
	return tp.Workers[index]
}


func (tp *BaseTaskPool) ShutDown() {
	for i := 0; i < len(tp.Workers); i++ {
		tp.Workers[i].Close <- true
		close(tp.Workers[i].Channel)
		close(tp.Workers[i].Close)
	}
}

func (tp *BaseTaskPool) Start() {
	for i := 0; i < len(tp.Workers); i++ {
		go func(i int) {
			worker := tp.Workers[i]
			for  {
				select {
				case task,ok := <-worker.Channel:
					if !ok{
						return
					}
					worker.Load++
					result, err := (*task).Run()
					if err == nil{
						(*task).Success(result)
					} else{
						(*task).Failure()
					}
					// if is git commit or cancel pipeline runtime
					if (*task).IsCancel() {
						(*task).Cancel()
					}
				case <-worker.Close:
					return
				}
			}
			
		}(i)
	}
}

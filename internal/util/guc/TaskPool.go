package guc

import (
	"math"
	"sync"
)

type TaskPool interface {
	Reg(bean interface{}) error
	UnReg(bean interface{}) error
	ShutDown()
	Start()
}

type Config struct {
	MaxWorkers int
}

type Worker struct {
	Channel chan *Task
	Close   chan bool
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
		nums := len(tp.Workers[i].Channel)
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

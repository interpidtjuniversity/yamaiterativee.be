package guc

import "sync"

type taskPool interface {
	Reg(bean interface{}) error
	UnReg(bean interface{}) error
	ShutDown()
	Start()
}

type Config struct {
	maxWorkers int
}

type Worker struct {
	Channel chan *Task
	Close   chan bool
}

type TaskPool struct {
	Workers []*Worker
	Lock    sync.Mutex
}

func (tp *TaskPool) Init(config Config) {
	tp.Workers = make([]*Worker, config.maxWorkers)
	for i := 0; i < config.maxWorkers; i++ {
		tp.Workers[i].Channel = make(chan *Task)
	}
}


func (tp *TaskPool) ShutDown() {
	for i := 0; i < len(tp.Workers); i++ {
		tp.Workers[i].Close <- true
		close(tp.Workers[i].Channel)
		close(tp.Workers[i].Close)
	}
}

func (tp *TaskPool) Start() {
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
						(*task).Fail()
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

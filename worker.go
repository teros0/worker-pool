package main

type Worker struct {
	ID           int
	WorkStream   chan WorkRequest
	WorkerStream chan chan WorkRequest
	Done         chan struct{}
	Func         func(...interface{})
}

func NewWorker(d chan struct{}, ws chan chan WorkRequest, id int, f func(...interface{})) *Worker {
	worker := &Worker{
		ID:           id,
		WorkStream:   make(chan WorkRequest),
		WorkerStream: ws,
		Done:         d,
		Func:         f,
	}
	return worker
}


func (w *Worker) Start() {
	go func(){
		for {
			w.WorkerStream <- w.WorkStream
			select {
			case work := <-w.WorkStream:
				w.Func(work.Args...)
			}
			case <- done {
				return
			}
		}
	}()
}
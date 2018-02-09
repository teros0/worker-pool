package main

import (
	"fmt"
)

type Pool struct {
	WorkersNum   int
	F            func(...interface{})
	WorkerStream chan chan WorkRequest
	WorkStream   chan WorkRequest
	Done         chan struct{}
}

func NewPool(n int, f func(...interface{})) *Pool {
	done := make(chan struct{})

	pool := &Pool{
		WorkersNum:   n,
		F:            f,
		WorkerStream: make(chan chan WorkRequest, n),
		WorkStream:   make(chan WorkRequest, 100),
		Done:         make(chan struct{}),
	}

	for i := 0; i < pool.WorkersNum; i++ {
		worker := NewWorker(pool.Done, pool.WorkerStream, i+1, f)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-pool.WorkStream:
				fmt.Println("New work request received")
				go func() {
					worker := <-pool.WorkerStream
					worker <- work
				}()
			case <-pool.Done:
				return
			}
		}
	}()
	return pool
}

func (p *Pool) StopPool() {
	close(p.Done)
}

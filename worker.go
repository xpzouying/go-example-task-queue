package main

import "log"

type Worker interface {
	Run(a, b int) int
}

type WorkerPool struct {
	TaskQueue *TaskQueue
}

func (pool *WorkerPool) AddWorker(w Worker) {
	go func() {
		for elem := range pool.TaskQueue.queue {
			args := elem.Args
			a := args.A
			b := args.B

			res := w.Run(a, b)
			elem.Result <- TaskResult{Result: res}
			close(elem.Result)
		}
	}()
}

type WorkerAdd struct {
	name string
}

func (w *WorkerAdd) Run(a, b int) int {
	res := a + b

	log.Printf("DEBUG: worker=%s result=%d", w.name, res)
	return res
}

type WorkerSub struct {
	name string
}

func (w *WorkerSub) Run(a, b int) int {
	res := a - b

	log.Printf("DEBUG: worker=%s result=%d", w.name, res)
	return res
}

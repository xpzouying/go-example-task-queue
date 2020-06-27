package main

import (
	"fmt"
	"log"
)

func newAddWorkerPool() *WorkerPool {
	pool := &WorkerPool{TaskQueue: NewTaskQueue(TaskTypeAdd)}

	for i := 0; i < 2; i++ {
		w := &WorkerAdd{name: fmt.Sprintf("add-worker-%d", i)}
		pool.AddWorker(w)
	}

	return pool
}

func newSubWorkerPool() *WorkerPool {
	pool := &WorkerPool{TaskQueue: NewTaskQueue(TaskTypeSub)}

	for i := 0; i < 1; i++ {
		w := &WorkerSub{name: fmt.Sprintf("sub-worker-%d", i)}
		pool.AddWorker(w)
	}

	return pool
}

func main() {
	s := NewServer()
	s.RegisterTask(TaskTypeAdd)
	s.RegisterTask(TaskTypeSub)

	s.RegisterWorkerPool(TaskTypeAdd, newAddWorkerPool())
	s.RegisterWorkerPool(TaskTypeSub, newSubWorkerPool())

	for i := 0; i < 10; i++ {
		args := TaskArgs{
			A: i,
			B: 100,
		}

		var taskType TaskType
		if i%2 == 0 {
			taskType = TaskTypeAdd
		} else {
			taskType = TaskTypeSub
		}

		taskRes := s.RunJob(taskType, args)
		log.Printf("result: %d", taskRes.Result)
	}
}

package main

import (
	"log"
	"sync"
)

type Server struct {
	tasks map[TaskType]*TaskQueue
	mu    sync.RWMutex
}

func NewServer() *Server {
	return &Server{tasks: make(map[TaskType]*TaskQueue)}
}

func (s *Server) RegisterTask(t TaskType) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.tasks[t]; ok {
		log.Println("register task failed. task exists")
		return false
	}

	s.tasks[t] = NewTaskQueue(t)
	return true
}

func (s *Server) RegisterWorkerPool(t TaskType, pool *WorkerPool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	tq, ok := s.tasks[t]
	if !ok {
		log.Printf("register worker pool failed. task queue not exists.")
		return false
	}

	pool.TaskQueue = tq
	return true
}

func (s *Server) RunJob(t TaskType, args TaskArgs) TaskResult {
	s.mu.RLock()
	tq := s.tasks[t]
	s.mu.RUnlock()

	elem := &TaskElement{
		Args:   args,
		Result: make(chan TaskResult),
	}
	tq.queue <- elem

	res := <-elem.Result
	return res
}

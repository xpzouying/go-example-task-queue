package main

type TaskType int

const (
	TaskTypeAdd = iota
	TaskTypeSub
)

type TaskArgs struct {
	A int
	B int
}

type TaskResult struct {
	Result int
}

type TaskElement struct {
	Args   TaskArgs
	Result chan TaskResult
}

type TaskQueue struct {
	typ   TaskType
	queue chan *TaskElement
}

func NewTaskQueue(t TaskType) *TaskQueue {
	return &TaskQueue{typ: t, queue: make(chan *TaskElement)}
}

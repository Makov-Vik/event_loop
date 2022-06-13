package engine

import (
	"sync"
)

type CommandQueue struct {
	mu sync.Mutex
	a  []Command

	notEmpty chan struct{}
	wait     bool
}

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(c Command)
}

type EventLoop struct {
	q *CommandQueue

	stopSignal chan struct{}
	stop       bool
}

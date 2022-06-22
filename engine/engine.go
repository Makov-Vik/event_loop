package engine

import (
	"sync"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(c Command)
}

type EventLoop struct {
	q *сommandQueue

	stopSignal chan struct{}
	stop       bool
}

func NewStopCommand() *stopCommand {
	return &stopCommand{}
}

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
	h.(*EventLoop).stop = true
}

type сommandQueue struct {
	mu sync.Mutex
	a  []Command

	notEmpty chan struct{}
	wait     bool
}

func (l *EventLoop) Start() {
	l.q = &сommandQueue{
		notEmpty: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	l.stop = false
	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

func (l *EventLoop) Post(c Command) {
	l.q.push(c)
}


func (l *EventLoop) AwaitFinish() {
	l.Post(stopCommand{})
	<-l.stopSignal
}



func (cq *сommandQueue) push(c Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.a = append(cq.a, c)

	if cq.wait {
		cq.wait = false
		cq.notEmpty <- struct{}{}
	}
}

func (cq *сommandQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if cq.empty() {
		cq.wait = true
		cq.mu.Unlock()
		<-cq.notEmpty
		cq.mu.Lock()
	}

	res := cq.a[0]
	cq.a[0] = nil
	cq.a = cq.a[1:]
	return res
}

func (cq *сommandQueue) empty() bool {
	return len(cq.a) == 0
}
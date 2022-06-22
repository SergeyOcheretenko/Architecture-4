package eventloop

import "sync"

type EventLoop struct {
	commands  *commandsQueue
	stop      bool
	isStopped chan struct{}
}

func (eLoop *EventLoop) Start() {
	eLoop.commands = &commandsQueue{
		cond: *sync.NewCond(&sync.Mutex{}),
	}
	eLoop.isStopped = make(chan struct{})
	// listenCommands ?
}

func (eLoop *EventLoop) AwaitFinish() {
	// eLoop.Post(CommandFunc(func(handler Handler) {
	//
	// }))
	<-eLoop.isStopped
}

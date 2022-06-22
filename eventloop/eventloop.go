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
	go eLoop.listenCommands()
}

func (eLoop *EventLoop) listenCommands() {
	for !eLoop.stop || !eLoop.commands.isEmpty() {
		cmd := eLoop.commands.pull()
		cmd.Execute(eLoop)
	}
	eLoop.isStopped <- struct{}{}
}

func (eLoop *EventLoop) Post(cmd Command) {
	if eLoop.stop {
		return
	}
	eLoop.commands.push(cmd)
}

func (eLoop *EventLoop) AwaitFinish() {
	eLoop.Post(CommandFunc(func(handler Handler) {
		eLoop.stop = true
	}))
	<-eLoop.isStopped
}

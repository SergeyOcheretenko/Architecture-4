package main

import (
	"github.com/SergeyOcheretenko/Architecture-4/eventloop"
)

func main() {
	eLoop := new(eventloop.EventLoop)
	eLoop.Start()
	eLoop.Post(commands.printCmd{"Hello"})
	eLoop.Post(commands.printCmd{"World"})
	eLoop.Post(commands.printCmd{"!"})
	eLoop.AwaitFinish()
}

package main

import (
	"github.com/SergeyOcheretenko/Architecture-4/commands"
	"github.com/SergeyOcheretenko/Architecture-4/eventloop"
)

func main() {
	eLoop := new(eventloop.EventLoop)
	eLoop.Start()
	eLoop.Post(commands.PrintCmd("Hello"))
	eLoop.Post(commands.PrintCmd("World"))
	eLoop.Post(commands.PrintCmd("!"))
	eLoop.Post(&commands.PrintcCmd{Count: 5, Symbol: "m"})
	eLoop.Post(&commands.PrintcCmd{Count: 10, Symbol: "*?"})
	eLoop.AwaitFinish()
}

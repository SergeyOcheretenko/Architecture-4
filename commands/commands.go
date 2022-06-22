package commands

import (
	"fmt"
	"strings"

	"github.com/SergeyOcheretenko/Architecture-4/eventloop"
)

type PrintCmd string

func (cmd PrintCmd) Execute(h eventloop.Handler) {
	fmt.Println(string(cmd))
}

type PrintcCmd struct {
	Count  int
	Symbol string
}

func (cmd *PrintcCmd) Execute(h eventloop.Handler) {
	slice := []string{}
	for i := 0; i < cmd.Count; i++ {
		slice = append(slice, cmd.Symbol)
	}
	h.Post(PrintCmd(strings.Join(slice, "")))
}

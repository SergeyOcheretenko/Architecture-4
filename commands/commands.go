package commands

import (
	"fmt"

	"github.com/SergeyOcheretenko/Architecture-4/eventloop"
)

type printCmd struct {
	text string
}

func (cmd *printCmd) Execute(h eventloop.Handler) {
	fmt.Println(cmd.text)
}

type printcCmd struct {
	count  int
	symbol rune
}

func (cmd *printcCmd) Execute(h eventloop.Handler) {
	slice := []rune{}
	for i := 0; i < cmd.count; i++ {
		slice[i] = (cmd.symbol)
	}
	h.Post(&printCmd{text: string(slice)})
}

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/SergeyOcheretenko/Architecture-4/commands"
	"github.com/SergeyOcheretenko/Architecture-4/eventloop"
)

func parse(cmdLine string) eventloop.Command {
	var cmd eventloop.Command
	partsOfCmd := strings.Fields(cmdLine)
	cmdName := partsOfCmd[0]
	if strings.Compare(cmdName, "print") == 0 {
		arg := partsOfCmd[1]
		cmd = commands.PrintCmd{Text: arg}
	} else if strings.Compare(cmdName, "printc") == 0 {
		firstArgStr := partsOfCmd[1]
		firstArgInt, _ := strconv.Atoi(firstArgStr)
		secondArg := partsOfCmd[2]
		cmd = commands.PrintcCmd{Count: firstArgInt, Symbol: secondArg}
	}
	return cmd
}

func main() {

	eLoop := new(eventloop.EventLoop)
	eLoop.Start()
	filePath := "test.txt"
	if input, err := os.Open(filePath); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eLoop.Post(cmd)
		}
	}
	eLoop.AwaitFinish()
}

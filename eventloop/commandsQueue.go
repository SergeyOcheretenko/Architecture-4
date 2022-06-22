package eventloop

import "sync"

type commandsQueue struct {
  commands []Command
  cond     sync.Cond
}

func (q *commandsQueue) isEmpty() bool {
  q.cond.L.Lock()
  defer q.cond.L.Unlock()
  return len(q.commands) == 0
}

func (q *commandsQueue) push(cmd Command) {
  q.cond.L.Lock()
  q.commands = append(q.commands, cmd)
  q.cond.L.Unlock()
  q.cond.Broadcast()
}

func (q *commandsQueue) pull() Command {
  for len(q.commands) == 0 {
    q.cond.Wait()
  }
  q.cond.L.Lock()
  defer q.cond.L.Unlock()
  cmd := q.commands[0]
  q.commands[0] = nil
  q.commands = q.commands[1:]
  return cmd
}
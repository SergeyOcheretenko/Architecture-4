package eventloop

type Command interface {
  Execute(handler Handler)
}

type CommandFunc func(handler Handler)

func (cmdFunc CommandFunc) Execute(h Handler) {
  cmdFunc(h)
}
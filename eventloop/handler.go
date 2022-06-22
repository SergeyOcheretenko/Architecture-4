package eventloop

type Handler interface {
  Post(cmd Command)
}
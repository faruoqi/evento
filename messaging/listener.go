package messaging

type EventListener interface {
	Listen(eventNames ...string) (<-chan Event, <-chan error, error)
}

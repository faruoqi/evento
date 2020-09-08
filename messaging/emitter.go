package messaging

type EventEmitter interface {
	Emit(event Event) error
}

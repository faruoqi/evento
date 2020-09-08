package messaging

type Event interface {
	EventName() string
}

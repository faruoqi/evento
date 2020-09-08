package model

type DbEventHandler interface {
	FindAllEvents() ([]Event, error)
	FindEventByName(eventName string) (Event, error)
	FindEventByID(ID string) (Event, error)
	AddEvent(event Event) error
}

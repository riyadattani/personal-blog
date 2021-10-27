package in_mem

import (
	"io/fs"
	"personal-blog/pkg/event"
	"personal-blog/pkg/file_system"
)

type EventService interface {
	GetEvents() []event.Event
}

type EventStore struct {
	events []event.Event
}

func NewEventStore(eventsDir fs.FS) (*EventStore, error) {
	events, err1 := file_system.NewEvents(eventsDir)
	if err1 != nil {
		return nil, err1
	}

	return &EventStore{ events: events}, nil
}

func (i *EventStore) GetEvents() []event.Event {
	return i.events
}

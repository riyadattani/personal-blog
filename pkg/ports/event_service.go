package ports

import "personal-blog/pkg/event"

type EventService interface {
	GetEvents() []event.Event
}

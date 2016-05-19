package controller

import (
	"time"
)

type EventAgregator struct {
	listeners map[string][]func(EventData)
}

func NewEventAgregator() *EventAgregator {
	ea := EventAgregator{
		listeners: make(map[string][]func(EventData)),
	}
	return &ea
}

func (ea *EventAgregator) AddListener(name string, f func(EventData)) {
	ea.listeners[name] = append(ea.listeners[name], f)
}

func (ea *EventAgregator) PublishEvent(name string, eventData EventData) {
	if ea.listeners[name] != nil {
		for _, r := range ea.listeners[name] {
			r(eventData)
		}
	}
}

type EventData struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

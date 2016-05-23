package controller

import (
    "time"
)

type EventRaiser interface {
    AddListener(eventName string, f func(interface{}))
}

type EventAgregator struct {
    listeners map[string][]func(interface{})
}

func NewEventAgregator() *EventAgregator {
    ea := EventAgregator{
        listeners: make(map[string][]func(interface{})),
    }
    return &ea
}

func (ea *EventAgregator) AddListener(name string, f func(interface{})) {
    ea.listeners[name] = append(ea.listeners[name], f)
}

func (ea *EventAgregator) PublishEvent(name string, eventData interface{}) {
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

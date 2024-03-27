package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type Event interface {
	GetName() string
	GetPayload() interface{}
	SetPayload(interface{})
}

type EventHandler interface {
	Handle(event Event)
}

type EventDispatcher interface {
	Register(event Event, handler EventHandler) error
	Dispatch(event Event) error
}

type dispatcher struct {
	handlers map[string][]EventHandler
}

func NewDispatcher() EventDispatcher {
	return &dispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *dispatcher) Register(event Event, handler EventHandler) error {
	if _, ok := d.handlers[event.GetName()]; ok {
		for _, h := range d.handlers[event.GetName()] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	d.handlers[event.GetName()] = append(d.handlers[event.GetName()], handler)
	return nil
}

func (d *dispatcher) Dispatch(event Event) error {
	if handlers, ok := d.handlers[event.GetName()]; ok {
		for _, handler := range handlers {
			handler.Handle(event)
		}
	}
	return nil
}

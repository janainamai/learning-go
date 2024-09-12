package events

import (
	"sync"
	"time"
)

type (
	EventInterface interface {
		GetName() string
		GetDateTime() time.Time
		GetPayload() interface{}
	}

	EventHandlerInterface interface {
		Handle(event EventInterface, wg *sync.WaitGroup)
	}

	EventDispatcherInterface interface {
		Register(eventName string, handler EventHandlerInterface) error
		Dispacth(event EventInterface) error
		Remove(eventName string, handler EventHandlerInterface) error
		Has(eventName string, handler EventHandlerInterface) bool
		Clear() error
	}
)

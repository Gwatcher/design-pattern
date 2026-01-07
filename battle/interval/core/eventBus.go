package core

import (
	"sync"
	"time"
)

type Subscriber func(e Event)

type EventBus struct {
	mu   sync.RWMutex
	subs map[EventType][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{subs: make(map[EventType][]Subscriber)}
}

func (b *EventBus) Subscribe(t EventType, fn Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subs[t] = append(b.subs[t], fn)
}

func (b *EventBus) Publish(t EventType, msg string) {
	b.mu.RLock()
	list := append([]Subscriber(nil), b.subs[t]...)
	b.mu.RUnlock()

	e := Event{Type: t, Msg: msg, Time: time.Now()}
	for _, fn := range list {
		fn(e)
	}
}

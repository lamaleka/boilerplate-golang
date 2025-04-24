package sse

import "sync"

type SseRegistry[T any] struct {
	mu        sync.Mutex
	listeners map[string]chan T
}

func NewSseRegistry[T any]() *SseRegistry[T] {
	return &SseRegistry[T]{
		listeners: make(map[string]chan T),
	}
}

func (r *SseRegistry[T]) Unregister(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if ch, ok := r.listeners[id]; ok {
		close(ch)
		delete(r.listeners, id)
	}
}

func (r *SseRegistry[T]) Notify(id string, data T) {
	r.mu.Lock()
	ch, ok := r.listeners[id]
	r.mu.Unlock()

	if ok {
		select {
		case ch <- data:
		default:
		}
	}
}

func (r *SseRegistry[T]) Get(id string, bufferSize int) chan T {
	r.mu.Lock()
	defer r.mu.Unlock()

	if ch, exists := r.listeners[id]; exists {
		return ch
	}

	newCh := make(chan T, bufferSize)
	r.listeners[id] = newCh
	return newCh
}

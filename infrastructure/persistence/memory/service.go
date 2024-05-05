package memory

import "sync"

type Memory[K comparable, V any] struct {
	mutex sync.Mutex
	Items map[K]V
}

func NewMemory[K comparable, V any]() *Memory[K, V] {
	p := new(Memory[K, V])

	p.Items = make(map[K]V)

	return p
}

func (m *Memory[K, V]) Add(key K, value V) {
	m.mutex.Lock()

	m.Items[key] = value

	m.mutex.Unlock()
}

func (m *Memory[K, V]) Get(key K) (V, bool) {
	m.mutex.Lock()

	v, ok := m.Items[key]

	m.mutex.Unlock()

	return v, ok
}

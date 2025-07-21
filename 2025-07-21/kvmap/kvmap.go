package kvmap

type Map[K comparable, V any] struct {
	data map[K]V
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V)}
}

func (m *Map[K, V]) Set(key K, value V) {
	m.data[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	val, ok := m.data[key]
	return val, ok
}

func (m *Map[K, V]) Delete(key K) {
	delete(m.data, key)
}

func (m *Map[K, V]) Has(key K) bool {
	_, ok := m.data[key]
	return ok
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) Len() int {
	return len(m.data)
}

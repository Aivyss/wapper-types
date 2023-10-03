package types

type Map[K comparable, V any] map[K]V

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func (m *Map[K, V]) Raw() map[K]V {
	return *m
}

func (m *Map[K, V]) Set(key K, value V) {
	m.Raw()[key] = value
}

func (m *Map[K, V]) Get(key K) V {
	return m.Raw()[key]
}

func (m *Map[K, V]) Keys() List[K] {
	return Mappable[Entry[K, V], K](m.Entries()).Do(
		func(entry Entry[K, V]) K {
			return entry.Key
		},
	)
}

func (m *Map[K, V]) Values() List[V] {
	return Mappable[Entry[K, V], V](m.Entries()).Do(
		func(entry Entry[K, V]) V {
			return entry.Value
		},
	)
}

func (m *Map[K, V]) Entries() List[Entry[K, V]] {
	entries := List[Entry[K, V]]{}
	for key, value := range *m {
		entries = entries.Append(Entry[K, V]{
			Key:   key,
			Value: value,
		})
	}

	return entries
}

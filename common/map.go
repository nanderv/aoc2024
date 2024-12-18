package common

type MiMap[T any] struct {
	m map[int]T
	f func(T) int
}

func (m *MiMap[T]) Len() int {
	return len(m.m)
}
func (m *MiMap[T]) Has(a T) bool {
	i := m.f(a)
	_, ok := m.m[i]
	return ok
}

func (m *MiMap[T]) Get(a T) T {
	i := m.f(a)

	return m.m[i]
}
func (m *MiMap[T]) Set(a T) {
	i := m.f(a)
	m.m[i] = a
	return
}

func NewMiMap[T any](f func(T) int) MiMap[T] {
	return MiMap[T]{m: make(map[int]T), f: f}
}

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

func (m *MiMap[T]) MGet(a T) (T, bool) {
	i := m.f(a)
	v, ok := m.m[i]
	return v, ok
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

func (m *MiMap[T]) Keys() []T {
	keys := make([]T, m.Len())
	i := 0
	for _, v := range m.m {
		keys[i] = v
		i++
	}
	return keys
}

func NewMiMap[T any](f func(T) int) MiMap[T] {
	return MiMap[T]{m: make(map[int]T, 4000), f: f}
}

type SlowMiMap[T any] struct {
	m map[string]T
	f func(T) string
}

func (m *SlowMiMap[T]) Len() int {
	return len(m.m)
}
func (m *SlowMiMap[T]) Has(a T) bool {
	i := m.f(a)
	_, ok := m.m[i]
	return ok
}

func (m *SlowMiMap[T]) Get(a T) T {
	i := m.f(a)

	return m.m[i]
}
func (m *SlowMiMap[T]) Set(a T) {
	i := m.f(a)
	m.m[i] = a
	return
}

func (m *SlowMiMap[T]) Keys() []T {
	keys := make([]T, m.Len())
	i := 0
	for _, v := range m.m {
		keys[i] = v
		i++
	}
	return keys
}

func NewSlowMiMap[T any](f func(T) string) SlowMiMap[T] {
	return SlowMiMap[T]{m: make(map[string]T), f: f}
}

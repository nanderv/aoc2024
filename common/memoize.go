package common

func Memoize[T any, U any](f func(T) string, g func(T) U) func(T) U {
	mp := make(map[string]U)
	return func(a T) U {
		v, ok := mp[f(a)]
		if ok {
			return v
		}
		v = g(a)
		mp[f(a)] = v
		return v
	}

}

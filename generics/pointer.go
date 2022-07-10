package generics

func Pointer[T any](v T) *T {
	return &v
}

func Value[T any](p *T) T {
	if p == nil {
		return Default[T]()
	}
	return *p
}

func Default[T any]() T {
	var defaultVal T
	return defaultVal
}

package flinx

func Self[T any](t T) T {
	return t
}

func Not[T any](predicate func(T) bool) func(T) bool {
	return func(t T) bool {
		return !predicate(t)
	}
}

func ValidateQuery[T comparable](q Query[T], output []T) bool {
	next := q.Iterate()

	for _, oitem := range output {
		qitem, _ := next()

		if oitem != qitem {
			return false
		}
	}

	_, ok := next()
	_, ok2 := next()
	return !(ok || ok2)
}

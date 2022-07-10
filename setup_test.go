package flinx

type foo struct {
	f1 int
	f2 bool
	f3 string
}

func getF1(f foo) int {
	return f.f1
}

func getF2(f foo) bool {
	return f.f2
}
func getF3(f foo) string {
	return f.f3
}

func getSelf[T any](t T) T {
	return t
}

func toSlice[T any](q Query[T]) (result []T) {
	next := q.Iterate()

	for item, ok := next(); ok; item, ok = next() {
		result = append(result, item)
	}

	return
}

func validateQuery[T comparable](q Query[T], output []T) bool {
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

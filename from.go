package flinx

// Iterator is an alias for function to iterate over data.
type Iterator[T any] func() (item T, ok bool)

// Query is the type returned from query functions. It can be iterated manually
// as shown in the example.
type Query[T any] struct {
	Iterate func() Iterator[T]
}

// KeyValue is a type that is used to iterate over a map (if query is created
// from a map). This type is also used by ToMap() method to output result of a
// query into a map.
type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

func FromSlice[T any](source []T) Query[T] {

	length := len(source)

	return Query[T]{
		Iterate: func() Iterator[T] {
			index := 0

			return func() (item T, ok bool) {
				ok = index < length
				if ok {
					item = source[index]
					index++
				}

				return
			}
		},
	}
}

func FromMap[K comparable, V any](source map[K]V) Query[KeyValue[K, V]] {

	length := len(source)
	keys := make([]K, length)
	idx := 0
	for k := range source {
		keys[idx] = k
		idx++
	}
	return Query[KeyValue[K, V]]{
		Iterate: func() Iterator[KeyValue[K, V]] {
			index := 0
			return func() (item KeyValue[K, V], ok bool) {
				ok = index < length
				if ok {
					key := keys[index]
					item = KeyValue[K, V]{
						Key:   key,
						Value: source[key],
					}

					index++
				}

				return
			}
		},
	}
}

// FromChannel initializes a linq query with passed channel, linq iterates over
// channel until it is closed.
func FromChannel[T any](source <-chan T) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			return func() (item T, ok bool) {
				item, ok = <-source
				return
			}
		},
	}
}

// FromString initializes a linq query with passed string, linq iterates over
// runes of string.
func FromString(source string) Query[rune] {
	runes := []rune(source)
	length := len(source)

	return Query[rune]{
		Iterate: func() Iterator[rune] {
			index := 0

			return func() (item rune, ok bool) {
				ok = index < length
				if ok {
					item = runes[index]
					index++
				}

				return
			}
		},
	}
}

// FromIterable initializes a linq query with custom collection passed. This
// collection has to implement Iterable interface, linq iterates over items,
// that has to implement Comparable interface or be basic types.
func FromIterable[T any](sourceFunc Iterator[T]) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			return sourceFunc
		},
	}
}

// Range generates a sequence of integral numbers within a specified range.
func Range(start, count int) Query[int] {
	return Query[int]{
		Iterate: func() Iterator[int] {
			index := 0
			current := start

			return func() (item int, ok bool) {
				if index >= count {
					return 0, false
				}

				item, ok = current, true

				index++
				current++
				return
			}
		},
	}
}

// Repeat generates a sequence that contains one repeated value.
func Repeat[T any](value T, count int) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			index := 0

			return func() (item T, ok bool) {
				if index >= count {
					var r T
					return r, false
				}

				item, ok = value, true

				index++
				return
			}
		},
	}
}

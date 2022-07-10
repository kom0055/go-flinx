package flinx

// Group is a type that is used to store the result of GroupBy method.
type Group[K, V any] struct {
	Key   K
	Group []V
}

// GroupBy method groups the elements of a collection according to a specified
// key selector function and projects the elements for each group by using a
// specified function.
func GroupBy[T any, K comparable, V any](keySelector func(T) K,
	elementSelector func(T) V) func(q Query[T]) Query[Group[K, V]] {
	return func(q Query[T]) Query[Group[K, V]] {
		return Query[Group[K, V]]{
			func() Iterator[Group[K, V]] {
				next := q.Iterate()
				set := map[K][]V{}

				for item, ok := next(); ok; item, ok = next() {
					key := keySelector(item)
					set[key] = append(set[key], elementSelector(item))
				}

				length := len(set)
				idx := 0
				groups := make([]Group[K, V], length)
				for k, v := range set {
					groups[idx] = Group[K, V]{k, v}
					idx++
				}

				index := 0

				return func() (item Group[K, V], ok bool) {
					ok = index < length
					if ok {
						item = groups[index]
						index++
					}

					return
				}
			},
		}
	}

}

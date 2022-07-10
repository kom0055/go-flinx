package flinx

// Aggregate applies an accumulator function over a sequence.
//
// Aggregate method makes it simple to perform a calculation over a sequence of
// values. This method works by calling f() one time for each element in source
// except the first one. Each time f() is called, Aggregate passes both the
// element from the sequence and an aggregated value (as the first argument to
// f()). The first element of source is used as the initial aggregate value. The
// result of f() replaces the previous aggregated value.
//
// Aggregate returns the final result of f().
func Aggregate[T any](f func(T, T) T) func(q Query[T]) T {
	return func(q Query[T]) (r T) {
		next := q.Iterate()

		result, exist := next()
		if !exist {
			return
		}

		for current, ok := next(); ok; current, ok = next() {
			result = f(result, current)
		}

		return result
	}

}

// AggregateWithSeed applies an accumulator function over a sequence. The
// specified seed value is used as the initial accumulator value.
//
// Aggregate method makes it simple to perform a calculation over a sequence of
// values. This method works by calling f() one time for each element in source
// except the first one. Each time f() is called, Aggregate passes both the
// element from the sequence and an aggregated value (as the first argument to
// f()). The value of the seed parameter is used as the initial aggregate value.
// The result of f() replaces the previous aggregated value.
//
// Aggregate returns the final result of f().
func AggregateWithSeed[T any](seed T, f func(T, T) T) func(q Query[T]) (t T) {
	return func(q Query[T]) (t T) {
		next := q.Iterate()
		result := seed

		for current, ok := next(); ok; current, ok = next() {
			result = f(result, current)

		}

		return result
	}

}

// AggregateWithSeedBy applies an accumulator function over a sequence. The
// specified seed value is used as the initial accumulator value, and the
// specified function is used to select the result value.
//
// Aggregate method makes it simple to perform a calculation over a sequence of
// values. This method works by calling f() one time for each element in source.
// Each time func is called, Aggregate passes both the element from the sequence
// and an aggregated value (as the first argument to func). The value of the
// seed parameter is used as the initial aggregate value. The result of func
// replaces the previous aggregated value.
//
// The final result of func is passed to resultSelector to obtain the final
// result of Aggregate.
func AggregateWithSeedBy[T any, V any](seed T, f func(T, T) T, resultSelector func(T) V) func(q Query[T]) (v V) {

	return func(q Query[T]) (v V) {
		next := q.Iterate()
		result := seed

		for current, ok := next(); ok; current, ok = next() {
			result = f(result, current)

		}
		s := resultSelector(result)

		return s
	}
}

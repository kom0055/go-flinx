package generics

import (
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/constraints"
)

func NumericCompare[T constraints.Float | constraints.Integer](v1, v2 T) int {
	return int(int64(v1) - int64(v2))
}

func OrderedCompare[T constraints.Ordered](t1, t2 T) int {
	if t1 == t2 {
		return 0
	}

	if t1 < t2 {
		return -1
	}
	return 1
}

func BoolCompare(a, b bool) int {
	switch {
	case a == b:
		return 0
	case a:
		return 1
	default:
		return -1
	}
}

func Equal(x, y any) bool {
	return cmp.Equal(x, y)
}

func Self[T any](x T) T {
	return x
}

func Any[T any](_ T) bool {
	return true
}

func NotAny[T any](_ T) bool {
	return false
}

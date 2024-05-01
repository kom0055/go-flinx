package flinx

import (
	"reflect"

	"golang.org/x/exp/constraints"

	"github.com/kom0055/go-flinx/hashset"
)

func Self[T any](t T) T {
	return t
}

func Not[T any](predicate func(T) bool) func(T) bool {
	return func(t T) bool {
		return !predicate(t)
	}
}

func Nil(t any) bool {
	valueOf := reflect.ValueOf(t)

	k := valueOf.Kind()

	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return valueOf.IsNil()
	default:
	}
	return t == nil
}

func NonNil(t any) bool {
	return !Nil(t)
}

func Gte[T constraints.Ordered](v T) func(t T) bool {
	return func(t T) bool {
		return t >= v
	}
}

func Lte[T constraints.Ordered](v T) func(t T) bool {
	return func(t T) bool {
		return t <= v
	}
}

func Gt[T constraints.Ordered](v T) func(t T) bool {
	return Not(Lte(v))
}

func Lt[T constraints.Ordered](v T) func(t T) bool {
	return Not(Gte(v))
}

func Eq[T comparable](v T) func(t T) bool {
	return func(t T) bool {
		return t == v
	}
}

func Neq[T comparable](v T) func(t T) bool {
	return Not(Eq(v))
}

func In[T comparable](v ...T) func(t T) bool {
	return func(t T) bool {
		return hashset.NewAny(v...).Has(t)
	}
}

func NotIn[T comparable](v ...T) func(t T) bool {
	return Not(In(v...))
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

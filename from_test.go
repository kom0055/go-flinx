package flinx

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestFrom(t *testing.T) {

	{
		tests := []struct {
			input  []int
			output []int
			want   bool
		}{
			{[]int{1, 2, 3}, []int{1, 2, 3}, true},
			{[]int{1, 2, 4}, []int{1, 2, 3}, false},
			{[]int{1, 2, 3}, []int{1, 2, 3}, true},
			{[]int{1, 2, 4}, []int{1, 2, 3}, false},
		}
		for _, test := range tests {
			if q := FromSlice(test.input); ValidateQuery(q, test.output) != test.want {
				if test.want {
					t.Errorf("From(%v)=%v expected %v", test.input, ToSlice(q), test.output)
				} else {
					t.Errorf("From(%v)=%v expected not equal", test.input, test.output)
				}
			}
		}
	}

	{
		tests := []struct {
			input  string
			output []rune
			want   bool
		}{

			{"str", []rune{'s', 't', 'r'}, true},
			{"str", []rune{'s', 't', 'g'}, false},
		}
		for _, test := range tests {
			if q := FromString(test.input); ValidateQuery(q, test.output) != test.want {
				if test.want {
					t.Errorf("From(%v)=%v expected %v", test.input, ToSlice(q), test.output)
				} else {
					t.Errorf("From(%v)=%v expected not equal", test.input, test.output)
				}
			}
		}
	}

	{
		tests := []struct {
			input  map[string]bool
			output []KeyValue[string, bool]
			want   bool
		}{

			{map[string]bool{"foo": true}, []KeyValue[string, bool]{{"foo", true}}, true},
			{map[string]bool{"foo": true}, []KeyValue[string, bool]{{"foo", false}}, false},
		}
		for _, test := range tests {
			if q := FromMap(test.input); ValidateQuery(q, test.output) != test.want {
				if test.want {
					t.Errorf("From(%v)=%v expected %v", test.input, ToSlice(q), test.output)
				} else {
					t.Errorf("From(%v)=%v expected not equal", test.input, test.output)
				}
			}
		}
	}
	{
		c := make(chan int, 3)
		c <- -1
		c <- 0
		c <- 1
		close(c)

		ct := make(chan int, 3)
		ct <- -10
		ct <- 0
		ct <- 10
		close(ct)

		tests := []struct {
			input  chan int
			output []int
			want   bool
		}{

			{c, []int{-1, 0, 1}, true},
			{ct, []int{-10, 0, 10}, true},
		}

		for _, test := range tests {
			if q := FromChannel(test.input); ValidateQuery(q, test.output) != test.want {
				if test.want {
					t.Errorf("From(%v)=%v expected %v", test.input, ToSlice(q), test.output)
				} else {
					t.Errorf("From(%v)=%v expected not equal", test.input, test.output)
				}
			}
		}
	}
	{
		s := foo{f1: 1, f2: true, f3: "string"}
		tests := []struct {
			input  foo
			output []interface{}
			want   bool
		}{

			{s, []interface{}{1, true, "string"}, true},
		}

		var fooIterator = func() Iterator[any] {
			i := 0

			return func() (item any, ok bool) {
				switch i {
				case 0:
					item = s.f1
					ok = true
				case 1:
					item = s.f2
					ok = true
				case 2:
					item = s.f3
					ok = true
				default:
					ok = false
				}

				i++
				return
			}
		}
		for _, test := range tests {
			res := Results(FromIterable(fooIterator()))
			assert.DeepEqual(t, res, test.output)

		}
	}

}

func TestFromChannel(t *testing.T) {
	c := make(chan int, 3)
	c <- 10
	c <- 15
	c <- -3
	close(c)

	w := []int{10, 15, -3}

	if q := FromChannel(c); !ValidateQuery(q, w) {
		t.Errorf("FromChannel() failed expected %v", w)
	}
}

func TestFromChannelT(t *testing.T) {
	c := make(chan int, 3)
	c <- 10
	c <- 15
	c <- -3
	close(c)

	w := []int{10, 15, -3}

	if q := FromChannel(c); !ValidateQuery(q, w) {
		t.Errorf("FromChannelT() failed expected %v", w)
	}
}

func TestFromString(t *testing.T) {
	s := "string"
	w := []rune{'s', 't', 'r', 'i', 'n', 'g'}

	if q := FromString(s); !ValidateQuery(q, w) {
		t.Errorf("FromString(%v)!=%v", s, w)
	}
}

func TestFromIterable(t *testing.T) {
	s := foo{f1: 1, f2: true, f3: "string"}
	w := []interface{}{1, true, "string"}

	var fooIterator = func() Iterator[any] {
		i := 0

		return func() (item any, ok bool) {
			switch i {
			case 0:
				item = s.f1
				ok = true
			case 1:
				item = s.f2
				ok = true
			case 2:
				item = s.f3
				ok = true
			default:
				ok = false
			}

			i++
			return
		}
	}
	res := Results(FromIterable(fooIterator()))
	assert.DeepEqual(t, res, w)

}

func TestRange(t *testing.T) {
	w := []int{-2, -1, 0, 1, 2}

	if q := Range(-2, 5); !ValidateQuery(q, w) {
		t.Errorf("Range(-2, 5)=%v expected %v", ToSlice(q), w)
	}
}

func TestRepeat(t *testing.T) {
	w := []int{1, 1, 1, 1, 1}

	if q := Repeat(1, 5); !ValidateQuery(q, w) {
		t.Errorf("Repeat(1, 5)=%v expected %v", ToSlice(q), w)
	}
}

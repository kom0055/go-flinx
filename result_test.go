package flinx

import (
	"github.com/kom0055/flinx/generics"
	"gotest.tools/v3/assert"
	"math"
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {

	input := []int{2, 4, 6, 8}

	r1 := All[int](func(i int) bool {
		return i%2 == 0
	})(FromSlice[int](input))
	r2 := All[int](func(i int) bool {
		return i%2 != 0
	})(FromSlice[int](input))
	if !r1 {
		t.Errorf("From(%v).All()=%v", input, r1)
	}

	if r2 {
		t.Errorf("From(%v).All()=%v", input, r2)
	}
}

func TestAny(t *testing.T) {

	{
		tests := []struct {
			input []int
			want  bool
		}{
			{[]int{1, 2, 2, 3, 1}, true},
			{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, true},
			{[]int{}, false},
		}

		for _, test := range tests {

			if r := Any[int](FromSlice[int](test.input)); r != test.want {
				t.Errorf("From(%v).Any()=%v expected %v", test.input, r, test.want)
			}
		}
	}
	{
		tests := []struct {
			input string
			want  bool
		}{

			{"sstr", true},
		}

		for _, test := range tests {
			if r := Any[rune](FromString(test.input)); r != test.want {
				t.Errorf("From(%v).Any()=%v expected %v", test.input, r, test.want)
			}
		}
	}
}

func TestAnyWith(t *testing.T) {
	tests := []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 2, 3, 1}, false},
		{[]int{1, 1, 1, 2, 1, 2, 3, 4, 2}, true},
		{[]int{}, false},
	}

	for _, test := range tests {

		if r := AnyWith[int](func(i int) bool {
			return i == 4
		})(FromSlice[int](test.input)); r != test.want {
			t.Errorf("From(%v).Any()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestAverage(t *testing.T) {
	{
		tests := []struct {
			input []int
			want  float64
		}{
			{[]int{1, 2, 2, 3, 1}, 1.8},
		}

		for _, test := range tests {
			if r := Average[int](FromSlice[int](test.input)); r != test.want {
				t.Errorf("From(%v).Average()=%v expected %v", test.input, r, test.want)
			}
		}
	}

	{
		tests := []struct {
			input []uint
			want  float64
		}{
			{[]uint{1, 2, 5, 7, 10}, 5.},
		}

		for _, test := range tests {
			if r := Average[uint](FromSlice[uint](test.input)); r != test.want {
				t.Errorf("From(%v).Average()=%v expected %v", test.input, r, test.want)
			}
		}
	}
	{
		tests := []struct {
			input []float32
			want  float64
		}{
			{[]float32{1., 1.}, 1.},
		}

		for _, test := range tests {
			if r := Average[float32](FromSlice[float32](test.input)); r != test.want {
				t.Errorf("From(%v).Average()=%v expected %v", test.input, r, test.want)
			}
		}
	}
}

func TestAverageForNaN(t *testing.T) {

	if r := Average[int](FromSlice[int]([]int{})); !math.IsNaN(r) {
		t.Errorf("From([]int{}).Average()=%v expected %v", r, math.NaN())
	}
}

func TestContains(t *testing.T) {
	{
		tests := []struct {
			input []int
			value int
			want  bool
		}{
			{[]int{1, 2, 2, 3, 1}, 10, false},
		}

		for _, test := range tests {
			if r := Contains[int](test.value)(FromSlice[int](test.input)); r != test.want {
				t.Errorf("From(%v).Contains(%v)=%v expected %v", test.input, test.value, r, test.want)
			}
		}
	}
	{
		tests := []struct {
			input []uint
			value uint
			want  bool
		}{
			{[]uint{1, 2, 5, 7, 10}, uint(5), true},
		}

		for _, test := range tests {
			if r := Contains[uint](test.value)(FromSlice[uint](test.input)); r != test.want {
				t.Errorf("From(%v).Contains(%v)=%v expected %v", test.input, test.value, r, test.want)
			}
		}

	}
	{
		tests := []struct {
			input []float32
			value float32
			want  bool
		}{
			{[]float32{}, 1., false},
		}

		for _, test := range tests {
			if r := Contains[float32](test.value)(FromSlice[float32](test.input)); r != test.want {
				t.Errorf("From(%v).Contains(%v)=%v expected %v", test.input, test.value, r, test.want)
			}
		}
	}
}

func TestCount(t *testing.T) {
	{
		tests := []struct {
			input []int
			want  int
		}{
			{[]int{1, 2, 2, 3, 1}, 5},
		}

		for _, test := range tests {

			if r := Count[int](FromSlice[int](test.input)); r != test.want {
				t.Errorf("From(%v).Count()=%v expected %v", test.input, r, test.want)
			}
		}
	}
	{
		tests := []struct {
			input []uint
			want  int
		}{
			{[]uint{1, 2, 5, 7, 10, 12, 15}, 7},
		}

		for _, test := range tests {

			if r := Count[uint](FromSlice[uint](test.input)); r != test.want {
				t.Errorf("From(%v).Count()=%v expected %v", test.input, r, test.want)
			}
		}
	}
	{
		tests := []struct {
			input []float32
			want  int
		}{
			{[]float32{}, 0},
		}

		for _, test := range tests {

			if r := Count[float32](FromSlice[float32](test.input)); r != test.want {
				t.Errorf("From(%v).Count()=%v expected %v", test.input, r, test.want)
			}
		}
	}
}

func TestCountWith(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 2, 3, 1}, 4},
		{[]int{}, 0},
	}

	for _, test := range tests {

		if r := CountWith[int](func(i int) bool {
			return i <= 2
		})(FromSlice[int](test.input)); r != test.want {
			t.Errorf("From(%v).CountWith()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		input []int
		want  []interface{}
	}{
		{[]int{1, 2, 2, 3, 1}, []any{1, true}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {
		if r, ok := First[int](FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).First()=%v %v expected %v", test.input, r, ok, test.want)
		}
	}
}

func TestFirstWith(t *testing.T) {
	tests := []struct {
		input []int
		want  []interface{}
	}{
		{[]int{1, 2, 2, 3, 1}, []any{3, true}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {

		if r, ok := FirstWith[int](func(i int) bool {
			return i > 2
		})(FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).FirstWith()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestForEach(t *testing.T) {
	tests := []struct {
		input []int
		want  interface{}
	}{
		{[]int{1, 2, 2, 35, 111}, []int{2, 4, 4, 70, 222}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		output := []int{}

		ForEach[int](func(i int) {
			output = append(output, i*2)
		})(FromSlice[int](test.input))

		if !reflect.DeepEqual(output, test.want) {
			t.Fatalf("From(%#v).ForEach()=%#v expected=%#v", test.input, output, test.want)
		}
	}
}

func TestForEachIndexed(t *testing.T) {
	tests := []struct {
		input []int
		want  interface{}
	}{
		{[]int{1, 2, 2, 35, 111}, []int{1, 3, 4, 38, 115}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		output := []int{}

		ForEachIndexed[int](func(index, item int) {
			output = append(output, item+index)
		})(FromSlice[int](test.input))

		if !reflect.DeepEqual(output, test.want) {
			t.Fatalf("From(%#v).ForEachIndexed()=%#v expected=%#v", test.input, output, test.want)
		}
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		input []int
		want  []interface{}
	}{
		{[]int{1, 2, 2, 3, 1}, []any{1, true}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {
		if r, ok := Last[int](FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).Last()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestLastWith(t *testing.T) {
	tests := []struct {
		input []int
		want  []any
	}{
		{[]int{1, 2, 2, 3, 1, 4, 2, 5, 1, 1}, []any{5, true}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {

		if r, ok := LastWith[int](func(i int) bool {
			return i > 2
		})(FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).LastWith()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input []int
		want  []any
	}{
		{[]int{1, 2, 2, 3, 1}, []any{3, true}},
		{[]int{1}, []any{1, true}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {
		if r, ok := Max[int](generics.NumericCompare[int])(FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).Max()=%v %v expected %v", test.input, r, ok, test.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input []int
		want  []any
	}{
		{[]int{1, 2, 2, 3, 0}, []any{0, true}},
		{[]int{1}, []any{1, true}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {
		if r, ok := Min[int](generics.NumericCompare[int])(FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).Min()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestResults(t *testing.T) {
	input := []int{1, 2, 3}
	want := []int{1, 2, 3}

	if r := Results[int](FromSlice[int](input)); !reflect.DeepEqual(r, want) {
		t.Errorf("From(%v).Raw()=%v expected %v", input, r, want)
	}
}

func TestSequenceEqual(t *testing.T) {
	tests := []struct {
		input  []int
		input2 []int
		want   bool
	}{
		{[]int{1, 2, 2, 3, 1}, []int{4, 6}, false},
		{[]int{1, -1, 100}, []int{1, -1, 100}, true},
		{[]int{}, []int{}, true},
	}

	for _, test := range tests {

		if r := SequenceEqual[int](FromSlice[int](test.input), FromSlice[int](test.input2)); r != test.want {
			t.Errorf("From(%v).SequenceEqual(%v)=%v expected %v", test.input, test.input2, r, test.want)
		}
	}
}

func TestSingle(t *testing.T) {
	tests := []struct {
		input []int
		want  []any
	}{
		{[]int{1, 2, 2, 3, 1}, []any{0, false}},
		{[]int{1}, []any{1, true}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {
		if r, ok := Single[int](FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).Single()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestSingleWith(t *testing.T) {
	tests := []struct {
		input []int
		want  []any
	}{
		{[]int{1, 2, 2, 3, 1}, []any{3, true}},
		{[]int{1, 1, 1}, []any{0, false}},
		{[]int{5, 1, 1, 10, 2, 2}, []any{0, false}},
		{[]int{}, []any{0, false}},
	}

	for _, test := range tests {
		if r, ok := SingleWith[int](func(i int) bool {
			return i > 2
		})(FromSlice[int](test.input)); r != test.want[0] || ok != test.want[1] {
			t.Errorf("From(%v).SingleWith()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestSumInts(t *testing.T) {
	tests := []struct {
		input []int
		want  int64
	}{
		{[]int{1, 2, 2, 3, 1}, 9},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, test := range tests {
		if r := Sum[int, int64](FromSlice[int](test.input)); r != test.want {
			t.Errorf("From(%v).SumInts()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestSumUInts(t *testing.T) {
	tests := []struct {
		input []uint
		want  uint64
	}{
		{[]uint{1, 2, 2, 3, 1}, 9},
		{[]uint{1}, 1},
		{[]uint{}, 0},
	}

	for _, test := range tests {
		if r := Sum[uint, uint64](FromSlice[uint](test.input)); r != test.want {
			t.Errorf("From(%v).SumInts()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestSumFloats(t *testing.T) {
	tests := []struct {
		input []float32
		want  float64
	}{
		{[]float32{1., 2., 2., 3., 1.}, 9.},
		{[]float32{1.}, 1.},
		{[]float32{}, 0.},
	}

	for _, test := range tests {
		if r := Sum[float32, float64](FromSlice[float32](test.input)); r != test.want {
			t.Errorf("From(%v).SumFloats()=%v expected %v", test.input, r, test.want)
		}
	}
}

func TestToChannel(t *testing.T) {
	c := make(chan int)
	input := []int{1, 2, 3, 4, 5}

	go func() {

		ToChannel[int](FromSlice[int](input), c)
	}()

	result := []int{}
	for value := range c {
		result = append(result, value)
	}

	if !reflect.DeepEqual(result, input) {
		t.Errorf("From(%v).ToChannel()=%v expected %v", input, result, input)
	}
}

func TestToChannelT(t *testing.T) {
	c := make(chan string)
	input := []string{"1", "2", "3", "4", "5"}

	go ToChannel[string](FromSlice[string](input), c)

	result := []string{}
	for value := range c {
		result = append(result, value)
	}

	if !reflect.DeepEqual(result, input) {
		t.Errorf("From(%v).ToChannelT()=%v expected %v", input, result, input)
	}
}

func TestToMap(t *testing.T) {
	input := make(map[int]bool)
	input[1] = true
	input[2] = false
	input[3] = true

	result := ToMap[int, bool](FromMap[int, bool](input))

	if !reflect.DeepEqual(result, input) {
		t.Errorf("From(%v).ToMap()=%v expected %v", input, result, input)
	}
}

func TestToMapBy(t *testing.T) {
	input := make(map[int]bool)
	input[1] = true
	input[2] = false
	input[3] = true

	result := ToMapBy[int, bool, KeyValue[int, bool]](
		func(t KeyValue[int, bool]) int {
			return t.Key
		},
		func(t KeyValue[int, bool]) bool {
			return t.Value
		},
	)(FromMap[int, bool](input))

	if !reflect.DeepEqual(result, input) {
		t.Errorf("From(%v).ToMapBy()=%v expected %v", input, result, input)
	}
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
		want   []int
	}{
		// output is nil slice
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			nil,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		// output is empty slice (cap=0)
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		// ToSlice() overwrites existing elements and reslices.
		{[]int{1, 2, 3},
			[]int{99, 98, 97, 96, 95},
			[]int{1, 2, 3},
		},
		// cap(out)>len(result): we get the same slice, resliced. cap unchanged.
		{[]int{1, 2, 3, 4, 5},
			make([]int, 0, 11),
			[]int{1, 2, 3, 4, 5},
		},
		// cap(out)==len(result): we get the same slice, cap unchanged.
		{[]int{1, 2, 3, 4, 5},
			make([]int, 0, 5),
			[]int{1, 2, 3, 4, 5},
		},
		// cap(out)<len(result): we get a new slice with len(out)=len(result) and cap doubled: cap(out')==2*cap(out)
		{[]int{1, 2, 3, 4, 5},
			make([]int, 0, 4),
			[]int{1, 2, 3, 4, 5},
		},
		// cap(out)<<len(result): trigger capacity to double more than once (26 -> 52 -> 104)
		{make([]int, 100),
			make([]int, 0, 26),
			make([]int, 100),
		},
		// len(out) > len(result): we get the same slice with len(out)=len(result) and cap unchanged: cap(out')==cap(out)
		{[]int{1, 2, 3, 4, 5},
			make([]int, 0, 50),
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		test.output = ToSlice[int](FromSlice[int](test.input))

		// test slice values
		assert.DeepEqual(t, test.output, test.want)
	}
}

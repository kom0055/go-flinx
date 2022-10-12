package flinx

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/kom0055/go-flinx/generics"
	"gotest.tools/v3/assert"
)

func Test_ExampleKeyValue(t *testing.T) {
	m := make(map[int]bool)
	m[10] = true
	assert.DeepEqual(t, Results(FromMap(m)), []KeyValue[int, bool]{{10, true}})
	// Output:
	// [{10 true}]
}

func TestExampleKeyValue_second(t *testing.T) {
	input := []KeyValue[int, bool]{
		{10, true},
	}

	m := ToMap(FromSlice(input))

	assert.DeepEqual(t, m, map[int]bool{10: true})
	// Output:
	// map[10:true]
}

// The following code example demonstrates how
// to use Range to generate a slice of values.
func TestExampleRange(t *testing.T) {
	// Generate a slice of integers from 1 to 10
	// and then select their squares.
	squares := ToSlice(Select(func(x int) int { return x * x })(Range(1, 10)))

	assert.DeepEqual(t, squares, []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100})

	//Output:
	//1
	//4
	//9
	//16
	//25
	//36
	//49
	//64
	//81
	//100
}

// The following code example demonstrates how to use Repeat
// to generate a slice of a repeated value.
func TestExampleRepeat(t *testing.T) {
	slice := ToSlice(Repeat("I like programming.", 5))

	assert.DeepEqual(t, slice,
		[]string{"I like programming.", "I like programming.", "I like programming.",
			"I like programming.", "I like programming."})

	//Output:
	//I like programming.
	//I like programming.
	//I like programming.
	//I like programming.
	//I like programming.

}

func TestExampleQuery(t *testing.T) {
	query := Where(func(i int) bool {
		return i <= 3
	})(FromSlice([]int{1, 2, 3, 4, 5}))

	slice := ToSlice(query)
	assert.DeepEqual(t, slice, []int{1, 2, 3})

	// Output:
	// 1
	// 2
	// 3
}

// The following code example demonstrates how to use Aggregate function
func TestExampleQuery_Aggregate(t *testing.T) {
	fruits := []string{"apple", "mango", "orange", "passionfruit", "grape"}
	aggreFn := Aggregate(func(r, i string) string {
		if len(r) > len(i) {
			return r
		}
		return i
	})
	// Determine which string in the slice is the longest.
	longestName := aggreFn(FromSlice(fruits))
	assert.DeepEqual(t, longestName, "passionfruit")
	// Output:
	// passionfruit
}

// The following code example demonstrates how to use AggregateWithSeed function
func TestExampleQuery_AggregateWithSeed(t *testing.T) {
	ints := []int{4, 8, 8, 3, 9, 0, 7, 8, 2}
	aggreFn := AggregateWithSeed(0, func(total, next int) int {
		if next%2 == 0 {
			return total + 1
		}
		return total
	})
	// Count the even numbers in the array, using a seed value of 0.
	numEven := aggreFn(FromSlice(ints))
	assert.DeepEqual(t, numEven, 6)
	// Output:
	// The number of even integers is: 6
}

// The following code example demonstrates how to use AggregateWithSeedBy function
func TestExampleQuery_AggregateWithSeedBy(t *testing.T) {
	input := []string{"apple", "mango", "orange", "passionfruit", "grape"}
	aggreFn := AggregateWithSeedBy("banana", func(longest, next string) string {
		if len(longest) > len(next) {
			return longest
		}
		return next

	}, func(result string) string {
		return fmt.Sprintf("The fruit with the longest name is %s.", result)
	})
	// Determine whether any string in the array is longer than "banana".
	longestName := aggreFn(FromSlice(input))
	assert.DeepEqual(t, longestName, "The fruit with the longest name is passionfruit.")
	// Output:
	// The fruit with the longest name is passionfruit.
}

// The following code example demonstrates how to
// use Distinct to return distinct elements from a slice of integers.
func TestExampleOrderedQuery_Distinct(t *testing.T) {
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55}
	orderByFn := OrderBy(generics.NumericCompare[int], func(i int) int {
		return i
	})

	distinctAges := ToSlice(Distinct(orderByFn(FromSlice(ages)).Query))
	assert.DeepEqual(t, distinctAges, []int{17, 21, 46, 55})
	// Output:
	// [17 21 46 55]
}

// The following code example demonstrates how to
// use DistinctBy to return distinct elements from a ordered slice of elements.
func TestExampleOrderedQuery_DistinctBy(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	distinctByFn := DistinctBy(func(item Product) int {
		return item.Code
	})
	orderByFn := OrderBy(
		strings.Compare, func(item Product) string {
			return item.Name
		},
	)
	//Order and exclude duplicates.
	noduplicates := ToSlice(distinctByFn(orderByFn(FromSlice(products)).Query))

	assert.DeepEqual(t, noduplicates, []Product{{Name: "apple", Code: 9}, {Name: "lemon", Code: 12}, {Name: "orange", Code: 4}})

	// Output:
	// apple 9
	// lemon 12
	// orange 4
}

// The following code example demonstrates how to use ThenBy to perform
// a secondary ordering of the elements in a slice.
func TestExampleOrderedQuery_ThenBy(t *testing.T) {
	fruits := []string{"grape", "passionfruit", "banana", "mango", "orange", "raspberry", "apple", "blueberry"}

	// Sort the strings first by their length and then
	//alphabetically by passing the identity selector function.

	thenByFn := ThenBy(strings.Compare, func(fruit string) string {
		return fruit
	})
	orderByFn := OrderBy(generics.NumericCompare[int], func(fruit string) int {
		return len(fruit)
	})
	query := ToSlice(thenByFn(orderByFn(FromSlice(fruits))).Query)
	assert.DeepEqual(t, query, []string{"apple", "grape", "mango", "banana", "orange", "blueberry", "raspberry", "passionfruit"})

	// Output:
	// apple
	// grape
	// mango
	// banana
	// orange
	// blueberry
	// raspberry
	// passionfruit
}

// The following code example demonstrates how to use All to determine
// whether all the elements in a slice satisfy a condition.
// Variable allStartWithB is true if all the pet names start with "B"
// or if the pets array is empty.
func TestExampleQuery_All(t *testing.T) {

	type Pet struct {
		Name string
		Age  int
	}

	pets := []Pet{
		{Name: "Barley", Age: 10},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 6},
	}

	// Determine whether all pet names
	// in the array start with 'B'.
	allStartWithB := All(func(pet Pet) bool { return strings.HasPrefix(pet.Name, "B") })(FromSlice(pets))
	assert.DeepEqual(t, allStartWithB, false)
	// Output:
	//
	//  All pet names start with 'B'? false
}

// The following code example demonstrates how to use Any to determine
// whether a slice contains any elements.
func TestExampleQuery_Any(t *testing.T) {

	numbers := []int{1, 2}

	hasElements := Any(FromSlice(numbers))
	assert.DeepEqual(t, hasElements, true)
	// Output:
	// Are there any element in the list? true
}

// The following code example demonstrates how to use AnyWith
// to determine whether any element in a slice satisfies a condition.
func TestExampleQuery_AnyWith(t *testing.T) {

	type Pet struct {
		Name       string
		Age        int
		Vaccinated bool
	}

	pets := []Pet{
		{Name: "Barley", Age: 8, Vaccinated: true},
		{Name: "Boots", Age: 4, Vaccinated: false},
		{Name: "Whiskers", Age: 1, Vaccinated: false},
	}

	// Determine whether any pets over age 1 are also unvaccinated.
	unvaccinated := AnyWith(func(p Pet) bool { return p.Age > 1 && p.Vaccinated == false })(FromSlice(pets))
	assert.DeepEqual(t, unvaccinated, true)
	// Output:
	//
	// Are there any unvaccinated animals over age one? true
}

// The following code example demonstrates how to use Append
// to include an elements in the last position of a slice.
func TestExampleQuery_Append(t *testing.T) {
	input := []int{1, 2, 3, 4}

	q := Append(FromSlice(input), 5)

	last, _ := Last(q)
	assert.DeepEqual(t, last, 5)
	// Output:
	// 5
}

// The following code example demonstrates how to use Average
// to calculate the average of a slice of values.
func TestExampleQuery_Average(t *testing.T) {
	grades := []int{78, 92, 100, 37, 81}

	average := Average(FromSlice(grades))

	assert.DeepEqual(t, average, 77.6)
	// Output:
	// 77.6
}

// The following code example demonstrates how to use Count
// to count the elements in an array.
func TestExampleQuery_Count(t *testing.T) {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	numberOfFruits := Count(FromSlice(fruits))
	assert.DeepEqual(t, numberOfFruits, 6)
	// Output:
	// 6
}

// The following code example demonstrates how to use Contains
// to determine whether a slice contains a specific element.
func TestExampleQuery_Contains(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	has5 := Contains(5)(FromSlice(slice))

	assert.DeepEqual(t, has5, true)
	// Output:
	// Does the slice contains 5? true
}

// The following code example demonstrates how to use CountWith
// to count the even numbers in an array.
func TestExampleQuery_CountWith(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenCount := CountWith(func(item int) bool { return item%2 == 0 })(FromSlice(slice))
	assert.DeepEqual(t, evenCount, 6)
	// Output:
	// 6
}

// The following example demonstrates how to use the DefaultIfEmpty
// method on the results of a group join to perform a left outer join.
//
// The first step in producing a left outer join of two collections is to perform
// an inner join by using a group join. In this example, the list of Person objects
// is inner-joined to the list of Pet objects based on a Person object that matches Pet.Owner.
//
// The second step is to include each element of the first (left) collection in the
// result set even if that element has no matches in the right collection.
// This is accomplished by calling DefaultIfEmpty on each sequence of matching
// elements from the group join.
// In this example, DefaultIfEmpty is called on each sequence of matching Pet elements.
// The method returns a collection that contains a single, default value if the sequence
// of matching Pet elements is empty for any Person element, thereby ensuring that each
// Person element is represented in the result collection.
func TestExampleQuery_DefaultIfEmpty(t *testing.T) {
	type Person struct {
		FirstName string
		LastName  string
	}

	type Pet struct {
		Name  string
		Owner Person
	}

	magnus := Person{FirstName: "Magnus", LastName: "Hedlund"}
	terry := Person{FirstName: "Terry", LastName: "Adams"}
	charlotte := Person{FirstName: "Charlotte", LastName: "Weiss"}
	arlene := Person{FirstName: "Arlene", LastName: "Huff"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	bluemoon := Pet{Name: "Blue Moon", Owner: terry}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	// Create two lists.
	people := []Person{magnus, terry, charlotte, arlene}
	pets := []Pet{barley, boots, whiskers, bluemoon, daisy}

	groupJoinFn := GroupJoin(func(person Person) Person { return person },
		func(pet Pet) Person { return pet.Owner },
		func(person Person, pets []Pet) Group[Person, Pet] {
			return Group[Person, Pet]{Key: person, Group: Results(FromSlice(pets))}
		})

	selectManyByFn := SelectManyBy(
		func(g Group[Person, Pet]) Query[Pet] {
			return DefaultIfEmpty(FromSlice(g.Group), Pet{})
		},
		func(pet Pet, group Group[Person, Pet]) string {
			return fmt.Sprintf("%s: %s", group.Key.FirstName, pet.Name)
		},
	)
	//(FromSlice(people),FromSlice(pets))
	results := ToSlice(selectManyByFn(groupJoinFn(FromSlice(people), FromSlice(pets))))

	assert.DeepEqual(t, results, []string{"Magnus: Daisy", "Terry: Barley", "Terry: Boots",
		"Terry: Blue Moon", "Charlotte: Whiskers", "Arlene: "})

	// Output:
	// Magnus: Daisy
	// Terry: Barley
	// Terry: Boots
	// Terry: Blue Moon
	// Charlotte: Whiskers
	// Arlene:

}

// The following code example demonstrates how to use Distinct
// to return distinct elements from a slice of integers.
func TestExampleQuery_Distinct(t *testing.T) {
	ages := []int{21, 46, 46, 55, 17, 21, 55, 55}

	distinctAges := ToSlice(Distinct(FromSlice(ages)))

	assert.DeepEqual(t, distinctAges, []int{21, 46, 55, 17})
	// Output:
	// [21 46 55 17]
}

// The following code example demonstrates how to
// use DistinctBy to return distinct elements from a ordered slice of elements.
func TestExampleQuery_DistinctBy(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	//Order and exclude duplicates.
	noduplicates := ToSlice(DistinctBy(func(item Product) int { return item.Code })(FromSlice(products)))

	assert.DeepEqual(t, noduplicates, []Product{{Name: "orange", Code: 4},
		{Name: "apple", Code: 9}, {Name: "lemon", Code: 12}})

	// Output:
	// orange 4
	// apple 9
	// lemon 12

}

// The following code example demonstrates how to use the Except
// method to compare two slices of numbers and return elements
// that appear only in the first slice.
func TestExampleQuery_Except(t *testing.T) {
	numbers1 := []float32{2.0, 2.1, 2.2, 2.3, 2.4, 2.5}
	numbers2 := []float32{2.2}

	onlyInFirstSet := ToSlice(Except(FromSlice(numbers1), FromSlice(numbers2)))

	assert.DeepEqual(t, onlyInFirstSet, []float32{2, 2.1, 2.3, 2.4, 2.5})

	// Output:
	//2
	//2.1
	//2.3
	//2.4
	//2.5

}

// The following code example demonstrates how to use the Except
// method to compare two slices of numbers and return elements
// that appear only in the first slice.
func TestExampleQuery_ExceptBy(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	fruits1 := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	fruits2 := []Product{
		{Name: "apple", Code: 9},
	}

	//Order and exclude duplicates.
	except := ToSlice(ExceptBy(func(item Product) int {
		return item.Code
	})(FromSlice(fruits1), FromSlice(fruits2)))

	assert.DeepEqual(t, except, []Product{{Name: "orange", Code: 4}, {Name: "lemon", Code: 12}})

	// Output:
	// orange 4
	// lemon 12

}

// The following code example demonstrates how to use First
// to return the first element of an array.
func TestExampleQuery_First(t *testing.T) {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19}
	n, _ := First(FromSlice(numbers))
	assert.DeepEqual(t, n, 9)

	// Output:
	// 9

}

// The following code example demonstrates how to use FirstWith
// to return the first element of an array that satisfies a condition.
func TestExampleQuery_FirstWith(t *testing.T) {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19}
	first, _ := FirstWith(func(item int) bool { return item > 80 })(FromSlice(numbers))

	assert.DeepEqual(t, first, 92)
	// Output:
	// 92

}

// The following code example demonstrates how to use Intersect
// to return the elements that appear in each of two slices of integers.
func TestExampleQuery_Intersect(t *testing.T) {
	id1 := []int{44, 26, 92, 30, 71, 38}
	id2 := []int{39, 59, 83, 47, 26, 4, 30}

	both := ToSlice(Intersect(FromSlice(id1), FromSlice(id2)))

	assert.DeepEqual(t, both, []int{26, 30})

	// Output:
	// 26
	// 30

}

// The following code example demonstrates how to use IntersectBy
// to return the elements that appear in each of two slices of products with same Code.
func TestExampleQuery_IntersectBy(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	store1 := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
	}

	store2 := []Product{
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	duplicates := ToSlice(IntersectBy(func(p Product) int {
		return p.Code
	})(FromSlice(store1), FromSlice(store2)))

	assert.DeepEqual(t, duplicates, []Product{{Name: "apple", Code: 9}})

	// Output:
	// apple  9

}

// The following code example demonstrates how to use Last
// to return the last element of an array.
func TestExampleQuery_Last(t *testing.T) {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54,
		83, 23, 87, 67, 12, 19}

	last, _ := Last(FromSlice(numbers))
	assert.DeepEqual(t, last, 19)

	//Output:
	//19

}

// The following code example demonstrates how to use LastWith
// to return the last element of an array.
func TestExampleQuery_LastWith(t *testing.T) {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54,
		83, 23, 87, 67, 12, 19}

	last, _ := LastWith(func(n int) bool { return n > 80 })(FromSlice(numbers))

	assert.DeepEqual(t, last, 87)

	//Output:
	//87

}

// The following code example demonstrates how to use Max
// to determine the maximum value in a slice.
func TestExampleQuery_Max(t *testing.T) {
	numbers := []int64{4294967296, 466855135, 81125}

	last, _ := Max(func(i, j int64) int {
		return int(i - j)
	})(FromSlice(numbers))

	assert.DeepEqual(t, last, int64(4294967296))
	//Output:
	//4294967296

}

// The following code example demonstrates how to use Min
// to determine the minimum value in a slice.
func TestExampleQuery_Min(t *testing.T) {
	grades := []int{78, 92, 99, 37, 81}

	min, _ := Min(func(i, j int) int {
		return i - j
	})(FromSlice(grades))

	assert.DeepEqual(t, min, 37)

	//Output:
	//37

}

// The following code example demonstrates how to use OrderByDescending
// to sort the elements of a slice in descending order by using a selector function
func TestExampleQuery_OrderByDescending(t *testing.T) {
	names := []string{"Ned", "Ben", "Susan"}

	result := ToSlice(OrderByDescending(
		strings.Compare, Self[string],
	)(FromSlice(names)).Query)

	assert.DeepEqual(t, result, []string{"Susan", "Ned", "Ben"})
	// Output:
	// [Susan Ned Ben]
}

// The following code example demonstrates how to use ThenByDescending to perform
// a secondary ordering of the elements in a slice in descending order.
func TestExampleOrderedQuery_ThenByDescending(t *testing.T) {

	fruits := []string{"apPLe", "baNanA", "apple", "APple", "orange", "BAnana", "ORANGE", "apPLE"}

	// Sort the strings first ascending by their length and
	// then descending using a custom case insensitive comparer.

	query := ToSlice(ThenByDescending(generics.NumericCompare[byte], func(i string) byte {

		return i[0]
	})(OrderBy(generics.NumericCompare[int], func(i string) int {
		return len(i)
	})(FromSlice(fruits))).Query)

	assert.DeepEqual(t, query, []string{"apPLe", "apple", "apPLE", "APple", "orange", "baNanA", "ORANGE", "BAnana"})

	// Output:
	// apPLe
	// apple
	// apPLE
	// APple
	// orange
	// baNanA
	// ORANGE
	// BAnana

}

// The following code example demonstrates how to use Concat
// to concatenate two slices.
func TestExampleQuery_Concat(t *testing.T) {
	assert.DeepEqual(t, Results(Concat(FromSlice([]int{1, 2, 3}), FromSlice([]int{4, 5, 6}))),
		[]int{1, 2, 3, 4, 5, 6})
	// Output:
	// [1 2 3 4 5 6]
}

func TestExampleQuery_GroupBy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := Results(OrderBy(func(i, j int) int {
		return i - j
	}, func(g Group[int, int]) int {
		return g.Key
	})(GroupBy(func(i int) int { return i % 2 }, func(i int) int {
		return i
	})(FromSlice(input))).Query)
	assert.DeepEqual(t, res,
		[]Group[int, int]{{Key: 0, Group: []int{2, 4, 6, 8}}, {Key: 1, Group: []int{1, 3, 5, 7, 9}}})

	// Output:
	// [{0 [2 4 6 8]} {1 [1 3 5 7 9]}]
}

// The following code example demonstrates how to use GroupJoin
// to perform a grouped join on two slices
func TestExampleQuery_GroupJoin(t *testing.T) {
	fruits := []string{
		"apple",
		"banana",
		"apricot",
		"cherry",
		"clementine",
	}
	res := Results(GroupJoin(
		func(i rune) rune { return i },
		func(i string) rune { return []rune(i)[0] },
		func(outer rune, inners []string) KeyValue[rune, []string] {
			return KeyValue[rune, []string]{outer, inners}
		},
	)(FromString("abc"), FromSlice(fruits)))
	assert.DeepEqual(t, res,
		[]KeyValue[rune, []string]{{Key: 'a', Value: []string{"apple", "apricot"}},
			{Key: 'b', Value: []string{"banana"}},
			{Key: 'c', Value: []string{"cherry", "clementine"}}})
	// Output:
	// [{a [apple apricot]} {b [banana]} {c [cherry clementine]}]
}

// The following code example demonstrates how to use IndexOf
// to retrieve the position of an item in the array and then
// update that item.
func TestExampleQuery_IndexOf(t *testing.T) {
	type Item struct {
		ID   uint64
		Name string
	}
	items := []Item{
		{
			ID:   1,
			Name: "Joe",
		},
		{
			ID:   2,
			Name: "Bob",
		},
		{
			ID:   3,
			Name: "Rickster",
		},
		{
			ID:   4,
			Name: "Jim",
		},
	}

	index := IndexOf(func(item Item) bool {
		return item.Name == "Rickster"
	})(FromSlice(items))
	assert.DeepEqual(t, index, 2)

	if index >= 0 {
		// We found the item in the array. Change the name using the index.
		items[index].Name = "Joshua"
		fmt.Println("Item found at:", index, "new name:", items[index].Name)
	}
	// Output:
	// Item found at: 2 new name: Joshua
}

// The following code example demonstrates how to use Join
// to perform an inner join of two slices based on a common key.
func TestExampleQuery_Join(t *testing.T) {
	fruits := []string{
		"apple",
		"banana",
		"apricot",
		"cherry",
		"clementine",
	}

	q :=
		Join(
			Self[int],
			func(i string) int {
				return len(i)
			},
			func(outer int, inner string) KeyValue[int, string] {
				return KeyValue[int, string]{outer, inner}
			},
		)(Range(1, 10), FromSlice(fruits))

	assert.DeepEqual(t, Results(q), []KeyValue[int, string]{
		{Key: 5, Value: "apple"}, {Key: 6, Value: "banana"}, {Key: 6, Value: "cherry"},
		{Key: 7, Value: "apricot"}, {Key: 10, Value: "clementine"},
	})
	// Output:
	// [{5 apple} {6 banana} {6 cherry} {7 apricot} {10 clementine}]
}

// The following code example demonstrates how to use OrderBy
// to sort the elements of a slice.
func TestExampleQuery_OrderBy(t *testing.T) {

	q := ThenByDescending(
		generics.NumericCompare[int],
		Self[int],
	)(OrderBy(
		generics.NumericCompare[int],
		func(v int) int {
			return v % 2
		},
	)(Range(1, 10)))

	assert.DeepEqual(t, Results(q.Query), []int{10, 8, 6, 4, 2, 9, 7, 5, 3, 1})
	// Output:
	// [10 8 6 4 2 9 7 5 3 1]
}

// The following code example demonstrates how to use Prepend
// to include an elements in the first position of a slice.
func TestExampleQuery_Prepend(t *testing.T) {
	input := []int{2, 3, 4, 5}

	first, _ := First(Prepend(FromSlice(input), 1))
	assert.DeepEqual(t, first, 1)
	// Output:
	// 1
}

// The following code example demonstrates how to use Reverse
// to reverse the order of elements in a string.
func TestExampleQuery_Reverse(t *testing.T) {
	input := "apple"

	output := ToSlice(Reverse(FromString(input)))
	assert.DeepEqual(t, string(output), "elppa")

	// Output:
	// elppa
}

// The following code example demonstrates how to use Select
// to project over a slice of values.
func TestExampleQuery_Select(t *testing.T) {
	squares := ToSlice(Select(func(x int) int {
		return x * x
	})(Range(1, 10)))

	assert.DeepEqual(t, squares, []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100})
	// Output:
	// [1 4 9 16 25 36 49 64 81 100]
}

func TestExampleQuery_SelectMany(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6, 7}}

	res := Results(SelectMany(func(i []int) Query[int] {
		return FromSlice(i)
	})(FromSlice(input)))

	assert.DeepEqual(t, res, []int{1, 2, 3, 4, 5, 6, 7})
	// Output:
	// [1 2 3 4 5 6 7]
}

// The following code example demonstrates how to use Select
// to project over a slice of values and use the index of each element.
func TestExampleQuery_SelectIndexed(t *testing.T) {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}

	result := ToSlice(SelectIndexed(func(i int, s string) string {
		return s[:i]
	})(FromSlice(fruits)))

	assert.DeepEqual(t, result, []string{"", "b", "ma", "ora", "pass", "grape"})
	// Output:
	// [ b ma ora pass grape]

}

// The following code example demonstrates how to use SelectManyByIndexed
// to perform a one-to-many projection over an array and use the index of each outer element.
func TestExampleQuery_SelectManyByIndexed(t *testing.T) {
	type Pet struct {
		Name string
	}

	type Person struct {
		Name string
		Pets []Pet
	}

	magnus := Person{
		Name: "Hedlund, Magnus",
		Pets: []Pet{{Name: "Daisy"}},
	}

	terry := Person{
		Name: "Adams, Terry",
		Pets: []Pet{{Name: "Barley"}, {Name: "Boots"}},
	}
	charlotte := Person{
		Name: "Weiss, Charlotte",
		Pets: []Pet{{Name: "Whiskers"}},
	}

	people := []Person{magnus, terry, charlotte}

	results := ToSlice(SelectManyByIndexed(
		func(i int, p Person) Query[string] {
			return Select(func(pet Pet) string {
				return fmt.Sprintf("%d - %s", i, pet.Name)
			})(FromSlice(p.Pets))
		},

		func(pet string, person Person) string {
			return fmt.Sprintf("Pet: %s, Owner: %s", pet, person.Name)
		},
	)(FromSlice(people)))

	assert.DeepEqual(t, results, []string{"Pet: 0 - Daisy, Owner: Hedlund, Magnus",
		"Pet: 1 - Barley, Owner: Adams, Terry",
		"Pet: 1 - Boots, Owner: Adams, Terry",
		"Pet: 2 - Whiskers, Owner: Weiss, Charlotte",
	})

	// Output:
	// Pet: 0 - Daisy, Owner: Hedlund, Magnus
	// Pet: 1 - Barley, Owner: Adams, Terry
	// Pet: 1 - Boots, Owner: Adams, Terry
	// Pet: 2 - Whiskers, Owner: Weiss, Charlotte

}

// The following code example demonstrates how to use SelectManyIndexed
// to perform a one-to-many projection over an slice of log data and print out their contents.
func TestExampleQuery_SelectManyIndexed(t *testing.T) {
	type LogFile struct {
		Name  string
		Lines []string
	}

	file1 := LogFile{
		Name: "file1.log",
		Lines: []string{
			"INFO: 2013/11/05 18:11:01 main.go:44: Special Information",
			"WARNING: 2013/11/05 18:11:01 main.go:45: There is something you need to know about",
			"ERROR: 2013/11/05 18:11:01 main.go:46: Something has failed",
		},
	}

	file2 := LogFile{
		Name: "file2.log",
		Lines: []string{
			"INFO: 2013/11/05 18:11:01 main.go:46: Everything is ok",
		},
	}

	file3 := LogFile{
		Name: "file3.log",
		Lines: []string{
			"2013/11/05 18:42:26 Hello World",
		},
	}

	logFiles := []LogFile{file1, file2, file3}

	results := ToSlice(SelectManyIndexed(
		func(fileIndex int, file LogFile) Query[string] {
			return SelectIndexed(
				func(lineIndex int, line string) string {
					return fmt.Sprintf("File:[%d] - %s => line: %d - %s", fileIndex+1, file.Name, lineIndex+1, line)
				},
			)(FromSlice(file.Lines))
		},
	)(FromSlice(logFiles)))

	assert.DeepEqual(t, results, []string{
		"File:[1] - file1.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:44: Special Information",
		"File:[1] - file1.log => line: 2 - WARNING: 2013/11/05 18:11:01 main.go:45: There is something you need to know about",
		"File:[1] - file1.log => line: 3 - ERROR: 2013/11/05 18:11:01 main.go:46: Something has failed",
		"File:[2] - file2.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:46: Everything is ok",
		"File:[3] - file3.log => line: 1 - 2013/11/05 18:42:26 Hello World",
	})
	// Output:
	// File:[1] - file1.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:44: Special Information
	// File:[1] - file1.log => line: 2 - WARNING: 2013/11/05 18:11:01 main.go:45: There is something you need to know about
	// File:[1] - file1.log => line: 3 - ERROR: 2013/11/05 18:11:01 main.go:46: Something has failed
	// File:[2] - file2.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:46: Everything is ok
	// File:[3] - file3.log => line: 1 - 2013/11/05 18:42:26 Hello World

}

// The following code example demonstrates how to use SelectMany
// to perform a one-to-many projection over a slice
func TestExampleQuery_SelectManyBy(t *testing.T) {

	type Pet struct {
		Name string
	}

	type Person struct {
		Name string
		Pets []Pet
	}

	magnus := Person{
		Name: "Hedlund, Magnus",
		Pets: []Pet{{Name: "Daisy"}},
	}

	terry := Person{
		Name: "Adams, Terry",
		Pets: []Pet{{Name: "Barley"}, {Name: "Boots"}},
	}
	charlotte := Person{
		Name: "Weiss, Charlotte",
		Pets: []Pet{{Name: "Whiskers"}},
	}

	people := []Person{magnus, terry, charlotte}
	results := ToSlice(SelectManyBy(
		func(person Person) Query[Pet] {
			return FromSlice(person.Pets)
		},
		func(pet Pet, person Person) string {
			return fmt.Sprintf("Owner: %s, Pet: %s", person.Name, pet.Name)
		},
	)(FromSlice(people)))

	assert.DeepEqual(t, results, []string{"Owner: Hedlund, Magnus, Pet: Daisy",
		"Owner: Adams, Terry, Pet: Barley",
		"Owner: Adams, Terry, Pet: Boots",
		"Owner: Weiss, Charlotte, Pet: Whiskers",
	})
	// Output:
	// Owner: Hedlund, Magnus, Pet: Daisy
	// Owner: Adams, Terry, Pet: Barley
	// Owner: Adams, Terry, Pet: Boots
	// Owner: Weiss, Charlotte, Pet: Whiskers
}

// The following code example demonstrates how to use SequenceEqual
// to determine whether two slices are equal.
func TestExampleQuery_SequenceEqual(t *testing.T) {
	type Pet struct {
		Name string
		Age  int
	}

	pets1 := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
		{Name: "Daisy", Age: 4},
	}

	pets2 := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
		{Name: "Daisy", Age: 4},
	}

	equal := SequenceEqual(FromSlice(pets1), FromSlice(pets2))
	assert.DeepEqual(t, equal, true)
	// Output:
	// Are the lists equals? true
}

// The following code example demonstrates how to use Single
// to select the only element of a slice.
func TestExampleQuery_Single(t *testing.T) {
	fruits1 := []string{"orange"}

	fruit1, _ := Single(FromSlice(fruits1))

	assert.DeepEqual(t, fruit1, "orange")
	// Output:
	// orange
}

// The following code example demonstrates how to use SingleWith
// to select the only element of a slice that satisfies a condition.
func TestExampleQuery_SingleWith(t *testing.T) {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}

	fruit, _ := SingleWith(func(f string) bool { return len(f) > 10 })(FromSlice(fruits))

	assert.DeepEqual(t, fruit, "passionfruit")
	// Output:
	// passionfruit
}

// The following code example demonstrates how to use Skip
// to skip a specified number of elements in a sorted array
// and return the remaining elements.
func TestExampleQuery_Skip(t *testing.T) {
	grades := []int{59, 82, 70, 56, 92, 98, 85}
	lowerGrades := ToSlice(

		Skip(
			OrderByDescending(
				generics.NumericCompare[int],
				func(g int) int { return g },
			)(FromSlice(grades)).Query,
			3,
		),
	)

	//All grades except the top three are:
	assert.DeepEqual(t, lowerGrades, []int{82, 70, 59, 56})
	// Output:
	// [82 70 59 56]
}

// The following code example demonstrates how to use SkipWhile
// to skip elements of an array as long as a condition is true.
func TestExampleQuery_SkipWhile(t *testing.T) {
	grades := []int{59, 82, 70, 56, 92, 98, 85}

	lowerGrades := ToSlice(

		SkipWhile(func(g int) bool { return g >= 80 })(
			OrderByDescending(
				generics.NumericCompare[int],
				func(g int) int { return g },
			)(FromSlice(grades)).Query,
		),
	)

	// All grades below 80:
	assert.DeepEqual(t, lowerGrades, []int{70, 59, 56})

	// Output:
	// [70 59 56]
}

// The following code example demonstrates how to use SkipWhileIndexed
// to skip elements of an array as long as a condition that depends
// on the element's index is true.
func TestExampleQuery_SkipWhileIndexed(t *testing.T) {
	amounts := []int{5000, 2500, 9000, 8000, 6500, 4000, 1500, 5500}

	query := ToSlice(
		SkipWhileIndexed(func(index int, amount int) bool { return amount > index*1000 })(FromSlice(amounts)),
	)

	assert.DeepEqual(t, query, []int{4000, 1500, 5500})
	// Output:
	// [4000 1500 5500]

}

// The following code example demonstrates how to use Sort
// to order elements of an slice.
func TestExampleQuery_Sort(t *testing.T) {
	amounts := []int{5000, 2500, 9000, 8000, 6500, 4000, 1500, 5500}

	query := ToSlice(
		Sort(func(i, j int) bool { return i < j })(FromSlice(amounts)),
	)
	assert.DeepEqual(t, query, []int{1500, 2500, 4000, 5000, 5500, 6500, 8000, 9000})

	// Output:
	// [1500 2500 4000 5000 5500 6500 8000 9000]

}

// The following code example demonstrates how to use SumFloats
// to sum the values of a slice.
func TestExampleQuery_SumFloats(t *testing.T) {
	numbers := []float64{43.68, 1.25, 583.7, 6.5}
	sum := Sum(FromSlice(numbers))
	assert.DeepEqual(t, sum, 635.130000)
	// Output:
	// The sum of the numbers is 635.130000.

}

// The following code example demonstrates how to use SumInts
// to sum the values of a slice.
func TestExampleQuery_SumInts(t *testing.T) {
	numbers := []int{43, 1, 583, 6}

	sum := Sum(FromSlice(numbers))

	assert.DeepEqual(t, sum, 633)
	// Output:
	// The sum of the numbers is 633.

}

// The following code example demonstrates how to use SumUInts
// to sum the values of a slice.
func TestExampleQuery_SumUInts(t *testing.T) {
	numbers := []uint{43, 1, 583, 6}

	sum := Sum(FromSlice(numbers))

	assert.DeepEqual(t, sum, uint(633))
	// Output:
	// The sum of the numbers is 633.

}

// The following code example demonstrates how to use Take
//
//	to return elements from the start of a slice.
func TestExampleQuery_Take(t *testing.T) {
	grades := []int{59, 82, 70, 56, 92, 98, 85}

	topThreeGrades := ToSlice(
		Take(OrderByDescending(generics.NumericCompare[int], func(g int) int { return g })(FromSlice(grades)).Query, 3),
	)
	assert.DeepEqual(t, topThreeGrades, []int{98, 92, 85})

	// Output:
	// The top three grades are: [98 92 85]
}

// The following code example demonstrates how to use TakeWhile
// to return elements from the start of a slice.
func TestExampleQuery_TakeWhile(t *testing.T) {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}

	query := ToSlice(TakeWhile(func(fruit string) bool { return fruit != "orange" })(FromSlice(fruits)))

	assert.DeepEqual(t, query, []string{"apple", "banana", "mango"})
	// Output:
	// [apple banana mango]
}

// The following code example demonstrates how to use TakeWhileIndexed
// to return elements from the start of a slice as long as
// a condition that uses the element's index is true.
func TestExampleQuery_TakeWhileIndexed(t *testing.T) {

	fruits := []string{"apple", "passionfruit", "banana", "mango",
		"orange", "blueberry", "grape", "strawberry"}

	query := ToSlice(TakeWhileIndexed(
		func(index int, fruit string) bool { return len(fruit) >= index },
	)(FromSlice(fruits)))

	assert.DeepEqual(t, query, []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry"})
	// Output:
	// [apple passionfruit banana mango orange blueberry]
}

// The following code example demonstrates how to use ToChannel
// to send a slice to a channel.
func TestExampleQuery_ToChannel(t *testing.T) {
	c := make(chan int)

	go func() {
		ToChannel(Repeat(10, 3), c)
	}()

	for i := range c {
		assert.DeepEqual(t, i, 10)
	}
	// Output:
	// 10
	// 10
	// 10
}

// The following code example demonstrates how to use ToChannelT
// to send a slice to a typed channel.
func TestExampleQuery_ToChannelT(t *testing.T) {
	c := make(chan string)

	go ToChannel(Repeat("ten", 3), c)

	for i := range c {
		assert.DeepEqual(t, i, "ten")
	}
	// Output:
	// ten
	// ten
	// ten
}

// The following code example demonstrates how to use ToMap to populate a map.
func TestExampleQuery_ToMap(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	map1 := ToMap(Select(
		func(item Product) KeyValue[int, string] {
			return KeyValue[int, string]{Key: item.Code, Value: item.Name}
		},
	)(FromSlice(products)))

	assert.DeepEqual(t, map1, map[int]string{4: "orange", 9: "apple", 12: "lemon"})

	// Output:
	// orange
	// apple
	// lemon
}

// The following code example demonstrates how to use ToMapBy
// by using a key and value selectors to populate a map.
func TestExampleQuery_ToMapBy(t *testing.T) {
	input := [][]any{{1, true}}

	result := ToMapBy(func(t []any) int {
		return t[0].(int)
	},
		func(t []any) bool {
			return t[1].(bool)
		},
	)(FromSlice(input))

	assert.DeepEqual(t, result, map[int]bool{1: true})
	// Output:
	// map[1:true]
}

// The following code example demonstrates how to use ToSlice to populate a slice.
func TestExampleQuery_ToSlice(t *testing.T) {
	result := ToSlice(Range(1, 10))
	assert.DeepEqual(t, result, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}

// The following code example demonstrates how to use Union
// to obtain the union of two slices of integers.
func TestExampleQuery_Union(t *testing.T) {
	q := Results(Union(Range(1, 10), Range(6, 10)))

	assert.DeepEqual(t, q, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	// Output:
	// [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]
}

// The following code example demonstrates how to use Where
// to filter a slices.
func TestExampleQuery_Where(t *testing.T) {
	fruits := []string{"apple", "passionfruit", "banana", "mango",
		"orange", "blueberry", "grape", "strawberry"}
	query := ToSlice(Where(func(f string) bool {
		return len(f) > 6
	})(FromSlice(fruits)))

	assert.DeepEqual(t, query, []string{"passionfruit", "blueberry", "strawberry"})

	// Output:
	// [passionfruit blueberry strawberry]
}

// The following code example demonstrates how to use WhereIndexed
// to filter a slice based on a predicate that involves the index of each element.
func TestExampleQuery_WhereIndexed(t *testing.T) {
	numbers := []int{0, 30, 20, 15, 90, 85, 40, 75}
	query := ToSlice(WhereIndexed(func(index, number int) bool {
		return number <= index*10
	})(FromSlice(numbers)))
	assert.DeepEqual(t, query, []int{0, 20, 15, 40})

	// Output:
	// [0 20 15 40]
}

// The following code example demonstrates how to use the Zip
// method to merge two slices.
func TestExampleQuery_Zip(t *testing.T) {
	number := []int{1, 2, 3, 4, 5}
	words := []string{"one", "two", "three"}

	q := Zip(func(a int, b string) []any {
		return []any{a, b}
	})(
		FromSlice(number), FromSlice(words),
	)
	assert.DeepEqual(t, ToSlice(q), [][]any{{1, "one"}, {2, "two"}, {3, "three"}})

	// Output:
	// [[1 one] [2 two] [3 three]]
}

// The following code example demonstrates how to use ThenByDescendingT to perform
// a order in a slice of dates by year, and then by month descending.
func TestExampleOrderedQuery_ThenByDescendingT(t *testing.T) {
	dates := []time.Time{
		time.Date(2015, 3, 23, 0, 0, 0, 0, time.Local),
		time.Date(2014, 7, 11, 0, 0, 0, 0, time.Local),
		time.Date(2013, 5, 4, 0, 0, 0, 0, time.Local),
		time.Date(2015, 1, 2, 0, 0, 0, 0, time.Local),
		time.Date(2015, 7, 10, 0, 0, 0, 0, time.Local),
	}

	orderedDates := ToSlice(Select(func(t time.Time) string {
		return t.Format("2006-Jan-02")
	})(ThenByDescending(generics.NumericCompare[int], func(date time.Time) int {
		return int(date.Month())
	})(OrderBy(generics.NumericCompare[int], func(date time.Time) int {
		return date.Year()
	})(FromSlice(dates))).Query))

	assert.DeepEqual(t, orderedDates, []string{
		"2013-May-04", "2014-Jul-11", "2015-Jul-10", "2015-Mar-23", "2015-Jan-02",
	})

	// Output:
	// 2013-May-04
	// 2014-Jul-11
	// 2015-Jul-10
	// 2015-Mar-23
	// 2015-Jan-02

}

// The following code example demonstrates how to use ThenByT to perform
// a orders in a slice of dates by year, and then by day.
func TestExampleOrderedQuery_ThenByT(t *testing.T) {
	dates := []time.Time{
		time.Date(2015, 3, 23, 0, 0, 0, 0, time.Local),
		time.Date(2014, 7, 11, 0, 0, 0, 0, time.Local),
		time.Date(2013, 5, 4, 0, 0, 0, 0, time.Local),
		time.Date(2015, 1, 2, 0, 0, 0, 0, time.Local),
		time.Date(2015, 7, 10, 0, 0, 0, 0, time.Local),
	}

	orderedDates := ToSlice(Select(func(t time.Time) string {
		return t.Format("2006-Jan-02")
	})(ThenBy(generics.NumericCompare[int], func(date time.Time) int {
		return int(date.Day())
	})(OrderBy(generics.NumericCompare[int], func(date time.Time) int {
		return date.Year()
	})(FromSlice(dates))).Query))

	assert.DeepEqual(t, orderedDates, []string{
		"2013-May-04", "2014-Jul-11", "2015-Jan-02", "2015-Jul-10", "2015-Mar-23",
	})

	// Output:
	// 2013-May-04
	// 2014-Jul-11
	// 2015-Jan-02
	// 2015-Jul-10
	// 2015-Mar-23

}

// The following code example demonstrates how to reverse
// the order of words in a string using AggregateT.
func TestExampleQuery_AggregateT(t *testing.T) {
	sentence := "the quick brown fox jumps over the lazy dog"
	// Split the string into individual words.
	words := strings.Split(sentence, " ")

	// Prepend each word to the beginning of the
	// new sentence to reverse the word order.

	reversed := Aggregate(
		func(workingSentence string, next string) string { return next + " " + workingSentence },
	)(FromSlice(words))
	assert.DeepEqual(t, reversed, "dog lazy the over jumps fox brown quick the")
	// Output:
	// dog lazy the over jumps fox brown quick the
}

// The following code example demonstrates how to use AggregateWithSeed function
func TestExampleQuery_AggregateWithSeedT(t *testing.T) {

	fruits := []string{"apple", "mango", "orange", "passionfruit", "grape"}

	// Determine whether any string in the array is longer than "banana".
	longestName := AggregateWithSeed(
		"banan", func(longest, next string) string {
			if len(next) > len(longest) {
				return next
			}
			return longest
		},
	)(FromSlice(fruits))
	assert.DeepEqual(t, longestName, "passionfruit")

	// Output:
	// The fruit with the longest name is passionfruit.

}

// The following code example demonstrates how to use AggregateWithSeedByT function
func TestExampleQuery_AggregateWithSeedByT(t *testing.T) {
	input := []string{"apple", "mango", "orange", "passionfruit", "grape"}

	// Determine whether any string in the array is longer than "banana".
	longestName := AggregateWithSeedBy("banana",
		func(longest string, next string) string {
			if len(longest) > len(next) {
				return longest
			}
			return next

		},
		// Return the final result
		func(result string) string {
			return fmt.Sprintf("The fruit with the longest name is %s.", result)
		})(FromSlice(input))
	assert.DeepEqual(t, longestName, "The fruit with the longest name is passionfruit.")
	// Output:
	// The fruit with the longest name is passionfruit.
}

// The following code example demonstrates how to use AllT
// to get the students having all marks greater than 70.
func TestExampleQuery_AllT(t *testing.T) {

	type Student struct {
		Name  string
		Marks []int
	}

	students := []Student{
		{Name: "Hugo", Marks: []int{91, 88, 76, 93}},
		{Name: "Rick", Marks: []int{70, 73, 66, 90}},
		{Name: "Michael", Marks: []int{73, 80, 75, 88}},
		{Name: "Fadi", Marks: []int{82, 75, 66, 84}},
		{Name: "Peter", Marks: []int{67, 78, 70, 82}},
	}

	approvedStudents := ToSlice(
		Select(func(t Student) string {
			return t.Name
		})(Where(
			func(student Student) bool {
				return All(
					func(mark int) bool { return mark > 70 },
				)(FromSlice(student.Marks))

			},
		)(FromSlice(students))))

	//List of approved students
	assert.DeepEqual(t, approvedStudents, []string{"Hugo", "Michael"})

	// Output:
	// Hugo
	// Michael
}

// The following code example demonstrates how to use AnyWithT
// to get the students with any mark lower than 70.
func TestExampleQuery_AnyWithT(t *testing.T) {
	type Student struct {
		Name  string
		Marks []int
	}

	students := []Student{
		{Name: "Hugo", Marks: []int{91, 88, 76, 93}},
		{Name: "Rick", Marks: []int{70, 73, 66, 90}},
		{Name: "Michael", Marks: []int{73, 80, 75, 88}},
		{Name: "Fadi", Marks: []int{82, 75, 66, 84}},
		{Name: "Peter", Marks: []int{67, 78, 70, 82}},
	}

	studentsWithAnyMarkLt70 := ToSlice(
		Select(
			func(t Student) string {
				return t.Name
			},
		)(

			Where(
				func(student Student) bool {

					return AnyWith(
						func(mark int) bool { return mark < 70 },
					)(FromSlice(
						student.Marks,
					))

				},
			)(FromSlice(students)),
		),
	)

	//List of students with any mark lower than 70

	assert.DeepEqual(t, studentsWithAnyMarkLt70, []string{"Rick", "Fadi", "Peter"})
	// Output:
	// Rick
	// Fadi
	// Peter

}

// The following code example demonstrates how to use CountWithT
// to count the elements in an slice that satisfy a condition.
func TestExampleQuery_CountWithT(t *testing.T) {
	type Pet struct {
		Name       string
		Vaccinated bool
	}

	pets := []Pet{
		{Name: "Barley", Vaccinated: true},
		{Name: "Boots", Vaccinated: false},
		{Name: "Whiskers", Vaccinated: false},
	}

	numberUnvaccinated := CountWith(
		func(p Pet) bool {

			return p.Vaccinated == false

		},
	)(FromSlice(pets))
	assert.DeepEqual(t, numberUnvaccinated, 2)

	//Output:
	//There are 2 unvaccinated animals.
}

// The following code example demonstrates how to use DistinctByT
// to return distinct elements from a slice of structs.
func TestExampleQuery_DistinctByT(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "apple", Code: 9},
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
	}

	//Exclude duplicates.
	noduplicates := ToSlice(
		Select(
			func(p Product) []any {

				return []any{p.Name, p.Code}
			},
		)(

			DistinctBy(func(item Product) int { return item.Code })(FromSlice(products)),
		),
	)

	assert.DeepEqual(t, noduplicates, [][]any{

		{"apple", 9},
		{"orange", 4},
		{"lemon", 12},
	})

	// Output:
	// apple 9
	// orange 4
	// lemon 12
}

// The following code example demonstrates how to use ExceptByT
func TestExampleQuery_ExceptByT(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	fruits1 := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	fruits2 := []Product{
		{Name: "apple", Code: 9},
	}

	//Order and exclude duplicates.

	expect := ToSlice(
		Select(
			func(p Product) []any {

				return []any{p.Name, p.Code}
			},
		)(
			ExceptBy(
				func(item Product) int { return item.Code },
			)(FromSlice(fruits1), FromSlice(fruits2)),
		))
	assert.DeepEqual(t, expect, [][]any{
		{"orange", 4},
		{"lemon", 12},
	})

	// Output:
	// orange 4
	// lemon 12

}

// The following code example demonstrates how to use FirstWithT
// to return the first element of an array that satisfies a condition.
func TestExampleQuery_FirstWithT(t *testing.T) {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54, 83, 23, 87, 435, 67, 12, 19}

	first, _ := FirstWith(
		func(item int) bool { return item > 80 },
	)(FromSlice(numbers))
	assert.DeepEqual(t, first, 92)
	// Output:
	// 92

}

// The following code example demonstrates how to use ForEach
// to output all elements of an array.
func TestExampleQuery_ForEach(t *testing.T) {
	fruits := []string{"orange", "apple", "lemon", "apple"}
	ForEach(
		func(fruit string) {
			fmt.Println(fruit)
		},
	)(FromSlice(fruits))

	// Output:
	// orange
	// apple
	// lemon
	// apple
}

// The following code example demonstrates how to use ForEachIndexed
// to output all elements of an array with its index.
func TestExampleQuery_ForEachIndexed(t *testing.T) {
	fruits := []string{"orange", "apple", "lemon", "apple"}

	ForEachIndexed(
		func(i int, fruit string) {
			fmt.Printf("%d.%s\n", i, fruit)
		},
	)(FromSlice(fruits))
	// Output:
	// 0.orange
	// 1.apple
	// 2.lemon
	// 3.apple
}

// The following code example demonstrates how to use ForEachT
// to output all elements of an array.
func TestExampleQuery_ForEachT(t *testing.T) {
	fruits := []string{"orange", "apple", "lemon", "apple"}
	ForEach(func(fruit string) {
		fmt.Println(fruit)
	})(FromSlice(fruits))
	// Output:
	// orange
	// apple
	// lemon
	// apple
}

// The following code example demonstrates how to use ForEachIndexedT
// to output all elements of an array with its index.
func TestExampleQuery_ForEachIndexedT(t *testing.T) {
	fruits := []string{"orange", "apple", "lemon", "apple"}

	ForEachIndexed(func(i int, fruit string) {
		fmt.Printf("%d.%s\n", i, fruit)
	})(FromSlice(fruits))
	// Output:
	// 0.orange
	// 1.apple
	// 2.lemon
	// 3.apple
}

// The following code example demonstrates how to use GroupByT
// to group the elements of a slice.
func TestExampleQuery_GroupByT(t *testing.T) {

	type Pet struct {
		Name string
		Age  int
	}
	// Create a list of pets.
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
		{Name: "Daisy", Age: 4},
	}

	// Group the pets using Age as the key value
	// and selecting only the pet's Name for each value.

	query := ToSlice(OrderBy(generics.NumericCompare[int], func(g Group[int, string]) int { return g.Key })(GroupBy(
		func(p Pet) int { return p.Age },
		func(p Pet) string { return p.Name },
	)(FromSlice(pets))).Query)
	assert.DeepEqual(t, query, []Group[int, string]{
		{1, []string{"Whiskers"}},
		{4, []string{"Boots", "Daisy"}},
		{8, []string{"Barley"}},
	})

	// Output:
	// 1
	//   Whiskers
	// 4
	//   Boots
	//   Daisy
	// 8
	//   Barley
}

// The following code example demonstrates how to use GroupJoinT
//
//	to perform a grouped join on two slices.
func TestExampleQuery_GroupJoinT(t *testing.T) {

	type Person struct {
		Name string
	}

	type Pet struct {
		Name  string
		Owner Person
	}

	magnus := Person{Name: "Hedlund, Magnus"}
	terry := Person{Name: "Adams, Terry"}
	charlotte := Person{Name: "Weiss, Charlotte"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	people := []Person{magnus, terry, charlotte}
	pets := []Pet{barley, boots, whiskers, daisy}

	// Create a slice where each element is a KeyValue
	// that contains a person's name as the key and a slice of strings
	// of names of the pets they own as a value.

	q := ToSlice(
		GroupJoin(
			func(p Person) Person { return p },
			func(p Pet) Person { return p.Owner },
			func(person Person, pets []Pet) KeyValue[string, []string] {

				return KeyValue[string, []string]{person.Name, ToSlice(
					Select(
						func(pet Pet) string { return pet.Name },
					)(FromSlice(pets)),
				)}
			},
		)(FromSlice(people), FromSlice(pets)),
	)

	assert.DeepEqual(t, q, []KeyValue[string, []string]{
		{
			"Hedlund, Magnus", []string{
				"Daisy",
			},
		},
		{
			"Adams, Terry", []string{
				"Barley",
				"Boots",
			},
		},
		{
			"Weiss, Charlotte", []string{
				"Whiskers",
			},
		},
	})

	// Output:
	// Hedlund, Magnus:
	//   Daisy
	// Adams, Terry:
	//   Barley
	//   Boots
	// Weiss, Charlotte:
	//   Whiskers
}

// The following code example demonstrates how to use IntersectByT
// to return the elements that appear in each of two slices of products
// with same Code.
func TestExampleQuery_IntersectByT(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	store1 := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
	}

	store2 := []Product{
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	duplicates := ToSlice(
		Select(
			func(p Product) []any {

				return []any{p.Name, p.Code}
			},
		)(
			IntersectBy(
				func(p Product) int { return p.Code },
			)(
				FromSlice(store1),
				FromSlice(store2),
			),
		),
	)
	assert.DeepEqual(t, duplicates, [][]any{
		{"apple", 9},
	})

	// Output:
	// apple  9

}

// The following code example demonstrates how to use JoinT
// to perform an inner join of two slices based on a common key.
func TestExampleQuery_JoinT(t *testing.T) {
	type Person struct {
		Name string
	}

	type Pet struct {
		Name  string
		Owner Person
	}

	magnus := Person{Name: "Hedlund, Magnus"}
	terry := Person{Name: "Adams, Terry"}
	charlotte := Person{Name: "Weiss, Charlotte"}

	barley := Pet{Name: "Barley", Owner: terry}
	boots := Pet{Name: "Boots", Owner: terry}
	whiskers := Pet{Name: "Whiskers", Owner: charlotte}
	daisy := Pet{Name: "Daisy", Owner: magnus}

	people := []Person{magnus, terry, charlotte}
	pets := []Pet{barley, boots, whiskers, daisy}

	// Create a list of Person-Pet pairs where
	// each element is an anonymous type that contains a
	// Pet's name and the name of the Person that owns the Pet.

	query := ToSlice(
		Join(
			func(person Person) Person { return person },
			func(pet Pet) Person { return pet.Owner },
			func(person Person, pet Pet) string { return fmt.Sprintf("%s - %s", person.Name, pet.Name) },
		)(FromSlice(people), FromSlice(pets)),
	)

	assert.DeepEqual(t, query, []string{
		"Hedlund, Magnus - Daisy",
		"Adams, Terry - Barley",
		"Adams, Terry - Boots",
		"Weiss, Charlotte - Whiskers",
	})

	//Output:
	//Hedlund, Magnus - Daisy
	//Adams, Terry - Barley
	//Adams, Terry - Boots
	//Weiss, Charlotte - Whiskers
}

// The following code example demonstrates how to use LastWithT
// to return the last element of an array.
func TestExampleQuery_LastWithT(t *testing.T) {
	numbers := []int{9, 34, 65, 92, 87, 435, 3, 54,
		83, 23, 87, 67, 12, 19}

	last, _ := LastWith(func(n int) bool { return n > 80 })(FromSlice(numbers))

	assert.DeepEqual(t, last, 87)

	//Output:
	//87

}

// The following code example demonstrates how to use OrderByDescendingT
// to order an slice.
func TestExampleQuery_OrderByDescendingT(t *testing.T) {
	type Player struct {
		Name   string
		Points int64
	}

	players := []Player{
		{Name: "Hugo", Points: 4757},
		{Name: "Rick", Points: 7365},
		{Name: "Michael", Points: 2857},
		{Name: "Fadi", Points: 85897},
		{Name: "Peter", Points: 48576},
	}

	//Order and get the top 3 players

	top3Players := ToSlice(
		Select(
			func(kv KeyValue[int64, Player]) string {
				return fmt.Sprintf(
					"Rank: #%d - Player: %s - Points: %d",
					kv.Key,
					kv.Value.Name,
					kv.Value.Points,
				)
			},
		)(

			SelectIndexed(
				func(i int, p Player) KeyValue[int64, Player] {
					return KeyValue[int64, Player]{Key: int64(i + 1), Value: p}
				},
			)(Take(
				OrderByDescending(
					generics.NumericCompare[int64],
					func(p Player) int64 { return p.Points },
				)(FromSlice(players)).Query,
				3,
			))),
	)
	assert.DeepEqual(t, top3Players, []string{
		"Rank: #1 - Player: Fadi - Points: 85897",
		"Rank: #2 - Player: Peter - Points: 48576",
		"Rank: #3 - Player: Rick - Points: 7365",
	})

	// Output:
	// Rank: #1 - Player: Fadi - Points: 85897
	// Rank: #2 - Player: Peter - Points: 48576
	// Rank: #3 - Player: Rick - Points: 7365
}

// The following code example demonstrates how to use OrderByT
// to sort the elements of a slice.
func TestExampleQuery_OrderByT(t *testing.T) {
	type Pet struct {
		Name string
		Age  int
	}
	// Create a list of pets.
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
		{Name: "Daisy", Age: 4},
	}

	orderedPets := ToSlice(
		Select(
			func(pet Pet) string {
				return fmt.Sprintf("%s - %d", pet.Name, pet.Age)

			},
		)(
			OrderBy(
				generics.NumericCompare[int],
				func(pet Pet) int { return pet.Age },
			)(
				FromSlice(pets),
			).Query,
		),
	)

	assert.DeepEqual(t, orderedPets, []string{
		"Whiskers - 1",
		"Boots - 4",
		"Daisy - 4",
		"Barley - 8",
	})

	// Output:
	// Whiskers - 1
	// Boots - 4
	// Daisy - 4
	// Barley - 8
}

// The following code example demonstrates how to use SelectT
// to project over a slice.
func TestExampleQuery_SelectT(t *testing.T) {
	squares := ToSlice(
		Select(
			func(x int) int { return x * x },
		)(Range(1, 10)),
	)
	assert.DeepEqual(t, squares, []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100})

	// Output:
	// [1 4 9 16 25 36 49 64 81 100]
}

// The following code example demonstrates how to use SelectIndexedT
// to determine if the value in a slice of int match their position
// in the slice.
func TestExampleQuery_SelectIndexedT(t *testing.T) {
	numbers := []int{5, 4, 1, 3, 9, 8, 6, 7, 2, 0}

	numsInPlace := ToSlice(
		Select(
			func(kv KeyValue[int, bool]) string {
				return fmt.Sprintf("%d: %t", kv.Key, kv.Value)
			},
		)(
			SelectIndexed(
				func(index, num int) KeyValue[int, bool] { return KeyValue[int, bool]{Key: num, Value: (num == index)} },
			)(
				FromSlice(numbers),
			),
		),
	)
	assert.DeepEqual(t, numsInPlace, []string{
		"5: false",
		"4: false",
		"1: false",
		"3: true",
		"9: false",
		"8: false",
		"6: true",
		"7: true",
		"2: false",
		"0: false",
	})

	// Output:
	// Number: In-place?
	// 5: false
	// 4: false
	// 1: false
	// 3: true
	// 9: false
	// 8: false
	// 6: true
	// 7: true
	// 2: false
	// 0: false

}

// The following code example demonstrates how to use SelectManyT
// to perform a one-to-many projection over a slice
func TestExampleQuery_SelectManyByT(t *testing.T) {

	type Pet struct {
		Name string
	}

	type Person struct {
		Name string
		Pets []Pet
	}

	magnus := Person{
		Name: "Hedlund, Magnus",
		Pets: []Pet{{Name: "Daisy"}},
	}

	terry := Person{
		Name: "Adams, Terry",
		Pets: []Pet{{Name: "Barley"}, {Name: "Boots"}},
	}
	charlotte := Person{
		Name: "Weiss, Charlotte",
		Pets: []Pet{{Name: "Whiskers"}},
	}

	people := []Person{magnus, terry, charlotte}
	results := ToSlice(
		SelectManyBy(
			func(person Person) Query[Pet] { return FromSlice(person.Pets) },
			func(pet Pet, person Person) string {
				return fmt.Sprintf("Owner: %s, Pet: %s", person.Name, pet.Name)
			},
		)(
			FromSlice(people),
		))

	assert.DeepEqual(t, results, []string{
		"Owner: Hedlund, Magnus, Pet: Daisy",
		"Owner: Adams, Terry, Pet: Barley",
		"Owner: Adams, Terry, Pet: Boots",
		"Owner: Weiss, Charlotte, Pet: Whiskers",
	})

	// Output:
	// Owner: Hedlund, Magnus, Pet: Daisy
	// Owner: Adams, Terry, Pet: Barley
	// Owner: Adams, Terry, Pet: Boots
	// Owner: Weiss, Charlotte, Pet: Whiskers
}

// The following code example demonstrates how to use SelectManyT
// to perform a projection over a list of sentences and rank the
// top 5 most used words
func TestExampleQuery_SelectManyT(t *testing.T) {
	sentences := []string{
		"the quick brown fox jumps over the lazy dog",
		"pack my box with five dozen liquor jugs",
		"several fabulous dixieland jazz groups played with quick tempo",
		"back in my quaint garden jaunty zinnias vie with flaunting phlox",
		"five or six big jet planes zoomed quickly by the new tower",
		"I quickly explained that many big jobs involve few hazards",
		"The wizard quickly jinxed the gnomes before they vaporized",
	}

	results := ToSlice(
		SelectIndexed(
			func(index int, wordGroup Group[string, string]) string {
				return fmt.Sprintf("Rank: #%d, Word: %s, Counts: %d", index+1, wordGroup.Key, len(wordGroup.Group))
			})(
			Take(ThenBy(
				strings.Compare,
				func(wordGroup Group[string, string]) string {
					return wordGroup.Key
				})(OrderByDescending(
				generics.NumericCompare[int],
				func(wordGroup Group[string, string]) int {
					return len(wordGroup.Group)
				})(GroupBy(
				func(word string) string { return word },
				func(word string) string { return word },
			)(
				SelectMany(
					func(sentence string) Query[string] {
						return FromSlice(strings.Split(sentence, " "))
					})(FromSlice(sentences))))).Query, 5),
		),
	)

	assert.DeepEqual(t, results, []string{
		"Rank: #1, Word: the, Counts: 4",
		"Rank: #2, Word: quickly, Counts: 3",
		"Rank: #3, Word: with, Counts: 3",
		"Rank: #4, Word: big, Counts: 2",
		"Rank: #5, Word: five, Counts: 2",
	})

	// Output:
	// Rank: #1, Word: the, Counts: 4
	// Rank: #2, Word: quickly, Counts: 3
	// Rank: #3, Word: with, Counts: 3
	// Rank: #4, Word: big, Counts: 2
	// Rank: #5, Word: five, Counts: 2
}

// The following code example demonstrates how to use SelectManyIndexedT
// to perform a one-to-many projection over an slice of log files and
// print out their contents.
func TestExampleQuery_SelectManyIndexedT(t *testing.T) {
	type LogFile struct {
		Name  string
		Lines []string
	}

	file1 := LogFile{
		Name: "file1.log",
		Lines: []string{
			"INFO: 2013/11/05 18:11:01 main.go:44: Special Information",
			"WARNING: 2013/11/05 18:11:01 main.go:45: There is something you need to know about",
			"ERROR: 2013/11/05 18:11:01 main.go:46: Something has failed",
		},
	}

	file2 := LogFile{
		Name: "file2.log",
		Lines: []string{
			"INFO: 2013/11/05 18:11:01 main.go:46: Everything is ok",
		},
	}

	file3 := LogFile{
		Name: "file3.log",
		Lines: []string{
			"2013/11/05 18:42:26 Hello World",
		},
	}

	logFiles := []LogFile{file1, file2, file3}
	results := ToSlice(SelectManyIndexed(
		func(fileIndex int, file LogFile) Query[string] {
			return SelectIndexed(func(lineIndex int, line string) string {
				return fmt.Sprintf("File:[%d] - %s => line: %d - %s", fileIndex+1, file.Name, lineIndex+1, line)
			})(FromSlice(file.Lines))
		})(FromSlice(logFiles)))

	assert.DeepEqual(t, results, []string{
		"File:[1] - file1.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:44: Special Information",
		"File:[1] - file1.log => line: 2 - WARNING: 2013/11/05 18:11:01 main.go:45: There is something you need to know about",
		"File:[1] - file1.log => line: 3 - ERROR: 2013/11/05 18:11:01 main.go:46: Something has failed",
		"File:[2] - file2.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:46: Everything is ok",
		"File:[3] - file3.log => line: 1 - 2013/11/05 18:42:26 Hello World",
	})
	// Output:
	// File:[1] - file1.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:44: Special Information
	// File:[1] - file1.log => line: 2 - WARNING: 2013/11/05 18:11:01 main.go:45: There is something you need to know about
	// File:[1] - file1.log => line: 3 - ERROR: 2013/11/05 18:11:01 main.go:46: Something has failed
	// File:[2] - file2.log => line: 1 - INFO: 2013/11/05 18:11:01 main.go:46: Everything is ok
	// File:[3] - file3.log => line: 1 - 2013/11/05 18:42:26 Hello World

}

// The following code example demonstrates how to use SelectManyByIndexedT
// to perform a one-to-many projection over an array and use the index of
// each outer element.
func TestExampleQuery_SelectManyByIndexedT(t *testing.T) {
	type Pet struct {
		Name string
	}

	type Person struct {
		Name string
		Pets []Pet
	}

	magnus := Person{
		Name: "Hedlund, Magnus",
		Pets: []Pet{{Name: "Daisy"}},
	}

	terry := Person{
		Name: "Adams, Terry",
		Pets: []Pet{{Name: "Barley"}, {Name: "Boots"}},
	}
	charlotte := Person{
		Name: "Weiss, Charlotte",
		Pets: []Pet{{Name: "Whiskers"}},
	}

	people := []Person{magnus, terry, charlotte}

	results := ToSlice(
		SelectManyByIndexed(
			func(index int, person Person) Query[string] {
				return Select(
					func(pet Pet) string {
						return fmt.Sprintf("%d - %s", index, pet.Name)
					},
				)(FromSlice(person.Pets))

			},
			func(indexedPet string, person Person) string {
				return fmt.Sprintf("Pet: %s, Owner: %s", indexedPet, person.Name)
			},
		)(FromSlice(people)),
	)
	assert.DeepEqual(t, results, []string{
		"Pet: 0 - Daisy, Owner: Hedlund, Magnus",
		"Pet: 1 - Barley, Owner: Adams, Terry",
		"Pet: 1 - Boots, Owner: Adams, Terry",
		"Pet: 2 - Whiskers, Owner: Weiss, Charlotte",
	})

	// Output:
	// Pet: 0 - Daisy, Owner: Hedlund, Magnus
	// Pet: 1 - Barley, Owner: Adams, Terry
	// Pet: 1 - Boots, Owner: Adams, Terry
	// Pet: 2 - Whiskers, Owner: Weiss, Charlotte

}

// The following code example demonstrates how to use SingleWithT
// to select the only element of a slice that satisfies a condition.
func TestExampleQuery_SingleWithT(t *testing.T) {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}
	fruit, _ := SingleWith(
		func(f string) bool { return len(f) > 10 },
	)(FromSlice(fruits))
	assert.DeepEqual(t, fruit, "passionfruit")
	// Output:
	// passionfruit
}

// The following code example demonstrates how to use SkipWhileT
// to skip elements of an array as long as a condition is true.
func TestExampleQuery_SkipWhileT(t *testing.T) {
	grades := []int{59, 82, 70, 56, 92, 98, 85}
	lowerGrades := ToSlice(
		SkipWhile(
			func(g int) bool { return g >= 80 })(
			OrderByDescending(generics.NumericCompare[int], func(g int) int { return g })(FromSlice(grades)).Query),
	)

	//"All grades below 80:
	assert.DeepEqual(t, lowerGrades, []int{70, 59, 56})
	// Output:
	// [70 59 56]
}

// The following code example demonstrates how to use SkipWhileIndexedT
// to skip elements of an array as long as a condition that depends
// on the element's index is true.
func TestExampleQuery_SkipWhileIndexedT(t *testing.T) {
	amounts := []int{5000, 2500, 9000, 8000, 6500, 4000, 1500, 5500}

	query := ToSlice(
		SkipWhileIndexed(func(index int, amount int) bool { return amount > index*1000 })(
			FromSlice(amounts),
		),
	)
	assert.DeepEqual(t, query, []int{4000, 1500, 5500})
	// Output:
	// [4000 1500 5500]

}

// The following code example demonstrates how to use SortT
// to order elements of an slice.
func TestExampleQuery_SortT(t *testing.T) {
	type Pet struct {
		Name string
		Age  int
	}
	// Create a list of pets.
	pets := []Pet{
		{Name: "Barley", Age: 8},
		{Name: "Boots", Age: 4},
		{Name: "Whiskers", Age: 1},
		{Name: "Daisy", Age: 4},
	}

	orderedPets := ToSlice(
		Select(
			func(pet Pet) string {
				return fmt.Sprintf("%s - %d", pet.Name, pet.Age)
			},
		)(
			Sort(func(pet1 Pet, pet2 Pet) bool { return pet1.Age < pet2.Age })(FromSlice(pets)),
		),
	)

	assert.DeepEqual(t, orderedPets, []string{
		"Whiskers - 1",
		"Boots - 4",
		"Daisy - 4",
		"Barley - 8",
	})
	// Output:
	// Whiskers - 1
	// Boots - 4
	// Boots - 4
	// Barley - 8

}

// The following code example demonstrates how to use TakeWhileT
// to return elements from the start of a slice.
func TestExampleQuery_TakeWhileT(t *testing.T) {
	fruits := []string{"apple", "banana", "mango", "orange", "passionfruit", "grape"}

	query := ToSlice(
		TakeWhile(
			func(fruit string) bool { return fruit != "orange" },
		)(FromSlice(fruits)),
	)
	assert.DeepEqual(t, query, []string{"apple", "banana", "mango"})
	// Output:
	// [apple banana mango]
}

// The following code example demonstrates how to use TakeWhileIndexedT
// to return elements from the start of a slice as long asa condition
// that uses the element's index is true.
func TestExampleQuery_TakeWhileIndexedT(t *testing.T) {

	fruits := []string{"apple", "passionfruit", "banana", "mango",
		"orange", "blueberry", "grape", "strawberry"}

	query := ToSlice(
		TakeWhileIndexed(
			func(index int, fruit string) bool { return len(fruit) >= index },
		)(FromSlice(fruits)),
	)
	assert.DeepEqual(t, query, []string{"apple", "passionfruit", "banana", "mango", "orange", "blueberry"})
	// Output:
	// [apple passionfruit banana mango orange blueberry]
}

// The following code example demonstrates how to use ToMapBy
// by using a key and value selectors to populate a map.
func TestExampleQuery_ToMapByT(t *testing.T) {
	type Product struct {
		Name string
		Code int
	}

	products := []Product{
		{Name: "orange", Code: 4},
		{Name: "apple", Code: 9},
		{Name: "lemon", Code: 12},
		{Name: "apple", Code: 9},
	}

	map1 := ToMapBy(
		func(item Product) int { return item.Code },
		func(item Product) string { return item.Name },
	)(FromSlice(products))
	assert.DeepEqual(t, map1, map[int]string{
		4:  "orange",
		9:  "apple",
		12: "lemon",
	})

	// Output:
	// orange
	// apple
	// lemon
}

// The following code example demonstrates how to use WhereT
// to filter a slices.
func TestExampleQuery_WhereT(t *testing.T) {
	fruits := []string{"apple", "passionfruit", "banana", "mango",
		"orange", "blueberry", "grape", "strawberry"}
	query := ToSlice(Where(func(f string) bool {
		return len(f) > 6
	})(FromSlice(fruits)))

	assert.DeepEqual(t, []string{"passionfruit", "blueberry", "strawberry"}, query)
	// Output:
	// [passionfruit blueberry strawberry]
}

// The following code example demonstrates how to use WhereIndexedT
// to filter a slice based on a predicate that involves the index of each element.
func TestExampleQuery_WhereIndexedT(t *testing.T) {
	numbers := []int{0, 30, 20, 15, 90, 85, 40, 75}

	query := ToSlice(WhereIndexed(func(index int, number int) bool { return number <= index*10 })(FromSlice(numbers)))

	assert.DeepEqual(t, []int{0, 20, 15, 40}, query)
	// Output:
	// [0 20 15 40]
}

// The following code example demonstrates how to use the ZipT
// method to merge two slices.
func TestExampleQuery_ZipT(t *testing.T) {
	number := []int{1, 2, 3, 4, 5}
	words := []string{"one", "two", "three"}

	s := Results(Zip(
		func(a int, b string) []any { return []any{a, b} },
	)(
		FromSlice(number),
		FromSlice(words),
	))
	assert.DeepEqual(t, s, [][]any{
		{1, "one"},
		{2, "two"},
		{3, "three"},
	})
	// Output:
	// [[1 one] [2 two] [3 three]]
}

// The following code example demonstrates how to use the FromChannelT
// to make a Query from typed channel.
func TestExampleFromChannelT(t *testing.T) {
	ch := make(chan string, 3)
	ch <- "one"
	ch <- "two"
	ch <- "three"
	close(ch)

	q := FromChannel(ch)

	assert.DeepEqual(t, Results(q), []string{"one", "two", "three"})
	// Output:
	// [one two three]
}

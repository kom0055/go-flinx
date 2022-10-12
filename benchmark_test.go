package flinx

import (
	"testing"

	"github.com/ahmetb/go-linq/v3"
)

const (
	size = 10000000
)

func BenchmarkSelectWhereFirst(b *testing.B) {

	b.Run("BenchmarkSelectWhereFirst_flinx", func(b *testing.B) {
		selectFn := Select(func(i int) int {
			return -i
		})
		whereFn := Where(func(i int) bool {
			return i > -5
		})

		for n := 0; n < b.N; n++ {
			First(whereFn(selectFn(Range(1, size))))
		}
	})
	b.Run("BenchmarkSelectWhereFirst_linq", func(b *testing.B) {

		for n := 0; n < b.N; n++ {
			linq.Range(1, size).Select(func(i interface{}) interface{} {
				return -i.(int)
			}).Where(func(i interface{}) bool {
				return i.(int) > -1000
			}).First()
		}
	})

	b.Run("BenchmarkSelectWhereFirst_generics", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			linq.Range(1, size).SelectT(func(i int) int {
				return -i
			}).WhereT(func(i int) bool {
				return i > -1000
			}).First()
		}
	})

}

func BenchmarkSum(b *testing.B) {
	b.Run("BenchmarkSum_flinx", func(b *testing.B) {
		whereFn := Where(func(i int) bool {
			return i%2 == 0
		})
		for n := 0; n < b.N; n++ {
			Sum(whereFn(Range(1, size)))
		}
	})

	b.Run("BenchmarkSum_linq", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			linq.Range(1, size).Where(func(i interface{}) bool {
				return i.(int)%2 == 0
			}).SumInts()
		}
	})

	b.Run("BenchmarkSum_generics", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			linq.Range(1, size).WhereT(func(i int) bool {
				return i%2 == 0
			}).SumInts()
		}
	})
}

func BenchmarkZipSkipTake(b *testing.B) {
	b.Run("BenchmarkZipSkipTake_flinx", func(b *testing.B) {

		selectFn := Select(func(i int) int {
			return i * 2
		})

		zipFn := Zip(func(i, j int) int {
			return i + j
		})

		for n := 0; n < b.N; n++ {

			Take(Skip(zipFn(Range(1, size), selectFn(Range(1, size))), 2), 5)
		}
	})

	b.Run("BenchmarkZipSkipTake_linq", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			linq.Range(1, size).Zip(linq.Range(1, size).Select(func(i interface{}) interface{} {
				return i.(int) * 2
			}), func(i, j interface{}) interface{} {
				return i.(int) + j.(int)
			}).Skip(2).Take(5)
		}
	})

	b.Run("BenchmarkZipSkipTake_generics", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			linq.Range(1, size).ZipT(linq.Range(1, size).SelectT(func(i int) int {
				return i * 2
			}), func(i, j int) int {
				return i + j
			}).Skip(2).Take(5)
		}
	})
}

func BenchmarkFromChannel(b *testing.B) {
	b.Run("BenchmarkFromChannel_flinx", func(b *testing.B) {
		allFn := All(func(i int) bool {
			return true
		})

		for n := 0; n < b.N; n++ {
			ch := make(chan int)
			go func() {
				for i := 0; i < size; i++ {
					ch <- i
				}

				close(ch)
			}()
			allFn(FromChannel(ch))
		}
	})

	b.Run("BenchmarkFromChannel_linq", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ch := make(chan interface{})
			go func() {
				for i := 0; i < size; i++ {
					ch <- i
				}

				close(ch)
			}()

			linq.FromChannel(ch).All(func(i interface{}) bool { return true })
		}
	})

	b.Run("BenchmarkFromChannel_linqt", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			ch := make(chan interface{})
			go func() {
				for i := 0; i < size; i++ {
					ch <- i
				}

				close(ch)
			}()

			linq.FromChannelT(ch).All(func(i interface{}) bool { return true })
		}
	})
}

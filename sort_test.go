package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("very small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5, 5)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("very big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 500000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays, elements in seria", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceSeria(100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays, elements reversed", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceReverse(100000)
			b.StartTimer()
			bubbleSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkSelectionSort(b *testing.B) {

	b.Run("very small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5, 5)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("very big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 500000)
			b.StartTimer()
			selectionSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkInsertionSort(b *testing.B) {
	b.Run("very small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5, 5)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("very big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 500000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays, elements in seria", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceSeria(100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays, elements reversed", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSliceReverse(100000)
			b.StartTimer()
			insertionSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkMergeSort(b *testing.B) {
	b.Run("very small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5, 5)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("very big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 500000)
			b.StartTimer()
			MergeSort(ar)
			b.StopTimer()
		}
	})
}

func BenchmarkQuickSort(b *testing.B) {
	b.Run("very small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(5, 5)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("very big arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100000, 500000)
			b.StartTimer()
			quickSort(ar)
			b.StopTimer()
		}
	})
}

//Массив с рандомными значениями
func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}
	return ar
}

//Массив со значениями по порядку,
//лучший случай для сортировки пузырьком и вставкой
func generateSliceSeria(size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = i
	}
	return ar
}

//Массив со значениями в обратном порядке,
//худший случай для сортировки пузырьком и вставкой
func generateSliceReverse(size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = size
		size--
	}
	return ar
}

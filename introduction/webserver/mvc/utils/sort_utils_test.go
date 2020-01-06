package utils

import (
	"math/rand"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	ints := []int{5, 4, 3, 2, 1, 0}
	BubbleSort(ints)
	for i := 0; i < len(ints); i++ {
		if i != ints[i] {
			t.Error()
		}
	}
}

func TestBubbleSortBestCase(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4, 5}
	BubbleSort(ints)
	for i := 0; i < len(ints); i++ {
		if i != ints[i] {
			t.Error()
		}
	}
}

func makeIntSlice(num int) []int {
	ints := make([]int, num)
	for i := 0; i < num; i++ {
		ints[i] = rand.Intn(1000)
	}
	return ints
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(makeIntSlice(1000))
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(makeIntSlice(1000))
	}
}

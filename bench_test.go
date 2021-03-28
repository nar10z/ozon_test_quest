package ozon

import (
	"testing"
)

func BenchmarkGetClustersQuantity(b *testing.B) {
	b.SetBytes(2)

	grid := [][]int{
		{0, 1, 1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 1, 0},
	}

	for i := 0; i < b.N; i++ {
		GetClustersQuantity(grid)
	}
}

func BenchmarkGetClustersQuantity2(b *testing.B) {
	b.SetBytes(2)

	grid := [][]int{
		{0, 1, 1, 0, 0, 0, 1},
		{0, 1, 1, 1, 0, 0, 1},
		{1, 0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 1, 1, 0},
	}

	for i := 0; i < b.N; i++ {
		GetClustersQuantity2(grid)
	}
}

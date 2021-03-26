package ozon

import "testing"

func TestGetClustersQuantity(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test #1",
			args{
				grid: [][]int{
					{0, 1, 1, 0, 0, 0, 1},
					{0, 1, 1, 1, 0, 0, 1},
					{1, 0, 1, 0, 0, 0, 0},
					{1, 1, 1, 0, 1, 1, 0},
				},
			},
			3,
		},
		{
			"Test #2",
			args{
				grid: [][]int{
					{0, 1, 1, 0, 0, 0, 1},
					{0, 1, 1, 1, 0, 0, 1},
					{1, 1, 1, 0, 1, 1, 0},
					{1, 0, 1, 0, 0, 0, 0},
					{1, 0, 1, 0, 0, 0, 1},
				},
			},
			4,
		},
		{
			"Test #3",
			args{
				grid: [][]int{
					{0, 1, 1, 0, 0, 0, 1},
					{1, 0, 1, 0, 1, 1, 0},
					{1, 0, 1, 0, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 1},
					{1, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 1, 1},
				},
			},
			8,
		},
		{
			"Test #4",
			args{
				grid: [][]int{
					{0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1},
				},
			},
			5,
		},
		{
			"Test #5",
			args{
				grid: [][]int{
					{0, 1, 1, 0, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1},
					{0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1},
					{0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1},
				},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetClustersQuantity(tt.args.grid); got != tt.want {
				t.Errorf("GetClustersQuantity() = %v, want %v", got, tt.want)
			}
		})
	}
}

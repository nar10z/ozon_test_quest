package ozon

import "math"

type gridInformation2 struct {
	grid    [][]int
	visited [][]bool
}

func GetClustersQuantity2(grid [][]int) int {
	l := len(grid)
	v := make([][]bool, l)

	for i, ints := range grid {
		v[i] = make([]bool, len(ints))
	}

	g := gridInformation2{
		grid:    grid,
		visited: v,
	}

	return g.calc()
}

func (gi *gridInformation2) calc() int {
	var qClusters int

	for i, ints := range gi.grid {
		for j, v := range ints {
			if gi.checkItem(v, i, j) {
				qClusters += 1
			}
		}
	}

	return qClusters
}

func (gi *gridInformation2) checkItem(v int, y, x int) bool {
	if v == 0 || gi.visited[y][x] {
		return false
	}

	gi.visited[y][x] = true

	left := gi.getGridItem(y, x-1)
	top := gi.getGridItem(y-1, x)
	right := gi.getGridItem(y, x+1)
	bottom := gi.getGridItem(y+1, x)

	if left != math.MinInt32 {
		gi.checkItem(left, y, x-1)
	}
	if top != math.MinInt32 {
		gi.checkItem(top, y-1, x)
	}
	if right != math.MinInt32 {
		gi.checkItem(right, y, x+1)
	}
	if bottom != math.MinInt32 {
		gi.checkItem(bottom, y+1, x)
	}

	return true
}

func (gi *gridInformation2) getGridItem(y, x int) int {
	if y < 0 || x < 0 {
		return math.MinInt32
	}

	if y >= len(gi.grid) {
		return math.MinInt32
	}
	if x >= len(gi.grid[y]) {
		return math.MinInt32
	}

	item := gi.grid[y][x]
	if item == 0 {
		return math.MinInt32
	}

	return item
}

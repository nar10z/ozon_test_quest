package ozon

type gridItem struct {
	x int
	y int

	value   int
	visited bool
}

type gridInformation struct {
	grid          [][]int
	convertedGrid [][]*gridItem

	qClusters int
}

//GetClustersQuantity returned quantity cluster in matrix
func GetClustersQuantity(grid [][]int) int {
	g := gridInformation{
		grid: grid,
	}

	g.convert()
	g.calc()

	return g.qClusters
}

//convert - converts the incoming grid to coords struct
func (gi *gridInformation) convert() {
	for i, v := range gi.grid {
		var g []*gridItem

		for j, v2 := range v {
			g = append(g, &gridItem{
				value: v2,
				x:     i,
				y:     j,
			})
		}

		gi.convertedGrid = append(gi.convertedGrid, g)
	}
}

//calc - calculate quantity clusters
func (gi *gridInformation) calc() {
	for _, items := range gi.convertedGrid {
		for _, item := range items {
			if gi.checkItem(item) {
				gi.qClusters += 1
			}
		}
	}
}

//checkItem - checking item in links in cluster
func (gi *gridInformation) checkItem(item *gridItem) bool {
	if item.value == 0 || item.visited {
		return false
	}

	item.visited = true

	left := gi.getGridItem(item.x, item.y-1)
	top := gi.getGridItem(item.x-1, item.y)
	right := gi.getGridItem(item.x, item.y+1)
	bottom := gi.getGridItem(item.x+1, item.y)

	if left != nil {
		gi.checkItem(left)
	}
	if top != nil {
		gi.checkItem(top)
	}
	if right != nil {
		gi.checkItem(right)
	}
	if bottom != nil {
		gi.checkItem(bottom)
	}

	return true
}

//getGridItem returned grid item from coords
func (gi *gridInformation) getGridItem(x, y int) *gridItem {
	if x < 0 || y < 0 {
		return nil
	}

	if x >= len(gi.convertedGrid) {
		return nil
	}
	if y >= len(gi.convertedGrid[x]) {
		return nil
	}

	item := gi.convertedGrid[x][y]
	if item.value == 0 {
		return nil
	}

	return item
}

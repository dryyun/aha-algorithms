package main

import "fmt"

var road [5][5]int
var markRoad [5]int
var short, tmpShort int

func main() {
	short = 99999
	road = [5][5]int{
		{0, 2, -1, -1, 10},
		{-1, 0, 3, -1, 7},
		{4, -1, 0, 4, -1},
		{-1, -1, -1, 0, 5},
		{-1, -1, 3, -1, 0},
	}
	// 寻找 1 -> 5 的最短路径，在数组里就是 0 -> 4 的最短路径
	// 深度优先搜索算法

	markRoad[0] = 1
	dfs52(0)

	fmt.Println(short)
}

func dfs52(city int) {
	if city == 4 {
		fmt.Println(markRoad)
		if tmpShort < short {
			short = tmpShort
		}
		return
	}
	for k, v := range road[city] {
		if v > 0 && markRoad[k] == 0 {
			markRoad[k] = 1
			tmpShort += road[city][k]
			dfs52(k)
			markRoad[k] = 0
			tmpShort -= road[city][k]
		}
	}
	return
}

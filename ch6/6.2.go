package main

import "fmt"

func main() {
	const MAX = 999999

	showRoad := func(road [][]int) {
		for _, row := range road {
			for _, col := range row {
				fmt.Printf("%8d", col)
			}
			fmt.Println()
		}
	}

	road := [][]int{
		{0, 1, 12, MAX, MAX, MAX},
		{MAX, 0, 9, 3, MAX, MAX},
		{MAX, MAX, 0, MAX, 5, MAX},
		{MAX, MAX, 4, 0, 13, 15},
		{MAX, MAX, MAX, MAX, 0, 4},
		{MAX, MAX, MAX, MAX, MAX, 0},
	}
	fmt.Println("Origin Road is ...")
	showRoad(road)

	resultRoad := road
	// 1 号顶点到 其余顶点的初始路程，在数组中就是 0 号
	for f := 0; f < len(road); f++ {

		find := f
		dis := [6]int{0}
		for k, v := range road[find] {
			dis[k] = v
		}
		n := len(road[find])

		book := [6]int{0}

		min := MAX
		tmp := 0 //tmp book id
		// 算法核心
		for i := 0; i < n-1; i++ { //运行次数，这里是 5 次
			// 找到离 0 最近的顶点
			min = MAX
			for j := 0; j < n; j++ {
				if book[j] == 0 && dis[j] < min {
					min = dis[j]
					tmp = j
				}
			}
			book[tmp] = 1
			for v := 0; v < n; v++ {
				if road[tmp][v] < MAX {
					if dis[v] > dis[tmp]+road[tmp][v] {
						dis[v] = dis[tmp] + road[tmp][v]
					}
				}
			}
		}
		resultRoad[f] = dis[:]
	}
	fmt.Println("Result Road is ...")
	showRoad(resultRoad)

}

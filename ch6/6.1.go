package main

import "fmt"

const MAX = 999999

func main() {
	road := [][]int{
		{0, 2, 6, 4},
		{MAX, 0, 3, MAX},
		{7, MAX, 0, 1},
		{5, MAX, 12, 0},
	}
	fmt.Println("Origin Road is ...")
	showRoad(road)

	n := 4
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if road[i][j] > road[i][k]+road[k][j] {
					road[i][j] = road[i][k] + road[k][j]
				}
			}
		}
	}
	fmt.Println("Result Road is ...")
	showRoad(road)

}

func showRoad(road [][]int) {
	for _, row := range road {
		for _, col := range row {
			fmt.Printf("%8d", col)
		}
		fmt.Println()
	}
}
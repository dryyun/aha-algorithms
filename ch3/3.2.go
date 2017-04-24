package main

import "fmt"

func main() {
	xo := 13
	yo := 13
	m := [13][13]string{
		{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", "G", "G", ".", "G", "G", "G", "#", "G", "G", "G", ".", "#"},
		{"#", "#", "#", ".", "#", "G", "#", "G", "#", "G", "#", "G", "#"},
		{"#", ".", ".", ".", ".", ".", ".", ".", "#", ".", ".", "G", "#"},
		{"#", "G", "#", ".", "#", "#", "#", ".", "#", "G", "#", "G", "#"},
		{"#", "G", "G", ".", "G", "G", "G", ".", "#", ".", "G", "G", "#"},
		{"#", "G", "#", ".", "#", "G", "#", ".", "#", ".", "#", "#", "#"},
		{"#", "#", "G", ".", ".", ".", "G", ".", ".", ".", ".", ".", "#"},
		{"#", "G", "#", ".", "#", "G", "#", "#", "#", ".", "#", "G", "#"},
		{"#", ".", ".", ".", "G", "#", "G", "G", "G", ".", "G", "G", "#"},
		{"#", "G", "#", ".", "#", "G", "#", "G", "#", ".", "#", "G", "#"},
		{"#", "G", "G", ".", "G", "G", "G", "#", "G", ".", "G", "G", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
	}

	// 输出地图
	for _, a1 := range m {
		for _, a2 := range a1 {
			fmt.Printf("%s", a2)
		}
		fmt.Println()
	}

	resultSum := 0
	resultX := 0
	resultY := 0
	currentSum := 0
	x := 0
	y := 0
	for i := 1; i < xo-1; i++ {
		for j := 1; j < yo-1; j++ {
			if m[i][j] == "." {
				currentSum = 0 //记数，最多消灭的地人数
				//向上统计
				x = i
				y = j
				for m[x][y] != "#" {
					if m[x][y] == "G" {
						currentSum++
					}
					x--
				}

				//向下统计
				x = i
				y = j
				for m[x][y] != "#" {
					if m[x][y] == "G" {
						currentSum++
					}
					x++
				}

				//向左统计
				x = i
				y = j
				for m[x][y] != "#" {
					if m[x][y] == "G" {
						currentSum++
					}
					y--
				}

				//向右统计
				x = i
				y = j
				for m[x][y] != "#" {
					if m[x][y] == "G" {
						currentSum++
					}
					y++
				}

				if currentSum > resultSum {
					resultSum = currentSum
					resultX = i
					resultY = j
				}
			}
		}
	}
	fmt.Println(resultSum)
	fmt.Printf("x = %d,y = %d\n", resultX, resultY)

	return
}

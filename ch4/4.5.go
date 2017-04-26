package main

import "fmt"

var theMap [10][10]int

var islandNextNode [4][2]int

type islandNode struct {
	x, y int
}

var bookNode [10][10]int
var colorNode [10][10]int

func main() {
	theMap = [10][10]int{
		{1, 2, 1, 0, 0, 0, 0, 0, 2, 3},
		{3, 0, 2, 0, 1, 2, 1, 0, 1, 2},
		{4, 0, 1, 0, 1, 2, 3, 2, 0, 1},
		{3, 2, 0, 0, 0, 1, 2, 4, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 5, 3, 0},
		{0, 1, 2, 1, 0, 1, 5, 4, 3, 0},
		{0, 1, 2, 3, 1, 3, 6, 2, 1, 0},
		{0, 0, 3, 4, 8, 9, 7, 5, 0, 0},
		{0, 0, 0, 3, 7, 8, 6, 0, 1, 2},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
	}

	for _, row := range theMap {
		for _, col := range row {
			if col > 0 {
				fmt.Printf("%3d", 1)
			} else {
				fmt.Printf("%3d", 0)
			}
		}
		fmt.Println()
	}

	islandNextNode = [4][2]int{
		{0, 1},  //右
		{1, 0},  //下
		{0, -1}, //左
		{-1, 0}, //上
	}

	islandQue := [100]islandNode{}

	startX := 5 // 6-1
	startY := 7 // 8-1

	//初始化
	head := 0
	tail := 0
	colorNode[startX][startY] = -1
	bookNode[startX][startY] = 1
	islandQue[tail] = islandNode{startX, startY}
	tail++

	for head < tail {
		for _, v := range islandNextNode {
			nextX := islandQue[head].x + v[0]
			nextY := islandQue[head].y + v[1]
			// 边界判断 && 陆地判断
			if nextX >= 0 && nextY >= 0 && nextX <= 9 && nextY <= 9 && theMap[nextX][nextY] > 0 && bookNode[nextX][nextY] == 0 {
				colorNode[nextX][nextY] = -1
				bookNode[nextX][nextY] = 1
				islandQue[tail] = islandNode{nextX, nextY}
				tail++
			}
		}
		head++
	}
	fmt.Println(tail)
	fmt.Println(islandQue)

	for x, row := range theMap {
		for y, col := range row {
			if colorNode[x][y] != 0 {
				fmt.Printf("%3d", colorNode[x][y])
			} else {
				fmt.Printf("%3d", col)
			}
		}
		fmt.Println()
	}

}

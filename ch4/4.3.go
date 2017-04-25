// 4.3 广度优先
package main

import "fmt"

type node struct {
	x, y, s, f int // s 是步数，f 父亲在 que 中的编号
}

var markNode [5][4]int
var mapNode [5][4]string
var nextNode [4][2]int

func main() {
	nextNode = [4][2]int{
		{0, 1},  //右
		{1, 0},  //下
		{0, -1}, //左
		{-1, 0}, //上
	}

	mapNode = [5][4]string{
		{".", ".", "#", "."}, // . 表示空地，# 表示障碍，G 表示目标人
		{".", ".", ".", "."},
		{".", ".", "#", "."},
		{".", "#", "G", "."},
		{".", ".", ".", "#"},
	}

	var que [2501]node

	head := 1
	tail := 1

	x := 0
	y := 0
	//从 0，0 开始
	que[tail].x = x
	que[tail].y = y
	que[tail].s = 0
	que[tail].f = 0
	tail++

	markNode[x][y] = 1
	flag := 0
	for head < tail {
		for _, nn := range nextNode {
			nex := que[head].x + nn[0]
			ney := que[head].y + nn[1]
			if nex >= 0 && nex <= 4 && ney >= 0 && ney <= 3 && mapNode[nex][ney] != "#" && markNode[nex][ney] == 0 {
				markNode[nex][ney] = 1
				que[tail].x = nex
				que[tail].y = ney
				que[tail].s = que[head].s + 1
				que[tail].f = head
				tail++

				if mapNode[nex][ney] == "G" {
					flag = 1
					break
				}
			}

		}
		if flag == 1 {
			break
		}
		head++ // 一个点结束后，出栈
	}

	if flag == 1 {
		fmt.Println(que[tail-1].s)
		for que[tail-1].f != 0 {
			fmt.Println(que[tail-1].x, que[tail-1].y)
			tail = que[tail-1].f + 1
		}
	} else {
		fmt.Println("No")
	}

}

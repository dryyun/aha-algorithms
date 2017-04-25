package main

import "fmt"

var next43 [4][2]int
var map43 [5][4]string
var mark43 [5][4]int
var step43, result43 [100][2]int
var short, tmpShort int

func main() {
	//坐标从 (0,0) 开始，(4,3) 结束
	// 按照顺时针方向去尝试，右、下、左、上
	next43 = [4][2]int{
		{0, 1},  //右
		{1, 0},  //下
		{0, -1}, //左
		{-1, 0}, //上
	}

	map43 = [5][4]string{
		{".", ".", "#", "."}, // . 表示空地，# 表示障碍，G 表示目标人
		{".", ".", ".", "."},
		{".", ".", "#", "."},
		{".", "#", "G", "."},
		{".", ".", ".", "#"},
	}

	short = 99999
	dfs43(0, 0)
	fmt.Println(short)
	for i := 0; i < short; i++ {
		fmt.Println(result43[i])
	}

}

func dfs43(x, y int) {
	if map43[x][y] == "G" { // 找到了
		if tmpShort < short {
			short = tmpShort
			result43 = step43
		}
		return
	}

	for _, ne := range next43 {
		nex := x + ne[0]
		ney := y + ne[1]
		//边界判断 && 障碍物判断 && 已有路径判断
		if nex >= 0 && nex <= 4 && ney >= 0 && ney <= 3 && map43[nex][ney] != "#" && mark43[nex][ney] == 0 {
			mark43[nex][ney] = 1
			step43[tmpShort] = [2]int{nex, ney}
			tmpShort++
			dfs43(nex, ney)
			//
			mark43[nex][ney] = 0
			tmpShort--
			step43[tmpShort] = [2]int{0, 0}
		}
	}

	return
}

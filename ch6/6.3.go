package main

import "fmt"

func main() {
	const MAX = 999999
	// n 表示顶点个数
	// m 表示边个数
	// x,y,z 三个数组，记录 x->y=z，x 点到 y点的距离是 z
	// dis 数组，记录源点到其他个点的最短路径
	// 对所有的边进行 n-1 次松弛

	n := 5
	m := 5
	x := [6]int{0, 2, 1, 1, 4, 3}
	y := [6]int{0, 3, 2, 5, 5, 4}
	z := [6]int{0, 2, -3, 5, 2, 3}
	dis := [6]int{0, 0, MAX, MAX, MAX, MAX}

	for k := 1; k <= n-1; k++ {
		for i := 1; i <= m; i++ {
			if dis[y[i]] > dis[x[i]]+z[i] {
				dis[y[i]] = dis[x[i]] + z[i]
			}
		}
	}

	//这里的结果过 dis[0] 忽略
	fmt.Println(dis)

}

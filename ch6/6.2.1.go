package main

import "fmt"

func main() {
	// 好的，这个用数组，邻接表 表示的 图比较难理解。。
	var n, m int
	var u, v, w [6]int     // 比 m 大 1
	var first, next [6]int // 比 n 大1

	fmt.Println("Input n , m :")
	fmt.Scanf("%d %d", &n, &m)

	for i := 1; i <= n; i++ {
		first[i] = -1
	}
	for i := 1; i <= m; i++ {
		fmt.Scanf("%d %d %d", &u[i], &v[i], &w[i])
		//下面两句是关键
		next[i] = first[u[i]]
		first[u[i]] = i
	}
	fmt.Println(u, v, w)
	fmt.Println(first, next)

}

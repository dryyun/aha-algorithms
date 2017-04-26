package main

import "fmt"

var e [5][5]int
var mark [5]int
var store []int

func main() {
	e = [5][5]int{
		{0, 1, 1, -1, 1},
		{1, 0, -1, 1, -1},
		{1, -1, 0, -1, 1},
		{-1, 1, -1, 0, -1},
		{1, -1, 1, -1, 0},
	}

	//深度优先搜索

	mark[0] = 1
	store = append(store, 0)
	dfs(0)
	for _, v := range store {
		fmt.Printf("%d ", v+1)
	}

	//广度优先搜索
	/*
		head := 0
		tail := 0
		que := [5]int{0}

		mark[0] = 1
		que[head] = 0
		tail++
		for head < tail {
			for k, v := range e[que[head]] {
				if v == 1 && mark[k] == 0 {
					mark[k] = 1
					que[tail] = k
					tail++
				}
			}
			head++
		}
		for _, v := range que {
			fmt.Printf("%d ", v+1)
		}
	*/
}

// 深度优先搜索
func dfs(cur int) {

	for k, v := range e[cur] {
		if v == 1 && mark[k] == 0 {
			mark[k] = 1
			store = append(store, k)
			dfs(k)
		}
	}
	return
}

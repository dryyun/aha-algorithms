// 4.1 深度优先搜索
package main

import "fmt"

var n, total int
var a, book [10]int

func main() {
	n = 5 // n <=9

	dsf(1)

	fmt.Println(total)
	return

}

func dsf(step int) {
	if step == n+1 {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d", a[i])
		}
		total++
		fmt.Println()
		return
	}
	for i := 1; i <= n; i++ {
		if book[i] == 0 {
			a[step] = i
			book[i] = 1
			dsf(step + 1)
			book[i] = 0 //非常重要的一步
		}
	}

	return
}

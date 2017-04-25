// 4.1 节 实现 3.1 的内容
package main

import "fmt"

var nn ,total31 int
var step, mark [10]int

func main() {
	nn = 9
	dfs31(1)
	fmt.Println(total31/2)
	return
}

func dfs31(s int) {

	if s == 10 {
		if step[1]*100+step[2]*10+step[3]+step[4]*100+step[5]*10+step[6] == step[7]*100+step[8]*10+step[9] {
			fmt.Printf("%d%d%d+%d%d%d=%d%d%d\n", step[1], step[2], step[3], step[4], step[5], step[6], step[7], step[8], step[9])
			total31 ++
		}

		return
	}
	for i := 1; i <= nn; i++ {
		if mark[i] == 0 {
			mark[i] = 1
			step[s] = i
			dfs31(s +1)
			mark[i] = 0
		}
	}
	return

}

package main

import "fmt"

func main() {

	// 建立最大堆，方法还是 siftdown ，是不是这个方法在细节上有点变化
	fmt.Println("建立堆原数据 ...")
	h := []int{0, 99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92}
	fmt.Println(h)
}

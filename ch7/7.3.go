package main

import "fmt"

func main() {
	h := []int{0, 23, 2, 5, 12, 7, 17, 25, 19, 36, 99, 22, 28, 46, 92}

	fmt.Println("Origin ...")
	fmt.Println(h)

	n := len(h) - 1 // 堆的大小，有 n 个元素

	swap := func(x, y int) {
		h[x], h[y] = h[y], h[x]
	}

	// 向下调整
	siftdown := func(i int) { //传入需要向下调整的节点的编号
		var t, flag int             // flag 用来标记是否需要继续向下调整
		for i*2 <= n && flag == 0 { // i*2 <=2 ，判断有没有子节点
			if h[i] > h[i*2] { // 判断左边子节点
				t = i * 2
			} else {
				t = i
			}
			if i*2+1 <= n { // 判断有没有 右边子节点
				if h[t] > h[i*2+1] {
					t = i*2 + 1
				}
			}
			if t != i {
				swap(t, i)
				i = t //更新 i 为刚才交换的节点的编号
			} else {
				flag = 1
			}
		}
	}

	siftdown(1)
	fmt.Println("siftdown result...")
	fmt.Println(h)

	fmt.Println("add 3 ...")
	h = append(h, 3)

	// 向上调整
	siftup := func(i int) {
		flag := 0
		if i == 1 {
			return
		}
		for i != 1 && flag == 0 {
			if h[i] < h[i/2] {
				swap(i, i/2)
				i = i / 2
			} else {
				flag = 1
			}
		}
	}
	siftup(len(h) - 1)
	fmt.Println("siftup result...")
	fmt.Println(h)

	fmt.Println("建立堆原数据 ...")
	h = []int{0, 99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92}
	fmt.Println(h)

	create := func(n int) { // n 元素个数
		for i := n / 2; i >= 1; i-- {
			siftdown(i)
		}
	}
	create(len(h) - 1)

	fmt.Println("最小堆结果 ...")
	fmt.Println(h)

	//堆排序
	//从小到大排序，可以先建立最小堆，然后每次删除顶部元素，就是排序之后的结果

	fmt.Println("排序 ...")
	fmt.Println(h)
	n = len(h) - 1
	deletemax := func() int { //删除最大元素，返回最小元素，// 其实并没有删除最大的
		t := h[1]
		h[1] = h[n]
		n--
		siftdown(1)
		return t
	}
	num := n
	fmt.Println("排序结果 ...")
	for i := 1; i <= num; i++ {
		fmt.Printf("%d ", deletemax())
	}

}

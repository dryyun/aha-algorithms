package main

import "fmt"

type queue struct {
	data [100]int
	head int
	tail int
}

func main() {
	arr := []int{6, 3, 1, 7, 5, 8, 9, 2, 4}

	q := queue{head: 0, tail: 0}
	for k, v := range arr {
		q.data[k] = v
		q.tail++
	}

	for q.head < q.tail { //当队列不为空
		fmt.Printf("%d ", q.data[q.head])
		q.head++
		q.data[q.tail] = q.data[q.head]
		q.tail++
		q.head++
	}
	return
}

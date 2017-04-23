package main

import "fmt"

type listNode struct {
	data int
	next *listNode
}

func main() {
	var head, current, tmp, loop *listNode

	arr := []int{2, 3, 5, 8, 9, 10, 18, 26, 32}

	for _, v := range arr {
		tmp = new(listNode)
		tmp.data = v

		if head == nil {
			head = tmp
			current = tmp
		} else {
			current.next = tmp
			current = tmp
		}
	}

	num := 2
	loop = head

	if loop != nil && loop.data > num { // num:=1
		tmp = &listNode{data: num, next: head}
		head = tmp
	} else { // num:=6 or num :=36

		for loop != nil {
			if loop.next == nil || loop.next.data > num {
				tmp = &listNode{data: num, next: loop.next}
				loop.next = tmp
				break
			}
			loop = loop.next
		}
	}

	//输出 list
	loop = head
	for loop != nil {
		fmt.Println(loop.data)
		loop = loop.next
	}

}

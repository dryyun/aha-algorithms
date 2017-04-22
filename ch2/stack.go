package main

import "fmt"

func main() {
	var str string
	fmt.Println("Input some string :")
	fmt.Scanln(&str)

	if IsPalindromic(str) == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	return
}

func IsPalindromic(str string) bool {
	top := 0
	len := len(str)
	mid := len / 2

	next := 0
	if len%2 == 0 {
		next = mid
	} else {
		next = mid + 1
	}

	var stack [100]string

	for i := 0; i < mid; i++ {
		top++
		stack[top] = string(str[i])
	}

	for i := next; i <= len-1; i++ {
		if string(str[i]) != stack[top] {
			break
		}
		top--
	}
	if top == 0 {
		return true
	}
	return false
}

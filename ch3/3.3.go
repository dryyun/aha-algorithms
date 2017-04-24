package main

import "fmt"

func main() {
	num := 18

	match := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6} // 0 - 9 每个数字对于的火柴数

	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			addM := 0
			add := i + j
			if add >= 10 {
				addM = match[add/10] + match[add%10]
			} else {
				addM = match[add]
			}
			addI := 0
			if i >= 10 {
				addI = match[i/10] + match[i%10]
			} else {
				addI = match[i]
			}

			addJ := 0
			if j >= 10 {
				addJ = match[j/10] + match[j%10]
			} else {
				addJ = match[j]
			}
			if addI+addJ+addM == num-4 {
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
			}
		}
	}

	return

}

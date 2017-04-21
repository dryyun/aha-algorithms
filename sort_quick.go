// 快速排序
package main

import (
	"fmt"
	"strings"
)

func main() {
	score := [...]int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Println(score)
	result := QuickSort(score[:], 0, len(score)-1, "asc")
	fmt.Println(result)
}

func QuickSort(score []int, left int, right int, order string) []int {
	if left < right {
		i := left
		j := right
		for i != j {
			//从小到大排序
			if strings.ToUpper(order) == "ASC" {
				//从右往左找
				for score[j] >= score[left] && i < j {
					j--
				}
				//从左往右找
				for score[i] <= score[left] && i < j {
					i++
				}
			} else { //从大到小排序
				//从右往左找
				for score[j] <= score[left] && i < j {
					j--
				}
				//从左往右找
				for score[i] >= score[left] && i < j {
					i++
				}
			}

			if i != j {
				score[i], score[j] = score[j], score[i]
			}
		}
		score[left], score[i] = score[i], score[left]
		score = QuickSort(score, left, i-1, order)
		score = QuickSort(score, i+1, right, order)
	}
	return score
}

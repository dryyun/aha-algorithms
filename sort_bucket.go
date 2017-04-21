// 桶排序
package main

import (
	"fmt"
	"strings"
)

func main() {

	bucketSize := 11 // 桶的大小

	score := [...]uint{5, 3, 5, 2, 10, 8}

	result := BucketSort(bucketSize, score[:], "asc")

	fmt.Println(result)

	return
}

func BucketSort(num int, score []uint, order string) []uint {
	var result []uint
	var bucket []int = make([]int, num)
	for i := 0; i < len(score); i++ {
		bucket[score[i]]++
	}

	if strings.ToUpper(order) == "ASC" {
		for k, v := range bucket {
			if v > 0 {
				for i := 0; i < v; i++ {
					result = append(result, uint(k))
				}
			}
		}
	} else {
		for i := num - 1; i >= 0; i-- {
			if bucket[i] > 0 {
				for j := 0; j < bucket[i]; j++ {
					result = append(result, uint(i))
				}
			}
		}
	}

	return result
}

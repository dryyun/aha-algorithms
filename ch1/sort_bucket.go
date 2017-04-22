// 桶排序
package main

import (
	"aha-algorithms"
	"fmt"
	"strings"
)

func main() {

	bucketSize := 21 // 桶的大小

	arr := aha_algorithms.RandArray(bucketSize - 1)

	fmt.Println("Initial array is:", arr)

	ascResult := BucketSort(bucketSize, arr, "asc")

	fmt.Println("ASC result :", ascResult)

	descResult := BucketSort(bucketSize, arr, "desc")

	fmt.Println("DESC result :", descResult)

	return
}

func BucketSort(num int, arr []int, order string) []int {
	var result []int
	var bucket []int = make([]int, num)
	for i := 0; i < len(arr); i++ {
		bucket[arr[i]]++
	}

	if strings.ToUpper(order) == "ASC" {
		for k, v := range bucket {
			if v > 0 {
				for i := 0; i < v; i++ {
					result = append(result, k)
				}
			}
		}
	} else {
		for i := num - 1; i >= 0; i-- {
			if bucket[i] > 0 {
				for j := 0; j < bucket[i]; j++ {
					result = append(result, i)
				}
			}
		}
	}

	return result
}

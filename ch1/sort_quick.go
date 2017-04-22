// 快速排序
package main

import (
	"aha-algorithms"
	"fmt"
	"strings"
)

func main() {

	arr := aha_algorithms.RandArray(20)

	fmt.Println("Initial array is:", arr)

	ascResult := QuickSort(arr, "asc")

	fmt.Println("ASC result :", ascResult)

	descResult := QuickSort(arr, "desc")

	fmt.Println("DESC result :", descResult)

	return
}

func QuickSort(arr []int, order string) []int {
	if len(arr) <= 1 {
		return arr
	}
	return quickSort(arr, 0, len(arr)-1, order)
}

func quickSort(arr []int, left int, right int, order string) []int {
	if left >= right {
		return arr
	}
	if 1 == right-left {
		if strings.ToUpper(order) == "ASC" && arr[right] < arr[left] {
			arr[right], arr[left] = arr[left], arr[right]
		} else if arr[right] > arr[left] {
			arr[right], arr[left] = arr[left], arr[right]
		}
		return arr
	}
	i := left
	j := right
	for i != j {
		//从小到大排序
		if strings.ToUpper(order) == "ASC" {
			//从右往左找
			for arr[j] >= arr[left] && i < j {
				j--
			}
			//从左往右找
			for arr[i] <= arr[left] && i < j {
				i++
			}
		} else { //从大到小排序
			//从右往左找
			for arr[j] <= arr[left] && i < j {
				j--
			}
			//从左往右找
			for arr[i] >= arr[left] && i < j {
				i++
			}
		}

		if i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[left], arr[i] = arr[i], arr[left]
	arr = quickSort(arr, left, i-1, order)
	arr = quickSort(arr, i+1, right, order)

	return arr
}

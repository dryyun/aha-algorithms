// 冒泡排序
package main

import (
	"aha-algorithms"
	"fmt"
	"strings"
)

type NameAndScore struct {
	Name  string
	Score int
}

func main() {

	arr := aha_algorithms.RandArray(20)

	fmt.Println("Initial array is:", arr)

	ascResult := BubbleSort(arr, "asc")

	fmt.Println("ASC result :", ascResult)

	descResult := BubbleSort(arr, "desc")

	fmt.Println("DESC result :", descResult)

	s := [...]NameAndScore{
		{"ttt", 5},
		{"qqq", 3},
		{"vvv", 10},
		{"xxx", 2},
		{"aaa", 8},
	}

	r := BubbleSortOnStruct(s[:], "asc")
	fmt.Println("ASC On Struct", r)

	return
}

func BubbleSort(arr []int, order string) []int {
	//从小到大，把小的往前排
	if strings.ToUpper(order) == "ASC" {
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] > arr[j] {
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
		}
	} else {
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] < arr[j] {
					arr[j], arr[i] = arr[i], arr[j]
				}
			}
		}
	}
	return arr
}

func BubbleSort2(arr []int, order string) []int {
	//从小到大，把大的往后排
	if strings.ToUpper(order) == "ASC" {
		for i := 0; i < len(arr); i++ {
			for j := 1; j < len(arr)-i; j++ {
				if arr[j-1] > arr[j] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
	} else {
		for i := 0; i < len(arr); i++ {
			for j := 1; j < len(arr)-i; j++ {
				if arr[j-1] < arr[j] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
	}
	return arr
}

func BubbleSortOnStruct(s []NameAndScore, order string) []NameAndScore {
	score := s
	//从小到大，把大的往后排
	if strings.ToUpper(order) == "ASC" {
		for i := 0; i < len(score); i++ {
			for j := 1; j < len(score)-i; j++ {
				if score[j-1].Score > score[j].Score {
					score[j], score[j-1] = score[j-1], score[j]
				}
			}
		}
	} else {
		for i := 0; i < len(score); i++ {
			for j := 1; j < len(score)-i; j++ {
				if score[j-1].Score < score[j].Score {
					score[j], score[j-1] = score[j-1], score[j]
				}
			}
		}
	}
	return score
}

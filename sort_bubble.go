// 冒泡排序
package main

import (
	"fmt"
	"strings"
)

type NameAndScore struct {
	Name  string
	Score int
}

func main() {
	score := [...]int{5}

	result := BubbleSort2(score[:], "asc")

	fmt.Println(result)

	s := [...]NameAndScore{
		{"ttt", 5},
		{"qqq", 3},
		{"vvv", 10},
		{"xxx", 2},
		{"aaa", 8},
	}

	r := BubbleSortOnStruct(s[:], "desc")
	fmt.Println(r)
}

func BubbleSort(score []int, order string) []int {
	//从小到大，把小的往前排
	if strings.ToUpper(order) == "ASC" {
		for i := 0; i < len(score); i++ {
			for j := i + 1; j < len(score); j++ {
				if score[i] > score[j] {
					score[j], score[i] = score[i], score[j]
				}
			}
		}
	} else {
		for i := 0; i < len(score); i++ {
			for j := i + 1; j < len(score); j++ {
				if score[i] < score[j] {
					score[j], score[i] = score[i], score[j]
				}
			}
		}
	}
	return score
}

func BubbleSort2(score []int, order string) []int {
	//从小到大，把大的往后排
	if strings.ToUpper(order) == "ASC" {
		for i := 0; i < len(score); i++ {
			for j := 1; j < len(score)-i; j++ {
				if score[j-1] > score[j] {
					score[j], score[j-1] = score[j-1], score[j]
				}
			}
		}
	} else {
		for i := 0; i < len(score); i++ {
			for j := 1; j < len(score)-i; j++ {
				if score[j-1] < score[j] {
					score[j], score[j-1] = score[j-1], score[j]
				}
			}
		}
	}
	return score
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

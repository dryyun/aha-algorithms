package aha_algorithms

import (
	"math/rand"
	"time"
)

func RandArray(n int) []int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n)
	for i := 0; i <= n-1; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

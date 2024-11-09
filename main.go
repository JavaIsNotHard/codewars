package main

import (
	"fmt"
	"math/rand"
)

func partition(arr []int, p int, r int) int {
	x := arr[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}

func randomizedPartition(arr []int, p int, r int) int {
	i := rand.Intn(r-p+1) + p
	arr[r], arr[i] = arr[i], arr[r]
	return partition(arr, p, r)
}

func randomizedSelect(arr []int, p int, r int, i int) int {
	if p == r {
		return arr[p]
	}

	q := randomizedPartition(arr, p, r)
	k := q - p + 1

	if i == k {
		return arr[q]
	} else if i < k {
		return randomizedSelect(arr, p, q-1, i)
	} else {
		return randomizedSelect(arr, q+1, r, i-k)
	}
}

func main() {
	var arr []int = []int{1, 2, 3, 4, 5}
	fmt.Println(randomizedSelect(arr, 0, len(arr)-1, 5))
}

package main

import "fmt"

func multi3or5(number int) int {
	var result []int
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = append(result, i)
		}
	}

	fmt.Println(result)

	var sum = 0
	for _, value := range result {
		sum += value
	}

	return sum
}

func main() {
	fmt.Println(multi3or5(10))
}

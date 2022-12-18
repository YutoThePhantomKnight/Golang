package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 34, 5, 6, 7, 94, 252, 4656, 1}
	result := filter(nums, func(num int) bool {
		if num <= 10 {
			return true
		}
		return false
	})

	fmt.Println(result)
}

func filter(nums []int, callback func(int) bool) []int {
	result := []int{}

	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

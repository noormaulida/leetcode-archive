package main

import "fmt"

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		index := 0
		min := nums[index]
		for i, num := range nums {
			if num < min {
				min = num
				index = i
			}
		}
		nums[index] = min * multiplier
	}

	return nums
}

func main() {
	fmt.Println("CASE 1")
	nums := []int{3, 12, 30, 17, 21}
	k := 5
	multiplier := 2

	result := getFinalState(nums, k, multiplier)
	fmt.Println(result)

	fmt.Println("CASE 2")
	nums = []int{1, 2}
	k = 3
	multiplier = 4

	result = getFinalState(nums, k, multiplier)
	fmt.Println(result)

}

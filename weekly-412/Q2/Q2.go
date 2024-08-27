package main

import (
	"fmt"
	"sort"
	"strconv"
)

func countPairs(nums []int) (ans int) {
	sort.Ints(nums)
	cnt := make(map[int]int)

	for _, x := range nums {
		vis := make(map[int]struct{})
		vis[x] = struct{}{}
		s := []rune(strconv.Itoa(x))

		for j := 0; j < len(s); j++ {
			for i := 0; i < j; i++ {
				s[i], s[j] = s[j], s[i]
				y, _ := strconv.Atoi(string(s))
				vis[y] = struct{}{}
				s[i], s[j] = s[j], s[i]
			}
		}

		for y := range vis {
			ans += cnt[y]
		}
		cnt[x]++
	}

	return
}

func main() {
	fmt.Println("CASE 1")
	numbers := []int{3, 12, 30, 17, 21}
	result := countPairs(numbers)
	fmt.Println(result)

	fmt.Println("CASE 2")
	numbers = []int{1, 1, 1, 1, 1}
	result = countPairs(numbers)
	fmt.Println(result)

	fmt.Println("CASE 3")
	numbers = []int{123, 231}
	result = countPairs(numbers)
	fmt.Println(result)

}

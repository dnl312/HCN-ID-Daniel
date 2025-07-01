package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func main() {
	arr := []int{1, 4, 3, 2}
	result := minPairSum(arr)
	fmt.Println("Jumlah pasangan terkecil maksimal:", result)
}

package main

import "fmt"

func main() {
	nums := []int{2, -1, -1, 1, 5}
	k := 1
	println(subarraySum(nums, k))
}

func subarraySum(nums []int, k int) int {
	count := 0
	hash := map[int]int{0: 1}
	preSum := 0

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		fmt.Printf("hash %v\n", hash)
		fmt.Printf("preSum: %v\n", preSum)
		if hash[preSum-k] > 0 {
			fmt.Printf("preSum: %v, preSum-k: %v, hash[preSum-k]: %v\n", preSum, preSum-k, hash[preSum-k])
			count += hash[preSum-k]
		}
		hash[preSum]++
	}

	fmt.Printf("%v", hash)

	return count
}

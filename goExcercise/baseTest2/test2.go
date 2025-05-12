package main

import (
	"fmt"
)

func main() {
	nums := []int{101, 952, 32, 4, 85, 156, 267, 648, 929, 110}
	n := len(nums)
	if n == 0 {
		fmt.Println(0)
	}
	if n == 1 {
		fmt.Println(nums[0])
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {

		dp[i] = max(dp[i-1], dp[i-2]+nums[i])

	}

	fmt.Println(dp)

}

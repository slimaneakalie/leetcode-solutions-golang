package dynamic_programming

import "math"

// Problem: https://leetcode.com/problems/coin-change/
// Solution's complexity: O(a * c) time, O(a) space where a is the amount and c is the number of coins

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for a := 1; a <= amount; a++ {
		dp[a] = math.MaxInt32
		for _, coin := range coins {
			remaining := a - coin
			if remaining >= 0 {
				dp[a] = min(dp[a], dp[remaining]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

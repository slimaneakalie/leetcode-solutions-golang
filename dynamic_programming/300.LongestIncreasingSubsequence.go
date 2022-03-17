package dynamic_programming

import "github.com/slimaneakalie/leetcode-solutions-golang/utils"

// Problem: https://leetcode.com/problems/longest-increasing-subsequence
// Solution's complexity: O(n^2) time, O(n) space where n is the size of the input

func lengthOfLIS(nums []int) int {
	size := len(nums)
	dp, maxLength := make([]int, size), 1
	dp[size-1] = 1

	for i := size - 2; i >= 0; i-- {
		maxSoFar := 0

		for j := i + 1; j < size; j++ {
			if nums[j] > nums[i] {
				maxSoFar = utils.MaxInts(maxSoFar, dp[j])
			}
		}

		dp[i] = maxSoFar + 1
		maxLength = max(maxLength, dp[i])
	}

	return maxLength
}

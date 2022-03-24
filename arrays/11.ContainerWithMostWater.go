package arrays

// Problem: https://leetcode.com/problems/container-with-most-water/
// Solution's complexity: O(n) time, O(1) space where n is the size of the input

func maxArea(height []int) int {
	left, right, maxAmount := 0, len(height)-1, 0

	for left < right {
		gap := right - left

		var minAmount int
		if height[left] < height[right] {
			minAmount, left = height[left], left+1
		} else {
			minAmount, right = height[right], right-1
		}

		maxAmount = max(gap*minAmount, maxAmount)
	}

	return maxAmount
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

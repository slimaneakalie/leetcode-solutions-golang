package arrays

import "sort"

// Problem: https://leetcode.com/problems/3sum/
// Solution's complexity: O(n²) time, O(n) space where n is the size of the input

func threeSum(nums []int) [][]int {
	size, triplets := len(nums), [][]int{}
	if size < 3 {
		return triplets
	}

	sortedNums := make([]int, size)
	copy(sortedNums, nums)

	sort.Ints(sortedNums)

	for i := 0; i < size-2; i++ {
		if i > 0 && sortedNums[i] == sortedNums[i-1] {
			continue
		}

		left, right, target := i+1, size-1, -sortedNums[i]
		for left < right {
			sum := sortedNums[left] + sortedNums[right]
			if sum == target {
				triplets = append(triplets, []int{sortedNums[i], sortedNums[left], sortedNums[right]})
				for left < right && sortedNums[left] == sortedNums[left+1] {
					left++
				}

				for right > left && sortedNums[right] == sortedNums[right-1] {
					right--
				}

				left, right = left+1, right-1
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return triplets
}

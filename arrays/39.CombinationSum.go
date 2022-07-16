package arrays

// Problem: https://leetcode.com/problems/combination-sum/
// Solution's complexity: O(2^T) time, O(2^T) space where T is the target

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	dfs(candidates, len(candidates), 0, 0, target, &[]int{}, &result)

	return result
}

func dfs(candidates []int, nbCandidates, currIdx, currTotal, target int, currComb *[]int, result *[][]int) {
	if currIdx >= nbCandidates || currTotal > target {
		return
	} else if currTotal == target {
		*result = append(*result, cloneIntSlice(*currComb))
		return
	}

	*currComb = append(*currComb, candidates[currIdx])
	dfs(candidates, nbCandidates, currIdx, currTotal+candidates[currIdx], target, currComb, result)

	*currComb = (*currComb)[:len(*currComb)-1]
	dfs(candidates, nbCandidates, currIdx+1, currTotal, target, currComb, result)
}

func cloneIntSlice(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	return c
}

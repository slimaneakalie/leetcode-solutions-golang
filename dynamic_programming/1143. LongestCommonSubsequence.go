package dynamic_programming

// Problem: https://leetcode.com/problems/longest-common-subsequence/
// Solution's complexity: O(l * s) time, O(s) space where l is the length of the longest string and s is the length of the shortest

func longestCommonSubsequence(text1 string, text2 string) int {
	shortStr, longStr := text1, text2
	shortStrSize, longStrSize := len(text1), len(text2)

	if shortStrSize > longStrSize {
		shortStr, longStr = text2, text1
		shortStrSize, longStrSize = longStrSize, shortStrSize
	}

	dpNxt := make([]int, shortStrSize+1)

	for i := longStrSize - 1; i >= 0; i-- {
		dpCurr := make([]int, shortStrSize+1)
		for j := shortStrSize - 1; j >= 0; j-- {
			if shortStr[j] == longStr[i] {
				dpCurr[j] = dpNxt[j+1] + 1
			} else {
				dpCurr[j] = max(dpCurr[j+1], dpNxt[j])
			}
		}

		dpNxt = dpCurr
	}

	return dpNxt[0]
}

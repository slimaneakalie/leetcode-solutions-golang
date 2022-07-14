package dynamic_programming

// Problem: https://leetcode.com/problems/word-break
// Solution's complexity: O(D*W) time, O(S) space where D is the size of the dictionary, W is the size of the longest word in the dictionary, and S is the size of the string s

func wordBreak(s string, wordDict []string) bool {
	sSize := len(s)
	dp := make([]bool, sSize+1)
	dp[sSize] = true
	for i := sSize - 1; i >= 0; i-- {
		for _, w := range wordDict {
			wSize := len(w)
			dp[i] = sSize-i >= wSize && s[i:i+wSize] == w && dp[i+wSize]
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}

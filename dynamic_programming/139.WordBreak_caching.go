package dynamic_programming

// Problem: https://leetcode.com/problems/word-break
// Solution's complexity: O(D*W) time, O(D*W + S) space where D is the size of the dictionary, W is the size of the longest word in the dictionary, and S is the size of the string s

func wordBreak(s string, wordDict []string) bool {
    wordMap := toWordMap(wordDict)
    memo := map[int]bool{}
    
    return wordBreakRecu(s, wordMap, memo, 0)
}

func toWordMap(wordDict []string) map[byte][]string{
    wordMap := map[byte][]string{}
    for _, w := range wordDict {
        arr, exists := wordMap[w[0]]
        if !exists {
            arr = []string{w}
        } else {
            arr = append(arr, w)
        }
        
        wordMap[w[0]] = arr
    }
    
    return wordMap
}

func wordBreakRecu(s string, wordMap map[byte][]string, memo map[int]bool, i int) bool{
    if i >= len(s){
        return true
    } else if val, exists := memo[i]; exists{
        return val
    }
    
    arr, exists := wordMap[s[i]]
    if !exists {
        memo[i] = false
        return false
    }
    
    for _, w := range arr {
        subLimit := i + len(w)
        if subLimit <= len(s) && s[i:subLimit] == w && wordBreakRecu(s, wordMap, memo, subLimit){
            memo[i] = true
            return true
        }
    }
    
    memo[i] = false
    return false
}
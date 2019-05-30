package main

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for i := range wordDict {
		dict[wordDict[i]] = true
	}
	dp := make(map[int]bool)
	wb(s, 0, &dict, &dp)
	return dp[0]
}

func wb(s string, idx int, dict *map[string]bool, dp *map[int]bool) bool {
	if idx == len(s) {
		return true
	}
	res, found := (*dp)[idx]
	if found {
		return res
	}
	for i := idx; i < len(s); i++ {
		partial := s[idx : i+1]
		_, found := (*dict)[partial]
		if found && wb(s, i+1, dict, dp) {
			(*dp)[idx] = true
			return true
		}
	}
	(*dp)[idx] = false
	return false
}

// func main() {
// 	s := "hello"
// 	fmt.Println(s[1:3])
// }

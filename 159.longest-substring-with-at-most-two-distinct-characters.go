package main

import (
	"math"
)

/*
 * @lc app=leetcode id=159 lang=golang
 *
 * [159] Longest Substring with At Most Two Distinct Characters
 */
func lengthOfLongestSubstringTwoDistinct(s string) int {
	i, j := 0, 0
	count := make(map[string]int)
	maxLen := math.MinInt32
	for j < len(s) {
		ch := string(s[j])
		if len(count) == 2 && count[ch] == 0 {
			for i <= j && len(count) == 2 {
				if remove := string(s[i]); count[remove] == 1 {
					delete(count, remove)
				} else {
					count[remove]--
				}
				i++
			}
		}
		count[ch]++
		if curLen := j - i; len(count) <= 2 && curLen > maxLen {
			maxLen = curLen
		}
		j++
	}
	if maxLen == math.MinInt32 {
		return 0
	}
	return maxLen + 1
}

/* func main() {
	fmt.Println(lengthOfLongestSubstringTwoDistinct("a"))       // 1
	fmt.Println(lengthOfLongestSubstringTwoDistinct(""))        // 0
	fmt.Println(lengthOfLongestSubstringTwoDistinct("eceba"))   // 3
	fmt.Println(lengthOfLongestSubstringTwoDistinct("ccaabbb")) //5
} */

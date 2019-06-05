package main

import (
	"math"
)

/*
 * @lc app=leetcode id=340 lang=golang
 *
 * [340] Longest Substring with At Most K Distinct Characters
 */
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	i, j := 0, 0
	count := make(map[string]int)
	maxLen := math.MinInt32
	for j < len(s) {
		ch := string(s[j])
		if len(count) == k && count[ch] == 0 {
			for i <= j && len(count) == k {
				if remove := string(s[i]); count[remove] == 1 {
					delete(count, remove)
				} else {
					count[remove]--
				}
				i++
			}
		}
		count[ch]++
		if curLen := j - i; len(count) <= k && curLen > maxLen {
			maxLen = curLen
		}
		j++
	}
	if maxLen == math.MinInt32 {
		return 0
	}
	return maxLen + 1
}

/*
func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))
	fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))
} */

package main

import (
	"math"
)

/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */
func minWindow(s string, t string) string {
	ref := make(map[string]int)
	for _, each := range t {
		letter := string(each)
		ref[letter] = ref[letter] + 1
	}
	count := make(map[string]int)
	i, j, matched := 0, 0, 0
	minLen := math.MaxInt32
	begin, end := -1, -1
	for j < len(s) {
		letter := string(s[j])
		if ref[letter] == 0 {
			j++
			continue
		}
		count[letter] = count[letter] + 1
		if count[letter] == ref[letter] {
			matched++
		}
		for i <= j && matched == len(ref) {
			if j-i < minLen {
				minLen = j - i
				begin = i
				end = j
			}
			remove := string(s[i])
			if ref[remove] == 0 {
				i++
				continue
			}
			if count[remove] == ref[remove] {
				matched--
			}
			count[remove] = count[remove] - 1
			i++
		}
		j++
	}
	if begin == -1 {
		return ""
	}
	return s[begin : end+1]
}

/* func main() {
	fmt.Println(minWindow("A", "A"))
	fmt.Println(minWindow("AA", "AA"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
} */

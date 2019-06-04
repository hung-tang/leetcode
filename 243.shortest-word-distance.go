package main

import (
	"math"
	"strings"
)

/*
 * @lc app=leetcode id=243 lang=golang
 *
 * [243] Shortest Word Distance
 */
func shortestDistance(words []string, word1 string, word2 string) int {
	idx1 := -1
	idx2 := -1
	res := math.MaxInt32
	for i, word := range words {
		if strings.Compare(word, word1) == 0 {
			idx1 = i
			if idx2 != -1 {
				res = min2(res, int(math.Abs(float64(idx1-idx2))))
			}
		}
		if strings.Compare(word, word2) == 0 {
			idx2 = i
			if idx1 != -1 {
				res = min2(res, int(math.Abs(float64(idx1-idx2))))
			}
		}
	}
	return res
}

func min2(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

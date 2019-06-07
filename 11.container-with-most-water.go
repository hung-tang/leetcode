package main

import (
	"math"
)

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
func maxArea(height []int) int {
	if height == nil || len(height) <= 1 {
		return 0
	}
	i, j := 0, len(height)-1
	max := math.MinInt32
	for i < j {
		len := j - i
		max = max3(max, min3(height[i], height[j])*len)
		if height[i] <= height[j] {
			i++
			continue
		}
		j--
	}
	if max == math.MinInt32 {
		return 0
	}
	return max
}

func min3(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func max3(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

/* func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h)) // 49
} */

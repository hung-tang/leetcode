package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
func threeSum(nums []int) [][]int {
	res := make([]Helper, 0, 5)
	sort.Ints(nums)
	set := make(map[Helper]bool)
	for i := 0; i < len(nums)-2; i++ {
		// a + b + c = 0 => -c = a + b
		j := i + 1
		k := len(nums) - 1
		c := -nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			if sum == c {
				h := Helper{nums[j], nums[k], -c}
				_, found := set[h]
				if !found {
					res = append(res, Helper{nums[j], nums[k], -c})
				}
				set[h] = true
				j++
				k--
			} else {
				if sum > c {
					k--
				} else {
					j++
				}
			}
		}
	}

	x := make([][]int, 0, 5)
	for _, h := range res {
		r := make([]int, 0, 3)
		r = append(r, h.a)
		r = append(r, h.b)
		r = append(r, h.c)
		x = append(x, r)
	}
	return x
}

// func main() {
// 	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
// }

type Helper struct {
	a int
	b int
	c int
}

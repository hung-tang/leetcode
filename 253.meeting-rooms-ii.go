package main

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=253 lang=golang
 *
 * [253] Meeting Rooms II
 */
func minMeetingRooms(intervals [][]int) int {
	f := make([][]int, 0)
	s := make([][]int, 0)
	for i := range intervals {
		f = append(f, intervals[i])
		s = append(s, intervals[i])
	}
	sort.Slice(f, func(i, j int) bool {
		return f[i][0] < f[j][0]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i][1] < s[j][1]
	})
	res := 0
	count := 0
	i := 0
	j := 0
	for i < len(f) && j < len(s) {
		for j < len(s) && s[j][1] <= f[i][0] {
			j++
			count--
		}
		count++
		res = int(math.Max(float64(res), float64(count)))
		i++
	}
	return res
}

/* func main() {
	x := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}
	fmt.Println(minMeetingRooms(x))

	x = [][]int{
		{7, 10},
		{2, 4},
	}
	fmt.Println(minMeetingRooms(x))
} */

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
	f, s := make([]int, len(intervals)), make([]int, len(intervals))
	for i := range intervals {
		f[i], s[i] = intervals[i][0], intervals[i][1]
	}
	sort.Ints(f)
	sort.Ints(s)
	res := 0
	count := 0
	i := 0
	j := 0
	for i < len(f) && j < len(s) {
		for j < len(s) && s[j] <= f[i] {
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

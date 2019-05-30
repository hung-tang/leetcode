package main

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	// empty case
	if intervals == nil || len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	if newInterval == nil {
		return res
	}
	// when newInterval is the first element
	if newInterval[1] < intervals[0][0] {
		res = append(res, newInterval)
		res = append(res, intervals...)
		return res
	}
	// when newInterval is the last element
	if newInterval[0] > intervals[len(intervals)-1][1] {
		res = append(res, intervals...)
		res = append(res, newInterval)
		return res
	}
	// find out the interval (overlappingIdx) where we start to insert newInterval
	overlappingIdx := 0
	for i := 0; i < len(intervals); i++ {
		// non-overlap
		if newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
		} else { // overlap
			overlappingIdx = i
			break
		}
	}

	// 2 cases to consider: (1) non-overlap, (2) overlap
	// Case 1 - non-overlap
	if newInterval[1] < intervals[overlappingIdx][0] {
		res = append(res, newInterval)
		for i := overlappingIdx; i < len(intervals); i++ {
			res = append(res, intervals[i])
		}
		return res
	}

	// Case 2 - overlap, use expansion strategy
	expand := newInterval
	for i := overlappingIdx; i < len(intervals); i++ {
		current := intervals[i]
		if current[0] > expand[1] { // expansion is done when there's no overlap
			res = append(res, expand)
			for j := i; j < len(intervals); j++ {
				res = append(res, intervals[j])
			}
			return res
		}
		// keep expanding if there's overlap
		if current[1] > expand[1] { // merge - new minimum
			expand[1] = current[1]
		}
		if current[0] < expand[0] { // merge - new maximum
			expand[0] = current[0]
		}
	}
	// edge case - add the expansion if this is the last item
	res = append(res, expand)
	return res
}

/* func main() {

	var x = [][]int{
		{3, 5},
		{12, 15},
	}
	fmt.Println(insert(x, []int{6, 6}))

	x = [][]int{
		{1, 5},
	}
	fmt.Println(insert(x, []int{2, 3}))

	x = [][]int{
		{1, 3},
		{6, 9},
	}
	fmt.Println(insert(x, []int{2, 5}))

	x = [][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}
	fmt.Println(insert(x, []int{4, 8}))

	x = [][]int{
		{4, 5},
		{6, 9},
	}
	fmt.Println(insert(x, []int{1, 3}))

	x = [][]int{
		{1, 3},
		{6, 9},
	}
	fmt.Println(insert(x, []int{10, 15}))
} */

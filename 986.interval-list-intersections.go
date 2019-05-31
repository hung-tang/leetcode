package main

/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */
func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	if A == nil || len(A) == 0 || B == nil || len(B) == 0 {
		return res
	}

	i, j := 0, 0
	var prev, cur []int
	for i < len(A) && j < len(B) {
		if prev == nil {
			if A[i][0] < B[j][0] || (A[i][0] == B[j][0] && A[i][1] < B[j][1]) {
				prev = A[i]
				i++
				continue
			}
			prev = B[j]
			j++
			continue
		}

		if A[i][0] < B[j][0] || (A[i][0] == B[j][0] && A[i][1] < B[j][1]) {
			cur = A[i]
			i++
		} else {
			cur = B[j]
			j++
		}
		overlapHelper(&cur, &prev, &res)
	}
	for i < len(A) {
		cur = A[i]
		i++
		overlapHelper(&cur, &prev, &res)
	}
	for j < len(B) {
		cur = B[j]
		j++
		overlapHelper(&cur, &prev, &res)
	}
	return res
}

func overlapHelper(cur *[]int, prev *[]int, res *[][]int) {
	// non-overlap
	if (*cur)[0] > (*prev)[1] {
		*prev = *cur
		return
	}

	// 3 cases for overlap: (1) prev is longer, (2) cur is longer, (3) prev and cur same length
	if (*cur)[1] == (*prev)[1] { // case 3
		*res = append(*res, []int{(*cur)[0], (*cur)[1]})
		*prev = *cur
	} else if (*cur)[1] > (*prev)[1] { // case 2
		*res = append(*res, []int{(*cur)[0], (*cur)[0] + (*prev)[1] - (*cur)[0]})
		*prev = *cur
	} else { // case 1
		*res = append(*res, []int{(*cur)[0], (*cur)[1]})
	}
}

/* func main() {

	a := [][]int{
		{4, 6},
		{7, 8},
		{10, 17},
	}
	b := [][]int{
		{5, 10},
	}
	fmt.Println(intervalIntersection(a, b))

	a = [][]int{
		{3, 5},
		{9, 20},
	}
	b = [][]int{
		{4, 5},
		{7, 10},
		{11, 12},
		{14, 15},
		{16, 29},
	}
	fmt.Println(intervalIntersection(a, b))

	a = [][]int{
		{1, 8},
		{16, 20},
	}
	b = [][]int{
		{2, 11},
		{12, 13},
	}
	fmt.Println(intervalIntersection(a, b))

	a = [][]int{
		{0, 2},
		{5, 10},
		{13, 23},
		{24, 25},
	}
	b = [][]int{
		{1, 5},
		{8, 12},
		{15, 24},
		{25, 26},
	}
	fmt.Println(intervalIntersection(a, b))
	a = [][]int{
		{8, 15},
	}
	b = [][]int{
		{2, 6},
		{8, 10},
		{12, 20},
	}
	fmt.Println(intervalIntersection(a, b))
} */

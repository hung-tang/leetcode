package main

/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
func largestRectangleArea(heights []int) int {
	s := make([][]int, 0)
	i := 0
	res := 0
	for i < len(heights) {
		cur := heights[i]
		res = max2(res, cur)
		if len(s) == 0 || cur >= s[len(s)-1][1] {
			s = append(s, []int{1, heights[i]})
			i++
			continue
		}
		count := 0
		for len(s) > 0 && s[len(s)-1][1] > cur {
			count += s[len(s)-1][0]
			res = max2(res, s[len(s)-1][1]*count)
			s = s[:len(s)-1]
		}
		s = append(s, []int{count + 1, heights[i]})
		i++
	}
	if len(s) == 0 {
		return res
	}
	count := 0
	for len(s) > 0 {
		count += s[len(s)-1][0]
		res = max2(res, s[len(s)-1][1]*count)
		s = s[:len(s)-1]
	}
	return res
}

func max2(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

/*
func main() {
	a := []int{3, 6, 5, 7, 4, 8, 1, 0}
	fmt.Println(largestRectangleArea(a))

	a = []int{9, 0}
	fmt.Println(largestRectangleArea(a))

	a = []int{0}
	fmt.Println(largestRectangleArea(a))

	a = []int{2, 1, 2}
	fmt.Println(largestRectangleArea(a))

	a = []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(a))

} */

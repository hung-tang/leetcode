package main

/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
func trap(height []int) int {
	res := 0
	if height == nil || len(height) == 0 {
		return res
	}
	s := make([]int, 0) //idx
	for i := 0; i < len(height); {
		h := height[i]
		if len(s) == 0 || h <= height[s[len(s)-1]] {
			s = append(s, i)
			i++
			continue
		}
		if len(s) > 0 {
			// pop
			top := height[s[len(s)-1]]
			s = s[:len(s)-1]
			for len(s) > 0 && height[s[len(s)-1]] == top {
				s = s[:len(s)-1]
			}
			if len(s) > 0 {
				peek := height[s[len(s)-1]]
				len := i - s[len(s)-1] - 1
				width := min(h, peek) - top
				res += len * width
			}
		}
	}
	return res
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

/* func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(a))

	a = []int{5, 4, 1, 2}
	fmt.Println(trap(a)) // ans 1

	a = []int{3, 2, 1, 1, 1, 0, 2, 3}
	fmt.Println(trap(a)) // ans 11

	a = []int{3, 3, 1, 0, 1, 2, 3}
	fmt.Println(trap(a)) // ans 8

	a = []int{3, 2, 1, 1, 4}
	fmt.Println(trap(a)) // ans 5

	a = []int{2, 1, 0, 0, 2, 3}
	fmt.Println(trap(a)) // ans 5

	a = []int{5, 2, 1, 2, 1, 5}
	fmt.Println(trap(a)) // ans 14
} */

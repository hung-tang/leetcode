package main

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	open := make([]int, 0)
	regular := make([]string, 0)
	max := math.MinInt32
	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if ch == "(" {
			regular = append(regular, "(")
			open = append(open, len(regular)-1)
			continue
		}
		// ch == ')'
		if len(open) == 0 {
			regular = append(regular, ")")
			continue
		}
		idx := open[len(open)-1]
		open = open[:len(open)-1]

		if idx+1 >= len(regular) {
			regular = regular[:len(regular)-1] // remove "("
			regular = append(regular, "1")
			helper2(&regular, &max)
		} else {
			num, _ := strconv.Atoi(regular[len(regular)-1])
			regular = regular[:len(regular)-2]
			regular = append(regular, strconv.Itoa((num + 1)))
			helper2(&regular, &max)
		}
	}
	if max == math.MinInt32 {
		return 0
	}
	return max
}
func helper2(regular *[]string, max *int) {
	count := 0
	for len(*regular) > 0 && (*regular)[len(*regular)-1] != "(" && (*regular)[len(*regular)-1] != ")" {
		num, _ := strconv.Atoi((*regular)[len(*regular)-1])
		count += num
		*regular = (*regular)[:len(*regular)-1]
	}
	*max = max4(*max, count*2)
	*regular = append(*regular, strconv.Itoa(count))
}

func max4(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

/* func main() {
	fmt.Println(longestValidParentheses("()(())"))
	fmt.Println(longestValidParentheses("(()(((()"))
	fmt.Println(longestValidParentheses("()(()"))
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("))()())()))(()())()(("))
} */

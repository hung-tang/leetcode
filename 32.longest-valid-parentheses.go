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
	prevOpen := make([]int, 0)
	stack := make([]string, 0)
	max := math.MinInt32
	for i := 0; i < len(s); i++ {
		if ch := string(s[i]); ch == "(" {
			stack = append(stack, "(")
			prevOpen = append(prevOpen, len(stack)-1)
			continue
		}
		// ch == ')'
		if len(prevOpen) == 0 {
			stack = append(stack, ")")
			continue
		}

		idx := prevOpen[len(prevOpen)-1]
		prevOpen = prevOpen[:len(prevOpen)-1] // pop

		if idx+1 >= len(stack) { // case: "("
			stack = stack[:len(stack)-1] // remove the "("
			stack = append(stack, "1")
			addAndMergeNumbers(&stack, &max)
		} else { // case: "(5"
			num, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-2] // remove the "(5"
			stack = append(stack, strconv.Itoa((num + 1)))
			addAndMergeNumbers(&stack, &max)
		}
	}
	if max == math.MinInt32 {
		return 0
	}
	return max
}
func addAndMergeNumbers(stack *[]string, max *int) {
	count := 0
	for len(*stack) > 0 && (*stack)[len(*stack)-1] != "(" && (*stack)[len(*stack)-1] != ")" {
		num, _ := strconv.Atoi((*stack)[len(*stack)-1])
		count += num
		*stack = (*stack)[:len(*stack)-1]
	}
	*max = max4(*max, count*2)
	*stack = append(*stack, strconv.Itoa(count))
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

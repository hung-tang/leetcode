package main

import (
	"math"
)

/*
 * @lc app=leetcode id=687 lang=golang
 *
 * [687] Longest Univalue Path
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	lup(root, &res)
	return res
}

func lup(node *TreeNode, res *int) (int, int) {
	if node == nil {
		return -1, -1
	}
	lval, llen := lup(node.Left, res)
	rval, rlen := lup(node.Right, res)
	if llen == -1 && rlen == -1 { // both left and right are empty
		return node.Val, 0
	}
	if llen == -1 { // only left is empty
		if rval == node.Val {
			*res = int(math.Max(float64(*res), float64(rlen+1)))
			return node.Val, rlen + 1
		}
		return node.Val, 0
	}
	if rlen == -1 { //only right is empty
		if lval == node.Val {
			*res = int(math.Max(float64(*res), float64(llen+1)))
			return node.Val, llen + 1
		}
		return node.Val, 0
	}
	// both left and right is not empty
	// 4 cases: (1) left val matches node val only.
	//          (2) right val matches node val only
	//          (3) left and right both matches node val
	//          (4) left and right both don't match node val
	if lval != node.Val && rval != node.Val { // case 4
		return node.Val, 0
	}
	if lval == node.Val && rval == node.Val { // case 3
		*res = int(math.Max(float64(*res), float64(rlen+llen+2)))
		if llen == rlen || llen > rlen {
			return node.Val, llen + 1
		}
		return node.Val, rlen + 1
	}
	if rval == node.Val { // case 2
		*res = int(math.Max(float64(*res), float64(rlen+1)))
		return node.Val, rlen + 1
	}
	// case 1
	*res = int(math.Max(float64(*res), float64(llen+1)))
	return node.Val, llen + 1
}

/* func main() {
	n1 := TreeNode{Val: 5}
	n2 := TreeNode{Val: 4}
	n3 := TreeNode{Val: 5}
	n4 := TreeNode{Val: 1}
	n5 := TreeNode{Val: 1}
	n6 := TreeNode{Val: 5}
	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = &n5
	n3.Right = &n6

	n1 := TreeNode{Val: 1}
	n2 := TreeNode{Val: 1}
	n3 := TreeNode{Val: 1}
	n4 := TreeNode{Val: 1}
	n5 := TreeNode{Val: 1}
	n6 := TreeNode{Val: 1}
	n7 := TreeNode{Val: 1}
	n1.Right = &n2
	n2.Left = &n3
	n2.Right = &n4
	n3.Left = &n5
	n3.Right = &n6
	n4.Left = &n7
	fmt.Println(longestUnivaluePath(&n1))

}*/

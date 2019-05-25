package main

import "fmt"

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2 != nil && l2.Val == 0 && l2.Next == nil {
		return l1
	}
	n1 := l1
	n2 := l2
	dummy := ListNode{Val: 0, Next: nil}
	head := &dummy
	res := &dummy
	var carry int
	for n1 != nil && n2 != nil {
		v := n1.Val + n2.Val + carry
		carry = v / 10
		n := ListNode{Val: v % 10, Next: nil}
		head.Next = &n
		head = &n
		n1 = n1.Next
		n2 = n2.Next
	}
	for n1 != nil {
		v := n1.Val + carry
		carry = v / 10
		n := ListNode{Val: v % 10, Next: nil}
		head.Next = &n
		head = &n
		n1 = n1.Next
	}
	for n2 != nil {
		v := n2.Val + carry
		carry = v / 10
		n := ListNode{Val: v % 10, Next: nil}
		head.Next = &n
		head = &n
		n2 = n2.Next
	}
	if carry != 0 {
		n := ListNode{Val: carry, Next: nil}
		head.Next = &n
	}
	return res.Next
}

func reverse(l *ListNode) *ListNode {
	var n *ListNode
	var cur = l
	for cur != nil {
		nx := cur.Next
		cur.Next = n
		n = cur
		cur = nx
	}
	return n
}

func print(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

// func main() {
// 	al3 := ListNode{Val: 8, Next: nil}
// 	al2 := ListNode{Val: 9, Next: &al3}
// 	// al1 := ListNode{Val: 9, Next: &al2}

// 	bl3 := ListNode{Val: 1, Next: nil}
// 	// bl2 := ListNode{Val: 2, Next: &bl3}
// 	// bl1 := ListNode{Val: 1, Next: &bl2}

// 	print(addTwoNumbers(&al2, &bl3))
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

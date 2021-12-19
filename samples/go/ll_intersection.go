package main

import (
	. "./list"
)

func scanList(head *ListNode, cf func(*int) bool) *ListNode {
	for head != nil {
		if cf(&(head.Val)) {
			return head
		}
		head.Val = -head.Val
		head = head.Next
	}
	return nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	greater := func(a *int) bool { return *a < 0 }
	less := func(a *int) bool { return *a > 0 }
	scanList(headA, greater)
	res := scanList(headB, greater)
	scanList(headA, less)
	scanList(headB, less)
	return res
}

func main() {}

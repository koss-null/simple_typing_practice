package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	cur := head
	prev := &ListNode{Next: head}
	before_head := prev
	for cur != nil {
		cur_next := cur.Next
		if cur_next != nil {
			prev.Next = cur_next
			save_next := cur_next.Next
			cur_next.Next = cur
			cur.Next = save_next
			prev = cur
			cur = save_next
		} else {
			break
		}
	}
	return before_head.Next
}

func main() {}

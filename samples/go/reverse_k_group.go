package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	visited := make(map[int]interface{})
	for head != nil {
		if visited[head.Val] != nil {
			break
		}
		visited[head.Val] = struct{}{}
		if head.Next != nil {
		}
		head = head.Next
	}
}

func reverseTwo(head *ListNode, next *ListNode) (*ListNode, *ListNode) {
	next_next := next.Next
	next.Next = head
	return next, next_next
}

func isGreaterK(head *ListNode, k int) (*ListNode, bool) {
	var prev *ListNode
	for ; k > 0 && head != nil; k-- {
		prev = head
		head = head.Next
	}
	return prev, k == 0
}

func reverseK(head *ListNode, k int) (first *ListNode, last *ListNode, ok bool) {
	if last_el, ok := isGreaterK(head, k); !ok {
		return last_el, nil, false
	}

	prev, cur, next := &ListNode{Next: head}, head, head.Next
	for ; k > 1; k-- {
		var new_next *ListNode
		prev.Next, new_next = reverseTwo(cur, next)
		cur = next
		next = new_next
	}
	return prev.Next, next, true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 0 || k == 1 {
		return head
	}

	var next *ListNode
	prev := &ListNode{Next: head}
	mem_head := prev
	cur := head
	for cur != nil {
		var prev_next *ListNode
		var ok bool
		prev_next, next, ok = reverseK(cur, k)
		if !ok {
			break
		}
		prev.Next = prev_next
		cur.Next = next
		prev = cur
		cur = next
	}
	return mem_head.Next
}

func main() {
	head := ListNode{Val: 0, Next: nil}
	cur := &head
	for i := 1; i < 2; i++ {
		cur.Next = &ListNode{Val: i, Next: nil}
		cur = cur.Next
	}
	new_head := reverseKGroup(&head, 2)
}

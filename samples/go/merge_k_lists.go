package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Heap struct {
	data [][]int
}

func (h *Heap) BubbleUp(i int) {
	prev := (i - 1) / 2
	if h.data[prev][0] > h.data[i][0] {
		h.data[prev], h.data[i] = h.data[i], h.data[prev]
		if prev != 0 {
			h.BubbleUp(prev)
		}
	}
}

func (h *Heap) Insert(val, idx int) {
	h.data = append(h.data, []int{val, idx})
	h.BubbleUp(len(h.data) - 1)
}

func (h *Heap) BubbleDown(idx int) {
	lf, rg := idx*2+1, idx*2+2
	if lf < len(h.data) {
		if rg < len(h.data) {
			if h.data[rg][0] < h.data[lf][0] && h.data[idx][0] > h.data[rg][0] {
				h.data[rg], h.data[idx] = h.data[idx], h.data[rg]
				h.BubbleDown(rg)
			} else if h.data[idx][0] > h.data[lf][0] {
				h.data[lf], h.data[idx] = h.data[idx], h.data[lf]
				h.BubbleDown(lf)
			}
		}
		if h.data[lf][0] < h.data[idx][0] {
			h.data[lf], h.data[idx] = h.data[idx], h.data[lf]
			h.BubbleDown(lf)
		}
	}
}

func (h *Heap) Pop() (val, idx int) {
	if len(h.data) == 0 {
		return 0, -1
	}
	val, idx = h.data[0][0], h.data[0][1]
	h.data[0], h.data[len(h.data)-1] = h.data[len(h.data)-1], h.data[0]
	h.data = h.data[:len(h.data)-1]
	h.BubbleDown(0)
	return
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := Heap{make([][]int, 0, 127)}
	for i, v := range lists {
		if v != nil {
			h.Insert(lists[i].Val, i)
			lists[i] = lists[i].Next
		}
	}
	result_head := ListNode{}
	cur := &result_head
	var prev *ListNode
	for len(h.data) != 0 {
		val, idx := h.Pop()

		cur.Val = val
		prev = cur
		cur.Next = &ListNode{}
		cur = cur.Next

		if lists[idx] != nil {
			h.Insert(lists[idx].Val, idx)
			lists[idx] = lists[idx].Next
		}
	}

	if prev != nil {
		prev.Next = nil
	} else {
		return nil
	}

	return &result_head
}

func main() {
}

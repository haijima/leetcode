package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return _addTwoNumbers(l1, l2, 0)
}

func _addTwoNumbers(l1 *ListNode, l2 *ListNode, d int) *ListNode {
	if l1 == nil && l2 == nil {
		if d > 0 {
			return &ListNode{1, nil}
		} else {
			return nil
		}
	}
	if l1 == nil {
		l1 = &ListNode{0, nil}
	}
	if l2 == nil {
		l2 = &ListNode{0, nil}
	}

	r := &ListNode{}
	tmp := l1.Val + l2.Val + d
	if tmp >= 10 {
		r.Val = tmp - 10
		r.Next = _addTwoNumbers(l1.Next, l2.Next, 1)
	} else {
		r.Val = tmp
		r.Next = _addTwoNumbers(l1.Next, l2.Next, 0)
	}
	return r
}

package num83

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	var q *ListNode
	for current != nil && current.Next != nil {
		q = current.Next
		for q != nil && q.Val == current.Val {
			q = q.Next
		}
		current.Next = q
		current = current.Next
	}

	return head
}

func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head
	next := p.Next
	for p != nil && next != nil {
		if p.Val == next.Val {
			next = next.Next
		} else {
			p.Next = next
			p = p.Next
			next = p.Next
		}
	}
	p.Next = next

	return head
}

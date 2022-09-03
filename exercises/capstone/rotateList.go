func rotateRight(head *ListNode, k int) *ListNode {
	end, p := head, head
	length := 1

	if head == nil {
		return head
	}

	for end.Next != nil {
		end = end.Next
		length++
	}

	k = k % length
	toMoveP := length - k - 1

	for toMoveP > 0 {
		p = p.Next
		toMoveP--
	}

	end.Next = head
	head = p.Next
	p.Next = nil

	return head
}
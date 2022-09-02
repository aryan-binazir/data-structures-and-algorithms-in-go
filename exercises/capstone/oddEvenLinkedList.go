package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headEven, currentEven := head.Next, head.Next
	currentOdd := head

	for currentOdd.Next != nil {
		currentOdd.Next = currentEven.Next

		if currentEven.Next != nil { // cannot lose last node in even, need to set that Next property to beginning of even linked list
			currentOdd = currentEven.Next
		}

		currentEven.Next = currentOdd.Next
		currentEven = currentOdd.Next
	}
	currentOdd.Next = headEven

	return head
}

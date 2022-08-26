package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeElements(head *ListNode, val int) *ListNode {

	for head.Val == val {
		head = head.Next
	}

	current := head

	for current != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// start at head
// loop through array until current.next == nil
// if current.Next.Val == target
// current.next = current.next.next
// current = current.next
// return head

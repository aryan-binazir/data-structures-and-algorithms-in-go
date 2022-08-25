/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	current := head
	for current.Next != nil {
		if current.Next.Val == val {
			fmt.Println(current.Val, current.Next.Val)
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
package listnode

// 反转链表，用三个指针，prev，curr，next
func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	curr := node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

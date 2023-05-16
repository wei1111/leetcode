package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	list2 := sortList(slow.Next)
	slow.Next = nil
	list1 := sortList(head)
	return sort(list1, list2)
}

func sort(list1, list2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result
	for list1 != nil && list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				head.Next = list1
				list1 = list1.Next
			} else {
				head.Next = list2
				list2 = list2.Next
			}
		}
		head = head.Next
	}
	if list1 == nil {
		head.Next = list2
	}
	if list2 == nil {
		head.Next = list1
	}
	return result.Next
}

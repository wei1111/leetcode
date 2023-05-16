package leetcode

import (
	"leetcode/utils"
	"testing"
)

func Test_sortList(t *testing.T) {
	head := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	got := sortList(head)
	t.Logf("%v", utils.JsonStr(got))
}

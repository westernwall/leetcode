package AddTwoNumbers

import "testing"

func buildList(val ...int) *ListNode {
	l := new(ListNode)
	head := l
	for _, v := range val {
		l.Next = new(ListNode)
		l.Next.Val = v
		l = l.Next
	}

	return head.Next
}

func Test(t *testing.T) {
	l1 := buildList(2, 4, 3)
	l2 := buildList(5, 6, 6)
	addTwoNumbers(l1, l2)
}


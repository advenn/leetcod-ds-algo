package linked_list

func ReverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

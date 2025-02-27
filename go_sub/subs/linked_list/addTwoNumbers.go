package linked_list

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var out int
	var result = &ListNode{0, nil}
	var resultTail = result

	for l1 != nil || l2 != nil || carry > 0 {
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}

		carry, out = (val1+val2+carry)/10, (val1+val2+carry)%10
		resultTail.Next = &ListNode{out, nil}
		resultTail = resultTail.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return result.Next

}

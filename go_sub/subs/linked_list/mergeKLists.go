package linked_list

import (
	"math"
)

func Merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

func MergeKListsOld(lists []*ListNode) *ListNode {
	head := &ListNode{}
	ans := head
	pointers := make([]*ListNode, len(lists))
	copy(pointers, lists)

	for {
		nils := 0
		var smallest = math.MaxInt64
		var chosenIndex int = -1

		for index, list := range pointers {
			if list != nil {
				if list.Val < smallest {
					smallest = list.Val
					chosenIndex = index
				}
			} else {
				nils++
			}
		}
		if nils == len(pointers) {
			break
		}
		ans.Next = &ListNode{smallest, nil}
		ans = ans.Next
		pointers[chosenIndex] = pointers[chosenIndex].Next
	}
	return head.Next
}

func MergeKListsOldOld(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pointers := make([]*ListNode, len(lists))
	copy(pointers, lists)
	for len(pointers) > 1 {
		var merged = Merge2Lists(pointers[0], pointers[1])
		pointers[1] = merged
		pointers = pointers[1:]

	}
	return pointers[0]
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	pointers := make([]*ListNode, len(lists))
	copy(pointers, lists)

	for len(pointers) > 1 {
		var newPointers []*ListNode
		left := 0
		right := len(pointers) - 1

		for left <= right {
			var merged *ListNode
			if left != right {
				merged = Merge2Lists(pointers[left], pointers[right])
			} else {
				merged = pointers[left]
			}
			newPointers = append(newPointers, merged)
			left++
			right--
		}

		pointers = newPointers

	}
	return pointers[0]

}

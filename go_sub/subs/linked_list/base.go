package linked_list

import "fmt"

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func Arr2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head = &ListNode{arr[0], nil}
	cur := head
	for _, elem := range arr[1:] {
		cur.Next = &ListNode{elem, nil}
		cur = cur.Next
	}
	return head
}

func PrintList(prefix string, list *ListNode) {
	fmt.Printf("%s: ", prefix)
	for list != nil {
		fmt.Printf("%d -> ", list.Val)
		list = list.Next
	}
	fmt.Println("nil")
}

func PrintWholeList(lists []*ListNode) {
	for _, list := range lists {
		PrintList("whole", list)
	}
	fmt.Println()
}

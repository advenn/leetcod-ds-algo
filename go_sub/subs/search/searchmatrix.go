package search

import "fmt"

func SearchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1

	for left <= right {
		fmt.Printf("left: %v, right: %v\n", left, right)
		mid := (left + right) / 2
		arr := matrix[mid]
		if arr[0] <= target && target <= arr[len(arr)-1] {
			index := BinarySearch(arr, target)
			return index != -1
		} else if arr[0] < target && arr[len(arr)-1] < target {
			left = mid + 1
		} else if arr[0] > target && arr[len(arr)-1] > target {
			right = mid - 1
		}
	}

	return false

}

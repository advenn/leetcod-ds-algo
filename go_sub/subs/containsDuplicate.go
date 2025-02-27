package subs

import "fmt"

func containsDuplicate(nums []int) bool {
	var counter = make(map[int]int)
	for _, num := range nums {
		counter[num]++
		if counter[num] > 1 {
			return true
		}
	}
	return false

}

func main() {
	nums := []int{1, 2, 3, 1, 1}
	fmt.Println(
		containsDuplicate(nums),
	)

}

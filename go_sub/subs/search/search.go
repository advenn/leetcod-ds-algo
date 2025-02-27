package search

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] == 0 {

			}
		}
	}
	//max := slices.Max(nums)
	//first := nums[0]
	//fixedArr := append(nums[max])
}

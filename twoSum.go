package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	lenght := len(nums)
	for i := 0; i < lenght; i++ {
		if nums[i]+nums[i+1] == target {
			return []int{nums[i], nums[i+1]}
		}
	}
	return nil
}

package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k > 0 {
		first := len(nums) - k
		reverse(0, first-1, nums)
		reverse(first, len(nums)-1, nums)
		reverse(0, len(nums)-1, nums)
	}

}

func reverse(i, j int, nums []int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

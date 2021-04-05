
func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func rob(nums []int) int {

	inc := nums[0]
	exc := 0

	for i := 1; i < len(nums); i++ {
		tmp := inc
		inc = nums[i] + exc
		exc = max(tmp, exc)
	}
	return max(inc, exc)

}

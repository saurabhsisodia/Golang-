
//if there is any power of 2 which is greater than left and less than right,
// then result will always be zero (0)
// otherwise, right-left is not too large, so calculate Bitwise AND, using loop
func rangeBitwiseAnd(left int, right int) int {
	if left == 0 || right == 0 {
		return 0
	}
	if left == right {
		return left
	}
	x := math.Floor(math.Log2(float64(right)))
	if int(math.Pow(2, x)) > left {
		return 0
	}
	ans := left & (left + 1)
	for i := left + 2; i < right+1; i++ {
		ans = ans & i
	}
	return ans
}

// Another Method,
func rangeBitwiseAnd(left int, right int) int {

	var count int
	for left != right {
		if left == 0 || right == 0 {
			break
		}
		left >>= 1
		right >>= 1
		count += 1
	}
	if left != right {
		return 0
	}
	return left << count
}

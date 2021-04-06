package main

// Using sort comparator
import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {

	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) > strconv.Itoa(nums[j])+strconv.Itoa(nums[i])

	})
	ans := ""
	for _, v := range nums {
		ans += strconv.Itoa(v)
	}
	if ans[0] == '0' {
		return "0"
	}
	return ans

}

// Using Pure merge sort
/*
import (
	"strconv"
)

func largestNumber(nums []int) string {

	var slice []string
	for _, v := range nums {
		slice = append(slice, strconv.Itoa(v))
	}
	ans := mergeSort(slice)
	if ans[0][0] == '0' {
		return "0"
	}
	res := ""
	for _, v := range ans {
		res += v
	}
	return res

}

func merge(l []string, r []string) []string {
	var i, j int
	var res []string
	for i < len(l) && j < len(r) {
		if l[i]+r[j] > r[j]+l[i] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}
	for i < len(l) {
		res = append(res, l[i])
		i++
	}
	for j < len(r) {
		res = append(res, r[j])
		j++
	}
	return res
}

func mergeSort(arr []string) []string {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	l := mergeSort(arr[:mid])
	r := mergeSort(arr[mid:])
	return merge(l, r)
}

*/
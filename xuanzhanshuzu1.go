package main

import "fmt"

//旋转数组  解法1
//1.反转整个数组，再反转前k项，再反转后几项
func rev(nums []int, start, end int) []int {
	var temp int
	for start < end {
		temp = nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--

	}
	return nums
}
func rot(nums []int, k int) []int {
	k %= len(nums)
	rev(nums, 0, len(nums)-1)
	rev(nums, 0, k-1)
	rev(nums, k, len(nums)-1)
	return nums

}

func main() {
	n1 := []int{1, 2, 3, 4, 5}
	fmt.Println(rot(n1, 2))
}

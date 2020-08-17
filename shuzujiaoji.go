package main

import "fmt"

/*
给定两个数组，编写一个函数来计算它们的交集。
1.
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
2.
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
*/
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}
func main() {
	n1 := []int{9, 9, 9}
	fmt.Println(plusOne(n1))
}

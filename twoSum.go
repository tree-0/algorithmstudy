package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if val, exist := m[target-k]; exist {
			res = append(res, val)
			res = append(res, i)
		}
		m[k] = i
	}
	return res
}
func main() {
	n1 := []int{2, 7, 11, 15}
	fmt.Println(twoSum(n1,9))
}

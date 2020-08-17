package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/
func moveZeroes(nums []int) []int {
	if nums == nil {
		return nil
	}
	i, tmp := 0, 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != 0 {
			i++
		}
		if nums[j] == 0 {
			continue
		}
		tmp = nums[i]
		nums[i] = nums[j]
		nums[j] = tmp

	}
	return nums
}
func main() {
	n1 := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(n1))
}

package main

import "fmt"

// 删除排序数组中的重复项
func rmDup(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nil
	}
	var i = 0
	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1, nums[:i+1]
}

func main() {
	var n1 = []int{2, 2, 4, 5}
	fmt.Println(rmDup(n1))
}

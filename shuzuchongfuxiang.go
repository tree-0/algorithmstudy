package main

import "fmt"

/*
m["1"] = 0  => true
m["2"] = 0  => truetrue
m["3"] = 0  => true =>>下一个m["3"]的值变为true
m["3"] = 0  =>可以取到值true 返回true

*/
func conDup(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

func main() {
	n1 := []int{1, 2, 3, 3}
	fmt.Println(conDup(n1))
}

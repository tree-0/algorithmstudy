package main

import "fmt"

func rev(s []byte) []byte{
	left := 0
	right := len(s) -1
	for left < right {
		s[left],s[right] = s[right],s[left]
		left++
		right--
	}
	return s
}

func main() {
	 n1 := []byte{'a','b','c'}
	 fmt.Println(rev(n1))
}

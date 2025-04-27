package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("p: %p\n", &s)
	clear(s)
	fmt.Printf("len: %d, cap: %d, p: %p, s:%v\n", len(s), cap(s), &s, s)

	a := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("p: %p\n", &a)
	// clear 切片, 原始数组也会清空
	clear(a[:])
	fmt.Printf("len: %d, cap: %d, p: %p, s:%v\n", len(a), cap(a), &a, a)

	s2 := make([]int, 5, 10)
	s2 = append(s2, 1, 2, 3)
	fmt.Printf("p: %p\n", &s2)
	clear(s2)
	fmt.Printf("len: %d, cap: %d, p: %p, s:%v\n", len(s2), cap(s2), &s2, s2)
}

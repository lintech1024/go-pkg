package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println("min:", min(4.1, 3.14, 9.9, 6))
	fmt.Println("max:", max(4.1, 3.14, 9.9, 6))

	fmt.Println("min:", min(4.1, 3.14, 9.9, 6, math.NaN())) // NaN 参与比较, 返回 NaN
	fmt.Println("max:", max(4.1, 3.14, 9.9, 6, math.NaN())) // NaN 参与比较, 返回 NaN

	a := [...]float64{4.1, 3.14, 9.9, 6}
	s := []float64{4.1, 3.14, 9.9, 6}
	// fmt.Println("min:", min(a...))           // 没法比较
	// fmt.Println("min:", min(a[0], a[1:]...)) // 没法比较

	// 只能用 slices 库
	fmt.Println("min:", slices.Min(a[:]))
	fmt.Println("min:", slices.Min(s))
	fmt.Println("max:", slices.Max(a[:]))
	fmt.Println("max:", slices.Max(s))
}

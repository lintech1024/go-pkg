package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	b  := make([]byte, 16)
	n, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Printf("%x\n", b)
}

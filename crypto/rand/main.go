package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	i, err := rand.Prime(rand.Reader, 512)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}

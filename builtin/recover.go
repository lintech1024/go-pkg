package main

import (
	"fmt"
	"runtime"
)

func main() {
	Go(func() {
		panic(nil)
	})
	Go(func() {
		panic("hello")
	})
}

func Go(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover: %s\n", r)
			for i := 0; ; i++ {
                pc, file, line, ok := runtime.Caller(i)
                if !ok {
                    break
                }
                fmt.Printf("Stack[%d]: %s:%d %s\n", i, file, line, runtime.FuncForPC(pc).Name())
            }
		} else {
			fmt.Println("no panic")
		}
	}()

	f()
}

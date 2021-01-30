package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
}
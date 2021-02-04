package main

import (
	// "fmt"
	"os"
	// "log"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		// log.Println("err happened", err)
		// log.Fatalln(err)
		panic(err)
	}
}
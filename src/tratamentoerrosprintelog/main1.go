package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	flog, err := os.Create("logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer flog.Close()
	log.SetOutput(flog)

	f, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println("err happened", err)
		log.Println("err happened", err)
		// log.Fatalln(err)
		// panic(err)
	}
	defer f.Close()

	fmt.Println("check the logs.txt file in the directory")
}
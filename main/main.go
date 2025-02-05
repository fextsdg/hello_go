package main

import (
	"HELLO_GO/test"
	"fmt"
)

func init() {
	// main()
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	fmt.Println(test.A)
}

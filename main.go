package main

import (
	"fmt"
)

func main() {
	var hello float32
	hello = 12.2
	fmt.Println(hello)
	for i := 0; i < 10000000; i++ {
		fmt.Println(i)
	}
}

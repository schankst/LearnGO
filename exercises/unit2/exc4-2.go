package main

import (
	"fmt"
)

func main() {

	a := 42
	fmt.Printf("decimal \t \t %d\n", a)
	fmt.Printf("binary  \t \t %b\n", a)
	fmt.Printf("hex  \t \t %#x\n", a)

	b := a << 1

	fmt.Printf("decimal \t \t %d\n", b)
	fmt.Printf("binary  \t \t %b\n", b)
	fmt.Printf("hex  \t \t %#x\n", b)

}

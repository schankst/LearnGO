package main

import (
	"fmt"
)

func main() {

	x := 42
	switch {
	case x < 10:
		fmt.Println("x < 10")
	case x < 100:
		fmt.Println("x>10 and x<100")
	case x > 100:
		fmt.Println("x>100")

	}
}

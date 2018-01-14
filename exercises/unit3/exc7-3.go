package main

import (
	"fmt"
)

func main() {

	x := 42

	if x < 10 {
		fmt.Println("x < 10")
	} else if x < 100 {
		fmt.Println("x>10 and x<100")
	} else {
		fmt.Println("x>100")
	}

}

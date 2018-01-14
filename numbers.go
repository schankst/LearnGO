package main

import "fmt"

func main() {

	for x := 33; x <= 122; x++ {

		fmt.Printf("decimal value %d, hex value %#x, unicode value %#U \n", x, x, x)
	}
}

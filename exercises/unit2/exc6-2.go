package main

import (
	"fmt"
)

const (
	a = 2014
	b = a + iota
	c = a + iota
	d = a + iota
)

func main() {

	fmt.Println(a, b, c, d)
}

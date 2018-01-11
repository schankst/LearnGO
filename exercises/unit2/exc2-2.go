package main

import (
	"fmt"
)

func main() {

	a := 2
	b := 4
	c := 2 == 4
	d := 2 <= 4
	e := 2 >= 4
	f := 2 != 4
	g := 2 < 4
	h := 2 > 4

	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)
	fmt.Printf("%T \t %v \n", c, c)
	fmt.Printf("%T \t %v \n", d, d)
	fmt.Printf("%T \t %v \n", e, e)
	fmt.Printf("%T \t %v \n", f, f)
	fmt.Printf("%T \t %v \n", g, g)
	fmt.Printf("%T \t %v \n", h, h)

}

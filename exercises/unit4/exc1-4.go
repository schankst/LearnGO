package main

import (
	"fmt"
)

func main() {

	var x [5]int
	//var y []uint

	y := []int{6, 7, 8, 2, 1, 3}

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(len(x))
	fmt.Println(cap(y))

	for i, v := range y {
		fmt.Println(i, v)
	}

	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}

	y = append(y, 67, 78, 73, 13, 67, 78)
	
	z:= y
	
	z append(z, y...)
	fmt.Println(z[:3])
	fmt.Println(z[4:])

}

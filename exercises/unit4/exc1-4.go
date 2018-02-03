package main

import (
	"fmt"
)

func main() {

	var x [5]int
	fmt.Printf("Array x %v\n", x)

	y := []int{6, 7, 8, 2, 1, 3}
	fmt.Printf("Slize y %v\n", y)

	// fmt.Println(x)
	// fmt.Println(y)

	// fmt.Println(len(x))
	// fmt.Println(cap(y))

	// for i, v := range y {
	// 	fmt.Println(i, v)
	// }

	// for i := 0; i < len(y); i++ {
	// 	fmt.Println(i, y[i])
	// }

	y = append(y, 67, 78, 73, 13, 67, 78)
	fmt.Printf("Slize y %v\n", y)

	z := y

	z = append(z, y...)
	fmt.Printf("Slize z %v\n", z)

	fmt.Println(z[:3])
	fmt.Println(z[4:])

}

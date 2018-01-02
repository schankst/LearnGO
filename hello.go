package main

import "fmt"

var a int64

func z() {
	println(a)
}

func main() {
	c := 0

	c, _ = fmt.Print("Hello World")
	a = 10
	fmt.Println(a, c)
	b := "Hello new world"
	fmt.Println(b)
}

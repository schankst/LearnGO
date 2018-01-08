package main

//import fmt package
import "fmt"

//declare some variables
var a int64
var y string

// func z has no parameters
func z() {
	y = ` hello "some text" printed out`
	fmt.Println(a + 1)
	fmt.Println(y)
}

//declare a new variable
var c int

//create your own datatype with base type int
type mytype int

//declare a new variable based on you own datatype
var m mytype

//main function, here starts the program
func main() {
	//c := 42

	m := 43
	fmt.Println(m)
	c, _ = fmt.Println("Hello World")
	a = 12
	fmt.Println(a, c)
	b := "Hello new world"
	s := fmt.Sprintf("%T\n", b)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", b)
	fmt.Printf("%T", s)
	//call function z()
	z()
	z()
}

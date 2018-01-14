package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().Local())
	fmt.Println(time.Now().Location())
	fmt.Println(time.Now().Zone())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

}

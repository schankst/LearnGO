package main

import (
	"fmt"
)

func main() {
	bd := 1972
	for {
		if bd != 2018 {
			fmt.Println(bd)
			bd++
		} else {
			break
		}
	}
}

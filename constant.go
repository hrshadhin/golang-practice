package main

import (
	"fmt"
)

func main() {

	const PI = 3.1416
	const (
		zero int = iota
		one
		two
		three
		four
	)

	fmt.Println("Pi ", PI)
	fmt.Println(zero)
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
	fmt.Println(four)
}

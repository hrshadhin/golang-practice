package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please give one or more integer/floats.")
		os.Exit(1)
	}

	arguments := os.Args
	sum := 0.00
	for i := 1; i < len(arguments); i++ {
		number, _ := strconv.ParseFloat(arguments[i], 64)
		sum = sum + number
	}

	fmt.Printf("Sum of all arguments is %.2f", sum)

}

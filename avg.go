package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args
	sum := 0.00
	numbersCount := 0
	for i := 1; i < len(arguments); i++ {
		number, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += number
			numbersCount++
		}

	}

	avg := (sum / float64(numbersCount))

	fmt.Printf("Avg of all aruments is %.2f\n", avg)
}

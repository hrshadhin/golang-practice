package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)

	fmt.Println("Type integer or type STOP to stop the program.")

	for scanner.Scan() {
		userInputs := scanner.Text()
		intNumber, err := strconv.ParseInt(userInputs, 10, 64)

		if err != nil {
			if userInputs == "STOP" {
				break
			} else {
				fmt.Println("Type STOP to stop the program!")
			}
		} else {
			fmt.Println(intNumber)
		}
	}
}

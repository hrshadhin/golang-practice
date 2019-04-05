package main

// #include <stdio.h>
// void callC() {
// 	printf("I am from C code.\n");
// }
import "C"

import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}

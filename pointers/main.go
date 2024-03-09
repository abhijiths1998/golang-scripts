package main

import "fmt"

func ptrCheck() {
	word := "Blasphemous"

	var wordMem = &word

	fmt.Println("Memory address for the pointer : ", wordMem)
	fmt.Println("Actual Value for the pointer : ", *wordMem)
}

func main() {
	ptrCheck()
}

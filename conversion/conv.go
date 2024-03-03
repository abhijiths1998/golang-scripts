package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
* Get multiple inputs and add all of them
**/

func getInput() {
	fmt.Println("Welcome to the program")
	fmt.Println("Now you may enter a value : ")

	reader := bufio.NewReader(os.Stdin)
	strVal, _ := reader.ReadString('\n')
	numVal, err := strconv.ParseFloat(strings.TrimSpace(strVal), 64)

	fmt.Println("Enter another value : ")

	newReader := bufio.NewReader(os.Stdin)
	secVal, _ := newReader.ReadString('\n')

	secNumVal, err2 := strconv.ParseFloat(strings.TrimSpace(secVal), 64)

	if err != nil {
		fmt.Printf("Error in first value %s \n", err)
	} else if err2 != nil {
		fmt.Printf("Error in Second value %s \n", err2)
	} else if err == nil && err2 == nil {
		fmt.Println("All Values are present !!")
		finalVal := numVal + secNumVal
		fmt.Printf("Final Value is %f \n", finalVal)
	}

}
func main() {
	getInput()
}

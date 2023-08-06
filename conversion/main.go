package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Function USER INPUT
/*
Get 2 input as a string and added string conversion using strconv.ParseInt.

In this case we have taken Marks as an example.

Taken exam and internal marks as a string input and find the total marks using to typecasting string to integer.
*/
func userInput() {
	println("Welcome to the student portal")
	println("Please enter your exam marks (1-80) : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	println("Please enter your Internal's marks (1-20) : ")
	readerInt := bufio.NewReader(os.Stdin)
	inputInt, _ := readerInt.ReadString('\n')
	examMarks, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	internalMarks, err2 := strconv.ParseInt(strings.TrimSpace(inputInt), 10, 64)

	if err != nil || err2 != nil {
		println(err.Error())
		println((err2.Error()))
	} else {
		println("Your Internal mark is : ", internalMarks)
		println("Your Exam mark is : ", examMarks)
		println("Your total mark is : ", internalMarks+examMarks)
	}

}

func main() {
	userInput()
}

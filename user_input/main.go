package main

import (
	"bufio"
	"os"
)

// User Input

/*
This function is used for demonstrating how user input works.

First it will create a readstream using bufio library and save this as a variable.

Then you call the function using ReadString method and mention when the buffer should end..
In this case when a new line comes (\n) the buffer will end.

We printed the same showing the rating.
*/
func userInput() {
	reader := bufio.NewReader(os.Stdin)
	println("Welcome to Dominos")
	println("Enter your rating for the pizza : ")
	input, _ := reader.ReadString('\n')
	println("Thanks for ordering from dominos. You gave a rating of -> ", input)
}

func main() {
	userInput()
}

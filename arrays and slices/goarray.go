package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var nameList = [4]string{"abhi", "ram", "jaadhu", "me"}

	fmt.Println("Name list : ", nameList)
	fmt.Println("Length of Name list : ", len(nameList))

	var nameList2 [3]string

	nameList2[0] = "deep"
	nameList2[1] = "random"
	nameList2[2] = "wow"

	fmt.Println("2nd name list : ", nameList2)
	fmt.Println("2nd name list : ", len(nameList2))

}

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Different ways to create slices

	//method 1

	var slice1 = []string{}

	slice1 = append(slice1, "adding", "random", "shit")

	fmt.Println("Current Slice1 : ", slice1)

	slice1 = append(slice1, "adding", "more", "random", "shit")

	fmt.Println("current slice1 : ", slice1)

	//method 2

	var slice2 = make([]float64, 4)

	slice2[0] = 45.67
	slice2[1] = 46.75
	slice2[2] = 34.56
	slice2[3] = 34.56

	fmt.Println("Current slice2 : ", slice2)

	slice2 = append(slice2, 56.78, 23, 45, 10.23, 24.5, 34.2)

	fmt.Println("Current slice2 : ", slice2)

	fmt.Println("Is this sorted ?? \n ", sort.Float64sAreSorted(slice2))

	isSorted := sort.Float64sAreSorted(slice2)

	if isSorted {
		fmt.Println("Slice is already sorted !!")
	} else {
		sort.Float64s(slice2)
		fmt.Println("Is this sorted now ?? \n ", sort.Float64sAreSorted(slice2))
		fmt.Println("Length of the slices \n ", len(slice2))
		fmt.Println("Sorted slices : ", slice2)

	}

	//learn to remove a value from the index

	var removeSlice = []string{"a", "b", "c", "d"}

	removeSlice = append(removeSlice, "d", "e", "f", "g")

	fmt.Println("Current list is : ", removeSlice)

	var index int = 4

	removeSlice = append(removeSlice[:index], removeSlice[index+1:]...)

	fmt.Println("Final list : ", removeSlice)

}

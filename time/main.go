package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func timeTravel() {
	fmt.Println("Welcome to time travel")

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	fmt.Println("Current time is : ", formattedTime)

	fmt.Println("How much hours you want to add/subtract? ")
	reader := bufio.NewReader(os.Stdin)
	timeVal, _ := reader.ReadString('\n')

	timeValNew, err := strconv.ParseFloat(strings.TrimSpace(timeVal), 64)

	newTime := currentTime.Add(time.Hour * time.Duration(timeValNew))
	formattedNewTime := newTime.Format("2006-01-02 15:04:05")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Original Time : ", formattedTime, "\n\nNew Time : ", formattedNewTime)
	}

}

func main() {
	timeTravel()
}

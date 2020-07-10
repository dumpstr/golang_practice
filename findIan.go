package main

import (
	"fmt"
	"strings"
)

func main() {
	var usrInput string
	var flagI, flagA, flagN int
	fmt.Printf("Please input a string.\n")

	x, err := fmt.Scan(&usrInput)
	_, _ = x, err

	convStr := strings.ToUpper(usrInput)

	if strings.HasPrefix(convStr, "I") == true {
		flagI = 1
	}
	if strings.HasSuffix(convStr, "N") == true {
		flagN = 1
	}
	if strings.ContainsAny(convStr, "A") == true {
		flagA = 1
	}

	if flagI == 1 && flagA == 1 && flagN == 1 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

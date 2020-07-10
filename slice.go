package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var usrInput string
	sli := make([]int, 3, 10)
	sli2 := make([]int, 3, 10)

	for i := 0; i < 9; i++ {

		fmt.Printf("Please input a number to add to the slice. Input x to exit.\n")

		x, err := fmt.Scan(&usrInput)
		_ = x
		if err != nil {
			panic(err)
		}

		if usrInput == "X" || usrInput == "x" {
			break
		}

		convStr, err := strconv.Atoi(usrInput)

		if err != nil {
			panic(err)
		}

		switch {
		case i > 2:
			sli = append(sli, convStr)
			sli2 = append(sli2, convStr)
		default:
			sli[i] = convStr
		}

		clone1 := copy(sli2, sli)
		_ = clone1
		sort.Slice(sli2, func(p, q int) bool {
			return sli2[p] < sli2[q]
		})

		fmt.Println("Loop:", i)
		fmt.Println("Slice Ascending:", sli2)
	}

	sort.Slice(sli, func(p, q int) bool {
		return sli[p] < sli[q]
	})

	fmt.Println("Final Ordered (ascending):", sli)
}

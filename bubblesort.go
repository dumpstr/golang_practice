package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func BubbleSort(s []int) {
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j-1] > s[j] {
				Swap(s, j)
			}
		}
	}
}

func Swap(b []int, q int) {
	swap := b[q]
	b[q] = b[q-1]
	b[q-1] = swap
}

func main() {
	sli := make([]int, 10)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter 10 integers to be sorted.")
	for p := 0; p <= 9; p++ {
		scanner.Scan()
		holder := scanner.Text()
		conv, err := strconv.Atoi(holder)
		sli[p] = conv
		if err != nil {
			panic(err)
		}
	}
	BubbleSort(sli)
	fmt.Println(sli)
}

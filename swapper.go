package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func BubbleSort(s []int, wg *sync.WaitGroup) {
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j-1] > s[j] {
				Swap(s, j)
			}
		}
	}
	wg.Done()
}

func Swap(b []int, q int) {
	swap := b[q]
	b[q] = b[q-1]
	b[q-1] = swap
}

func main() {
	var sli []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter some integers (separated by whitespace) to be sorted.")
	scanner.Scan()
	holder := strings.Fields(scanner.Text())
	for _, i := range holder {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		sli = append(sli, j)
	}
	quarter := len(sli) / 4
	wg := sync.WaitGroup{}
	wg.Add(4)
	fmt.Println("Subarray to be sorted: ", sli[:quarter])
	fmt.Println("Subarray to be sorted: ", sli[quarter:quarter*2])
	fmt.Println("Subarray to be sorted: ", sli[quarter*2:quarter*3])
	fmt.Println("Subarray to be sorted: ", sli[quarter*3:])
	go BubbleSort(sli[:quarter], &wg)
	go BubbleSort(sli[quarter:quarter*2], &wg)
	go BubbleSort(sli[quarter*2:quarter*3], &wg)
	go BubbleSort(sli[quarter*3:], &wg)
	wg.Wait()
	wg1 := sync.WaitGroup{}
	wg1.Add(1)
	BubbleSort(sli, &wg1)
	wg1.Wait()

	fmt.Println("Sorted array: ", sli)
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (.5 * a * math.Pow(t, 2)) + (v0 * t) + s0
	}
	return fn
}

func main() {
	sli := make([]float64, 4)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i <= 3; i++ {
		if i == 0 {
			fmt.Println("Enter acceleration (m^2/s).")
		}
		if i == 1 {
			fmt.Println("Enter initial velocity (m/s).")
		}
		if i == 2 {
			fmt.Println("Enter initial displacement (m).")
		}
		if i == 3 {
			fmt.Println("Enter time (s).")
		}
		scanner.Scan()
		holder := scanner.Text()
		conv, err := strconv.ParseFloat(holder, 64)
		sli[i] = conv
		if err != nil {
			panic(err)
		}
	}
	Disp := GenDisplaceFn(sli[0], sli[1], sli[2])
	fmt.Printf("Displacement for %v second(s): %v m", sli[3], Disp(sli[3]))
}

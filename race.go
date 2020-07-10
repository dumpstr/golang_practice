/*
A race condition is a situation where the outcome of the program is dependent on the interleavings.
This means that the program is nondeterministic. The two goroutines that I have written interates an int (x = 1) to 6 and multiplies it to 125000.
The race condition that occurs is that sometimes the first printed line is 6 then 125000 and other times it is flipped.
The reason this happens is due to how the goroutines/threads are executed by the scheduler. The interleaving of the instructions are nondeterministic.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	go func(x int) int {
		for i := 0; i < 5; i++ {
			x = x + 1
		}
		fmt.Println(x)
		return x
	}(x)

	go func(x int) int {
		for i := 0; i < 3; i++ {
			x = x * 50
		}
		fmt.Println(x)
		return x
	}(x)

	time.Sleep(time.Second)
	fmt.Println("Finished.")
}

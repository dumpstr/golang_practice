package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("wooper")
	case 1:
		fmt.Println("dooker")
	case 3:
		fmt.Println("dooper")
	default:
		fmt.Println("boop #", num)
	}
}

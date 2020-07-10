package main

import "fmt"

func main() {
	var inputFlt float32
	fmt.Printf("Please input a floating point number.\n")

	flt, err := fmt.Scan(&inputFlt)
	_, _ = flt, err
	convert := int(inputFlt)
	fmt.Println(convert)
}

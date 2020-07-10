package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	fname string
	lname string
}

func main() {
	var people []person
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input names file to be read.")
	x := scanner.Scan()
	_ = x
	nameInput := scanner.Text()
	file, err := os.Open(nameInput)
	check(err)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
	for fileScanner.Scan() {
		firstName := fileScanner.Text()
		if !fileScanner.Scan() {
			break
		}
		lastName := fileScanner.Text()
		p := person{firstName, lastName}
		people = append(people, p)

		fmt.Printf("first name: %v, last name: %v\n", firstName, lastName)
	}

	fmt.Println(people)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

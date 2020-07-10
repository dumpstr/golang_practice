package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var nameInput string
	var addressInput string
	idMap := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input a name.\n")

	x := scanner.Scan()
	nameInput = scanner.Text()

	fmt.Printf("Please input an address.\n")

	y := scanner.Scan()
	addressInput = scanner.Text()

	if x == false || y == false {
		panic("A problem occured. Invalid input.")
	}
	idMap[nameInput] = addressInput

	//turn into json output
	bmap, err := json.Marshal(idMap)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bmap))
}

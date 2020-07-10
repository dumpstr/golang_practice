package main

import (
	"fmt"
	"math/rand"

	"github.com/kniren/gota/dataframe"
)

const secondsInDay = 86400

type Ticket struct {
	spaceline string
	days      int
	tripType  string
	price     int
}

func main() {
	distance := 62100000
	company := ""
	trip := ""

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space X"
		case 1:
			company = "Space Adventures"
		case 2:
			company = "Virgin Galactic"
		}
	}

	speed := rand.Intn(15) + 16
	duration := distance / speed / secondsInDay
	price := 20.0 + speed

	if rand.Intn(2) == 1 {
		trip = "Round Trip"
		price *= 2
	} else {
		trip = "One Way"
	}
	tickets := []Ticket{
		Ticket{company, duration, trip, price},
	}
	df := dataframe.LoadStructs(tickets)
	fmt.Println(df)
}

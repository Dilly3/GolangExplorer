package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type Pizzaorder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	// print out message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("-----------------------------")

}

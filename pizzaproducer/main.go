package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfOrders = 15

var pizzasMade, pizzasFailed, total int

// Producer is a type for structs that holds two channels: one for pizzas, with all
// information for a given pizza order including whether it was made
// successfully, and another to handle end of processing (when we quit the channel)
type Producer struct {
	data chan PizzaBox
	quit chan string
}

// PizzaOrder is a type for structs that describes a given pizza order. It has the
// *** order number
// *** a message indicating what happened to the order,
// *** and a boolean indicating if the order was successfully completed.
type PizzaBox struct {
	pizzaNumber int
	message     string
	success     bool
}

// Close is simply a method of closing the channel when we are done with it (i.e.
// something is pushed to the quit channel)
func (p *Producer) Close() error {

	p.quit <- "channel closed"
	return nil
}

// makePizza attempts to make a pizza. We generate a random number from 1-12,
// and put in two cases where we can't make the pizza in time. Otherwise,
// we make the pizza without issue. To make things interesting, each pizza
// will take a different length of time to produce (some pizzas are harder than others).
func makePizza(pizzaNumber int) *PizzaBox {

	if pizzaNumber <= NumberOfOrders {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}
		pizzaNumber++
		p := PizzaBox{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p

	}

	return &PizzaBox{
		pizzaNumber: pizzaNumber,
	}
}

// pizzeria is a goroutine that runs in the background and
// calls makePizza to try to make one order each time it iterates through
// the for loop. It executes until it receives something on the quit
// channel. The quit channel does not receive anything until the consumer
// sends it (when the number of orders is greater than or equal to the
// constant NumberOfPizzas).
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 1

	// this loop will continue to execute, trying to make pizzas,
	// until the quit channel receives something.
	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza (we send something to the data channel -- a chan PizzaOrder)
			case pizzaMaker.data <- *currentPizza:

			// we want to quit, so send pizzMaker.quit to the quitChan (a chan error)
			case <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(pizzaMaker.quit)
				return
			}
		}
	}
}

func main() {
	// seed the random number generator
	//rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaBox),
		quit: make(chan string),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfOrders {
			if i.success {
				color.HiGreen(i.message)
				color.HiGreen("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.HiRed(i.message)
				color.HiRed("The customer is really mad!")
			}
		} else {
			color.HiCyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.HiRed("*** Error closing channel!", err)
			}
		}
	}

	// print out the ending message
	color.HiCyan("-----------------")
	color.HiCyan("Done for the day.")

	color.HiCyan("We made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.HiRed("It was an awful day...")
	case pizzasFailed >= 6:
		color.HiRed("It was not a very good day...")
	case pizzasFailed >= 4:
		color.HiYellow("It was an okay day....")
	case pizzasFailed >= 2:
		color.HiYellow("It was a pretty good day!")
	default:
		color.HiGreen("It was a great day!")
	}
}

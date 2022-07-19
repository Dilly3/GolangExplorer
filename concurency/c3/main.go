package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go greet(ch)
	time.Sleep(5 * time.Second)
	fmt.Println("main ready")
	greeting := <-ch
	time.Sleep(2 * time.Second)
	fmt.Println("greeter received greeting:", greeting)
	fmt.Println(greeting)
}
func greet(ch chan string) {
	fmt.Printf("Greeter ready!...\nGreeter waiting to send greeting...\n")
	ch <- "Hello, world!"
	fmt.Println("Greeter sent greeting!")
}

package main

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func main() {
	count := make(chan int)
	name := make(chan string, 4)

	Wg.Add(4)
	go counter("A", count)
	go counter("B", count)
	go counter("C", count)
	go register(name)
	count <- 1

	Wg.Wait()

	for {
		val, ok := <-name
		if !ok {
			fmt.Println(val)
			return
		}
		fmt.Println(val)
	}

}
func register(str chan string) {
	defer Wg.Done()
	names := []string{"michael", "zubby", "kenneth", "fredrick"}
	for _, val := range names {
		str <- val
	}
	close(str)
}
func counter(a string, ch chan int) {

	defer Wg.Done()
	for {
		val, ok := <-ch

		if !ok {
			println(a, "closed")
			return
		}

		println(a, val)

		if val == 12 {
			close(ch)
			return
		}

		val++

		ch <- val
	}
}

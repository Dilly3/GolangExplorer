package main

import (
	"fmt"
	"strconv"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	wg.Add(len(words))
	for i := 0; i < len(words); i++ {
		go printSomething(fmt.Sprintf("%s : %s", strconv.Itoa(i+1), words[i]), &wg)
	}
	wg.Wait()
}

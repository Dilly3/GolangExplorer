package main

import "fmt"
import G "goexplorer/Algo/Week5"

func main() {

	fmt.Println(G.CaesarCyp("One should not worry over things that have already happened and that cannot be changed.", 49))

}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

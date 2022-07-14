package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

var Wg = &sync.WaitGroup{}

func main() {

	Wg.Add(1)
	go echo(os.Stdin, os.Stdout)

	fmt.Println("outside of a goroutine")
	go func() {
		fmt.Println("inside of a goroutine")
	}()
	fmt.Println("outside of a goroutine again")
	runtime.Gosched()
	time.Sleep(20 * time.Second)
	Wg.Done()
	Wg.Wait()
	fmt.Println("timed out")
	os.Exit(0)

}
func echo(in io.Reader, out io.Writer) {

	n, _ := io.Copy(out, in)

	fmt.Println("copied", n, "bytes")

}

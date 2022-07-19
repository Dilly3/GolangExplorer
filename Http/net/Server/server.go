package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

var CallerGroup = make(map[net.Conn]bool)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	connectionChannel := make(chan net.Conn)
	listener, lErr := net.Listen("tcp", ":9090")
	HandleErr(lErr)
	defer listener.Close()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for {
			caller, err := listener.Accept()
			HandleErr(err)
			CallerGroup[caller] = true
			connectionChannel <- caller
		}

	}()
	for {

		reader := bufio.NewReader(<-connectionChannel)
		message, Merr := reader.ReadSlice('\n')
		HandleErr(Merr)

		fmt.Println(string(message))

		wg.Wait()
	}
}

package Server

import (
	"fmt"
	"log"
	"net"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func StartServer() {
	li, err1 := net.Listen("tcp", ":8080")
	HandleErr(err1)
	defer li.Close()

	for {
		conn, err2 := li.Accept()
		HandleErr(err2)

		_, err3 := fmt.Fprintf(conn, "Welcome to Ohdy Network")
		HandleErr(err3)
	}
}

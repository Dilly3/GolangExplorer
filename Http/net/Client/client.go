package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	caller, connErr := net.Dial("tcp", ":9090")
	HandleErr(connErr)
	defer caller.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is Your Name")
	username, readErr := reader.ReadString('\n')
	HandleErr(readErr)
	username = strings.Trim(username, "\n ")
	username = fmt.Sprintf("%s Connected \n", username)

	caller.Write([]byte(username))

}

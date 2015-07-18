package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	port := 8000
	fmt.Println("Launching server onto port ", port)
	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(port))

	connection, _ := listener.Accept()

	for {
		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print(string(message))
		newMessage := strings.ToUpper(message)
		connection.Write([]byte(newMessage + "\n"))
	}
}

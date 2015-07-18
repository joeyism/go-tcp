package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	port := 8000
	connect, _ := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(port))
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(connect, text+"\n")
		message, _ := bufio.NewReader(connect).ReadString('\n')
		fmt.Print("Response: " + message)
	}
}

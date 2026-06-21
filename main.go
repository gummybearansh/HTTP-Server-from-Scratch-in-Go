package main

import (
	"bufio"
	"fmt"
	"net"
	// "strconv"
	// "strings"
)

func main() { 
	// Listen on port -> returns err, Listener 
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer listener.Close()


	for { 
		// wait for connection
		connection, err := listener.Accept()
		if err != nil { 
			fmt.Println(err)
			return;
		}

		// handle the connection in a goroutine 
		// the loop returns to Accepting - so multiple connections can be served concurrently
		go func (connection net.Conn){
			// pass the TCP connection to bufio to create a reader on it
			reader := bufio.NewReader(connection)
			// Entire Parser in parser.go 
			_, err := ParseRequest(reader, connection)
			if err != nil {
				SendResponse(connection, 400, "BAD REQUEST", "", err.Error())
				return
			}

			// close connection
			connection.Close()
		} (connection)

	}
}

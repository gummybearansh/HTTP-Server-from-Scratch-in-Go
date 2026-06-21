package main

import (
	"bytes"
	"fmt"
	"net"
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
			// read upto 1000 bytes of request - headers etc
			buffer := make([]byte, 1000)
			count, err := connection.Read(buffer)
			if err != nil { 
				fmt.Println(err)
				return
			}
			// need to parse the Request line [METHOD] [PATH] [PROTOCOL]
			// to understand the lgoic 
			request_line, _, found := bytes.Cut(buffer[:count], []byte("\n"))
			if !found {
				fmt.Println("Invalid request")
				return
			}
			// fmt.Println(string(request_line))

			method, remaining, found := bytes.Cut(request_line, []byte("/"))
			if !found {
				fmt.Println("Invalid request")
				return
			}
			path, protocol, found := bytes.Cut(remaining, []byte(" "))
			if !found  {
				fmt.Println("Invalid request")
				return;
			}

			fmt.Println("method", string(method))
			fmt.Println("path", string(path))
			fmt.Println("protcol", string(protocol))


			// net.Conn is 2 way - can send back response from here 
			// HTTP resposne format: (every line in the header sectiion terminates with carriage return + newline)
			status_line := "HTTP/1.1 200 OK\r\n"
			headers := "Content-type: text/html\r\n" 
			empty_line := "\r\n" // 2x (one from previous header) exact sequence of carriage return & newline - physical barrier between metadata and body
			body := "<h1>Hello world</h1>"

			response := status_line + headers + empty_line + body
			b := []byte(response)

			// send response
			write_count, err := connection.Write(b)
			if err != nil { 
				fmt.Println(err)
				return;
			}
			fmt.Println(response[:write_count])

			// close connection
			connection.Close()
		} (connection)

	}
}

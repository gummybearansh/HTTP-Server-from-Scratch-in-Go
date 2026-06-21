package main

import (
	// "bytes"
	"bufio"
	"fmt"
	"net"
	"strings"
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
			// need to parse the Request line [METHOD] [PATH] [PROTOCOL]
			// read all the headers 
			var request_headers []string
			for { 
				line, err := reader.ReadString('\n')
				if err != nil { 
					fmt.Println(err)
					return;
				}
				// last line will be just this
				if line == "\r\n" { 
					break;
				}
				request_headers = append(request_headers, line)
			}

			if len(request_headers) == 0 {
				fmt.Println("Invalid request")
				return;
			}

			request_status_line := request_headers[0]
			// status line format 
			// [METHOD] <space> [PATH] <space> [PROTOCL]
			// splits on whitespaces
			status_line_array := strings.Fields(request_status_line)
			if len(status_line_array) < 3 {
				fmt.Print("Invalid request")
				return;
			}
			method := status_line_array[0]
			path := status_line_array[1]
			protocol := status_line_array[2]

			fmt.Println("method", method)
			fmt.Println("path", path)
			fmt.Println("protocol", protocol)

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
				fmt.Print(write_count)
				return;
			}
			// fmt.Println(response[:write_count])

			// close connection
			connection.Close()
		} (connection)

	}
}

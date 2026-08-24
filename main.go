package main

import (
	"bufio"
	"fmt"
	"net"
	// this is the module path name, but the package name is just limiter so i can use it with limiter.check
	"github.com/gummybearansh/token-bucket-limiter"
)

func main() { 
	// Listen on port -> returns err, Listener 
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer listener.Close()

	rateLimiter := limiter.NewLimiter(1.0, 3.0); // 1 token per second, max burst 3

	for { 
		// wait for connection
		connection, err := listener.Accept()
		if err != nil { 
			fmt.Println(err)
			return;
		}

		// handle the connection in a goroutine 
		// the loop returns to Accepting - so multiple connections can be served concurrently
		go HandleConnection(connection, rateLimiter)
	}
}

func HandleConnection(connection net.Conn, rateLimiter *limiter.RateLimiter) {
	// extract IP of the request 
	ip, _, err := net.SplitHostPort(connection.RemoteAddr().String())
	if err != nil {
		// if extraction fails, fallback to the raw string 
		ip = connection.RemoteAddr().String()
	}

	// rate limiting 
	if !rateLimiter.Allow(ip){
		// Write the raw HTTP 429 response
		SendResponse(connection, 429, "Too Many Requests", "Content-Length: 17", "Too Many Requests")
    connection.Close() // Terminate the TCP socket
    return       // Kill the Goroutine immediately
	}

	// pass the TCP connection to bufio to create a reader on it
	reader := bufio.NewReader(connection)
	// Entire Parser in parser.go 
	request, err := ParseRequest(reader, connection)
	if err != nil {
		// net.Conn is 2 way - can send and receive on the same connection
		Handle400(request, connection)
		return
	}

	RouteRequest(request, connection)
	// close connection
	connection.Close()
}

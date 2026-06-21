package main

import (
	"errors"
	"fmt"
	"net"
)

func SendResponse(connection net.Conn, status_code int, status_text, headers string, body string) (bool, error){
	status_line := fmt.Sprintf("HTTP/1.1 %d %s\r\n", status_code, status_text)
	empty_line := "\r\n" // 2x (one from previous header) exact sequence of carriage return & newline - physical barrier between metadata and body
	
	response := status_line + headers + empty_line + body
	// convert it to byte stream
	b := []byte(response)

	// send response
	_, err := connection.Write(b)
	if err != nil { 
		return false, errors.New("Error sending response")
	}

	// successful request sent
	return true, nil
}

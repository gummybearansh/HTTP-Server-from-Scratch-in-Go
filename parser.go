package main

import (
	"bufio"
	"net"
	"strings"
	"strconv"
	"errors"
)

// this is the type that can be initialised 
type HTTPRequest struct {
	Method string 
	Path string 
	Protocol string 
	// map[keyType]valueType
	Headers map[string]string
	Body string
}

func ParseRequest(reader *bufio.Reader, connection net.Conn) (*HTTPRequest, error){
	var request HTTPRequest
	request.Headers = make(map[string]string)

	// first line is the request line
	// [METHOD] <space> [PATH] <space> [PROTOCL]
	request_status_line, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Invalid Request line")
	}
	request_status_line = strings.TrimSpace(request_status_line)
	status_line_array := strings.Fields(request_status_line)
	if len(status_line_array) < 3 {
		// can return nil because i'm returning pointer to my struct
		return nil, errors.New("Invalid status line")
	}
	request.Method = status_line_array[0]
	request.Path = status_line_array[1]
	request.Protocol = status_line_array[2]

	// parse all headers
	for { 
		line, err := reader.ReadString('\n')
		if err != nil { 
			return nil, errors.New("Error reading headers")
		}
		// last line will be just this
		if line == "\r\n" { 
			break;
		}
		line = strings.TrimSpace(line)
		// split the header on the :
		headerParts := strings.SplitN(line, ":", 2)
		if len(headerParts) == 2 {
			key := strings.ToLower(strings.TrimSpace(headerParts[0]))
			value := strings.ToLower(strings.TrimSpace(headerParts[1]))

			// add to map
			request.Headers[key] = value
		}
	}

	if len(request.Headers) == 0 {
		return nil, errors.New("No headers received")
	}
	
	// one of the headers will be "Content-Length" - specifies bytes of body
	body_size := -1
	value, exists := request.Headers["content-length"]
	if exists { 
		body_size, err = strconv.Atoi(value)
		if err != nil {
			return nil, errors.New("error parsing content-length")
		}
	}

	// read the body if it exists
	var body_buffer []byte
	if body_size == -1 {
		body_buffer = nil
	} else {
		body_buffer = make([]byte, body_size)
		_, err := reader.Read(body_buffer)
		if err != nil {
			return nil, errors.New("error parsing body")
		}
	}
	request.Body = string(body_buffer)

	// successful parsed response - no errors
	return &request, nil
}

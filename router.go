package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

// values are of a specific function type wowowowowowo
var RouterMap = map[string]func(*HTTPRequest, net.Conn){
	"GET /": HandleHome,
	"GET /about": HandleAbout,
	"POST /": HandleHomePOST,
}

func RouteRequest(request *HTTPRequest, connection net.Conn){
	req := fmt.Sprintf("%s %s", request.Method, request.Path)
	handler, exists := RouterMap[req]
	if !exists {
		Handle404(request, connection)
	} else {
		handler(request, connection)
	}
}

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

func HeadersMapToString(headers_map map[string]string) (string){
	var builder strings.Builder
	for key, value := range headers_map {
		builder.WriteString(key)
		builder.WriteString(":")
		builder.WriteString(value)
		builder.WriteString("\r\n")
	}
	
	return builder.String()
}

func Handle404(HTTPRequest *HTTPRequest, connection net.Conn){
	fmt.Println("Hit 404")
	headers_map := map[string]string{
		"content-type": "text/html",
	}
	headers := HeadersMapToString(headers_map)
	body := "<h1>404 Page Not Found</h1>"
	SendResponse(connection, 404, "Not Found", headers, body)
}

func Handle400(HTTPRequest *HTTPRequest, connection net.Conn){
	fmt.Println("Hit 400")
	headers_map := map[string]string{
		"content-type": "text/html",
	}
	headers := HeadersMapToString(headers_map)
	body := "<h1>400 Bad Request</h1>"
	SendResponse(connection, 400, "Bad Request", headers, body)
}

func HandleHome(HTTPRequest *HTTPRequest, connection net.Conn){
	fmt.Println("hit /")
	headers_map := map[string]string{
		"content-type": "text/html",
	}
	headers := HeadersMapToString(headers_map)
	body := "<h1>HTTP Server From Scratch in Go wow\nCheck out the <a href='/about'>about</a> page for more info.</h1>"
	SendResponse(connection, 200, "OK", headers, body)
}

func HandleAbout(HTTPRequest *HTTPRequest, connection net.Conn){
	fmt.Println("hit /About")
	headers_map := map[string]string{
		"content-type": "text/html",
	}
	headers := HeadersMapToString(headers_map)
	body := "<h1>This entire web Server was written from scratch <br>- handling TCP connection, <br>- Accepting requests, <br>- Parsing the request - from the Status line (METHOD PATH PROTOCOL), the headers, the body all of it.<br>Thank you for visiting.</h1>"
	SendResponse(connection, 200, "OK", headers, body)
}

func HandleHomePOST(HTTPRequest *HTTPRequest, connection net.Conn){
	fmt.Println("Post to /")
	headers_map := map[string]string{
		"content-type": "text/html",
	}
	headers := HeadersMapToString(headers_map)
	body := "<h1>HTTP Server From Scratch in Go wow<br>Also this is what you sent<br><br></h1>"
	fmt.Println(HTTPRequest.Body)
	body += HTTPRequest.Body
	SendResponse(connection, 200, "OK", headers, body)
}




// proxy-trainingexercise-golang
package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

/*
https://pkg.go.dev/net will probably do the most work for me,
user makes requests to url, the proxy forwards them to the url, including headers and
status codes.

*/

//https://www.upguard.com/blog/proxy-server

//my time limit is 2 days, but I hope I'll be faster than that

//TODO: improve Error Handling

func getHeaders(r http.Request) bool {
	// Loop over header names
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
	return true
}

func handleConnection(c net.Conn) {
	defer c.Close()
	bufReader := bufio.NewReader(c)

	for {
		// Read tokens delimited by newline
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", bytes)
	}
}

func listen(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("an error occured:", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("an error occured:", err)
		}
		go handleConnection(conn)
	}

}

func main() {
	listen(":8888")
}

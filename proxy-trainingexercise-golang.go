// proxy-trainingexercise-golang
package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
https://pkg.go.dev/net will probably do the most work for me,
user makes requests to url, the proxy forwards them to the url, including headers and
status codes.*/

//https://www.upguard.com/blog/proxy-server

//my time limit is 2 days, but I hope I'll be faster than that

func main() {
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status, err)
}

// proxy-trainingexercise-golang
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

/*
https://pkg.go.dev/net will probably do the most work for me,
user makes requests to url, the proxy forwards them to the url, including headers and
status codes.

*/

//https://www.upguard.com/blog/proxy-server

//my time limit is 2 days, but I hope I'll be faster than that

//TODO: improve Error Handling
/*
	fmt.Println("FEHLER!: ", err)
	return*/
//now doing a breaking change, so to github it goes

func handleConnection(c net.Conn) { /*
		fmt.Println("FEHLER!: ", err)
		return*/
	defer fmt.Println("closing")
	defer c.Close()
	var headers []string
	bufReader := bufio.NewReader(c)
	for {
		// Read tokens delimited by newline
		bytes, err := bufReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				fmt.Println("EOF reached!")
				break
			}
		}
		headers = append(headers, string(bytes))
	}
	method := strings.Split(headers[0], " ")[0]
	url := strings.Split(headers[0], " ")[1]
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("error while forming request")
	}
	for i := 1; i < len(headers)-1; i++ {
		line := strings.Split(headers[i], ": ")
		req.Header.Set(line[0], line[1])
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

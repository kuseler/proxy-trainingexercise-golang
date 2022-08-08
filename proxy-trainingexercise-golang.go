// proxy-trainingexercise-golang
package main

import (
	"bufio"
	"fmt"
	"net"
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

func handleConnection(c net.Conn) {
	defer c.Close()
	bufReader := bufio.NewReader(c)

	for {
		// Read tokens delimited by newline
		bytes, err := bufReader.ReadBytes('\n')
		if err != nil {
			//fmt.Println("FEHLER!: ", err)
			return
		}

		//fmt.Printf("%s", bytes)
		headers := strings.Split(string(bytes), "\n")
		for i := 0; i < len(headers); i++ {
			fmt.Print(headers[i])
		}
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

// proxy-trainingexercise-golang
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

/*
https://pkg.go.dev/net will probably do the most Fetchwork for me.
user makes requests to url, the proxy forwards them to the url, including headers and
status codes.
*/

//https://www.upguard.com/blog/proxy-server
//I have 1 day left, but I hope I'll be faster than that
//TODO: improve Error Handling
//TODO improve readability

func handleConnection(c net.Conn) { //now doing a breaking change, so to github it goes
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
	url := strings.Split(headers[0], " ")[1][1:]
	req, err := http.NewRequest(method, url, nil)
	fmt.Println(headers)
	if err != nil {
		fmt.Println("error while forming request")
	}
	for i := 1; i < len(headers)-1; i++ {
		if !strings.Contains(headers[i], "Sec-Fetch") {
			line := strings.Split(headers[i], ": ")
			req.Header.Set(strings.TrimSpace(line[0]), strings.TrimSpace(line[1]))
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error while making Request", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error trying to read response body", err)
	}
	fmt.Printf("client: response body: %s\n", resBody)
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

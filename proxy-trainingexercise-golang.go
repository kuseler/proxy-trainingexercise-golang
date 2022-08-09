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
https://pkg.go.dev/net will probably do the most work for me.
user makes requests to url, the proxy forwards them to the url, including headers and
status codes.
*/

//https://www.upguard.com/blog/proxy-server
//I have 1 day left
//TODO: improve Error Handling
//TODO improve readability
//weird bug where the page has to be reloaded in order for the program to do anything

func getHeaders(connection net.Conn) []string {
	var headers []string
	bufReader := bufio.NewReader(connection)
	for {
		// Read tokens delimited by newline
		bytes, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF reached!")
				break
			}
			fmt.Println(err)
		}
		headers = append(headers, string(bytes))
	}
	return headers
}

func genRequest(headers []string) *http.Request {
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
	return req
}

func handleConnection(c net.Conn) {
	defer fmt.Println("closing")
	defer c.Close()
	headers := getHeaders(c)
	// here is the actual request
	request := genRequest(headers)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("error while making Request", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error trying to read response body", err)
	}
	fmt.Println(string(resBody))
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

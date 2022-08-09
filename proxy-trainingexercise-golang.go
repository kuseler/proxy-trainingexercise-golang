// proxy-trainingexercise-golang
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
https://pkg.go.dev/net will probably do the most work for me.
user makes requests to url, the proxy forwards them to the url, including headers and
status codes.
*/

//https://www.upguard.com/blog/proxy-server
//I have 1 day left
//TODO improve readability

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		url := req.URL.String()
		fmt.Printf("url: %s", url)
		fwReq, _ := http.NewRequest(req.Method, "http://example.org", req.Body)
		fwReq.Header = req.Header.Clone()
		resp, err := http.DefaultClient.Do(fwReq)
		if err != nil {
			fmt.Println(err)
		}
		respBody, _ := ioutil.ReadAll(resp.Body)
		for k, v := range req.Header.Clone() {
			w.Header().Set(k, v[0])
		}
		w.Write(respBody)
		fmt.Println(resp)
		fmt.Println(req.Header)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}

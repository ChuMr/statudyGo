package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:9000/client")
	if err != nil {
		fmt.Printf("http get failed,err: %v", err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(b))
}

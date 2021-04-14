package main

import (
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./test.html")
	w.Write(b)
}

func main() {
	http.HandleFunc("/hello/", f1)
	http.ListenAndServe("127.0.0.1:9001", nil)
}

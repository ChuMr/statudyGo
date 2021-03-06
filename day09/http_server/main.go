package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/test", f1)
	http.HandleFunc("/client", f2)
	http.ListenAndServe("127.0.0.1:9111", nil)
}

func f1(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))

	w.Write([]byte("err ok!"))
}

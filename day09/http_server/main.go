package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/test", f1)
	http.ListenAndServe("127.0.0.1:9000", nil)
}

func f1(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadFile("test.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)

}

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello World!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":5000", nil)
}

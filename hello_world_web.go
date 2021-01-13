package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "yes")
	})

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("there is a new request")
	io.WriteString(w, "Hello World")
}

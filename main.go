package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, otherMethod())
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func otherMethod() string {
	var a int = 15
	var b int = 5

	var c int = a + b

	var result string = "Hello! A+B is: " + strconv.Itoa(c)

	return result
}

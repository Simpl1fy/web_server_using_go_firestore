package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Welcome to the basic Http server")
	if err != nil {
		fmt.Println("error", err)
		return
	}
}

func main() {
	fmt.Println("Hello World this is a web server in go")

	http.HandleFunc("/", handler)
	fmt.Println("Listening to port 5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("Error while listening and serving =", err)
		return
	}
}

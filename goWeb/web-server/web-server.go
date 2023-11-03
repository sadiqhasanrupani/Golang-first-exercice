package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	_ = http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprint(w, "Hello world")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			log.Println("Warning...")
		}
		fmt.Println(n)
	})

	http.ListenAndServe(":8080", nil)
}

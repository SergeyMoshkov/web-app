package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Home it is function for home page
func Home(w http.ResponseWriter, _ *http.Request) {
	sum := AddValues(2, 4)
	n, err := fmt.Fprintf(w, "This is home page and sum 2 + 4 is %d", sum)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
}

// About it is function for about page
func About(w http.ResponseWriter, _ *http.Request) {
	sum := AddValues(2, 2)
	n, err := fmt.Fprintf(w, "This is about page and sum 2 + 2 is %d", sum)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
}

// AddValues this function added two digit
func AddValues(x, y int) int {
	return x + y
}

// main the main function our application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Application run on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

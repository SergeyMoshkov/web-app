package main

import (
	"fmt"
	"log"
	"net/http"
)

// Home it is function for home page
func Home(w http.ResponseWriter, _ *http.Request) {
	sum := AddValues(2, 4)
	n, err := fmt.Fprintf(w, "This is home page and sum 2 + 4 is %d", sum)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
}

// About it is function for home page
func About(w http.ResponseWriter, _ *http.Request) {
	sum := AddValues(2, 2)
	n, err := fmt.Fprintf(w, "This is about page and sum 2 + 2 is %d", sum)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)
}

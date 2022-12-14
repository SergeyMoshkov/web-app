package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home it is function for home page
func Home(w http.ResponseWriter, _ *http.Request) {
	sum := AddValues(2, 4)
	_, _ = fmt.Fprintf(w, "This is home page and sum 2 + 4 is %d", sum)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(n)
}

// About it is function for about page
func About(w http.ResponseWriter, _ *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is about page and sum 2 + 2 is %d", sum)
	//if err != nil {
	//	log.Println(err)
	//}
	//fmt.Println(n)
}

// AddValues this function added two digit
func AddValues(x, y int) int {
	return x + y
}

// Divide this function divide two digit
func Divide(w http.ResponseWriter, _ *http.Request) {
	f, err := divideValues(100.0, -1.0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by zero")
		return
	}
	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, -1.0, f))
}

// Divide this function divide two digit and check divide by zero
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

// main the main function our application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Application run on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

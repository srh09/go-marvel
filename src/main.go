package main

import (
	"fmt"
	"log"
	"net/http"
)

// func testHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

// func main() {
// 	fmt.Println("Starting the application-------")
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", testHandler)

// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	fmt.Println("and other items-----")
}

func main() {
	fmt.Println("test here----")
	http.HandleFunc("/monkeys", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

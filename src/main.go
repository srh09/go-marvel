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

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("and other ite222s-----")
}

func main() {
	mux := http.NewServeMux()
	fmt.Println("test again------")
	http.HandleFunc("/monkeys", handler)
	// mux.HandleFunc("/", handler2)
	mux.Handle("/", http.FileServer(http.Dir("./src/static")))
	log.Fatal(http.ListenAndServe(":8000", mux))
}

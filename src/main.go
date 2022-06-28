package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/srh09/go-marvel/src/api/handlers"
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

// func handler2(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("and other ite222s-----")
// }

// func handleMonkey(w http.ResponseWriter, r *http.Request) {
// 	id := strings.TrimPrefix(r.URL.Path, "/monkey/")
// 	fmt.Printf("The monkey was identified with id: %s\n", id)
// 	fmt.Println(r.URL.Path)
// }

func main() {
	mux := http.NewServeMux()
	fmt.Println("test again------")
	mux.HandleFunc("/monkey/", handlers.HandleMonkey)
	mux.HandleFunc("/monkeys", handler)
	// mux.HandleFunc("/", handler2)
	mux.Handle("/", http.FileServer(http.Dir("./src/static")))
	log.Fatal(http.ListenAndServe(":8000", mux))
}

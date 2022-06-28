package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleMonkey(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/monkey/")
	fmt.Printf("The monkey was identified with id: %s\n", id)
	fmt.Println(r.URL.Path)
}

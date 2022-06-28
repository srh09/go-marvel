package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/srh09/go-marvel/src/util"
)

func HandleMonkey(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/monkey/")
	fmt.Printf("The monkey was identified with id: %s\n", id)
	fmt.Println(r.URL.Path)

	s := "example.com/character/123/comic/456?one=apple&two=bananas"
	path, pathParams, queryParams := util.ParseURL(s)
	fmt.Println(path)
	fmt.Println(pathParams)
	fmt.Println(queryParams)
}

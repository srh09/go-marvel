package util

import (
	"net/url"
	"strconv"
	"strings"
)

func Split(r rune) bool {
	return r == '/' || r == '?'
}

func ParseURL(s string) (string, []int, url.Values) {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fragments := strings.FieldsFunc(s, Split)
	pathParams := []int{}
	for _, fragment := range fragments {
		i, err := strconv.Atoi(fragment)
		if err != nil {
			continue
		}
		pathParams = append(pathParams, i)
	}
	path := u.Path
	queryParams, _ := url.ParseQuery(u.RawQuery)

	return path, pathParams, queryParams
}

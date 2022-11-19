package main

import "net/http"

type Variable struct {
	url      string
	method   string
	body     []byte
	header   http.Header
	request  int
	duration int
}

// func (variable *Variable) convertToInt(value string) (int, error) {
// 	r, err := strconv.Atoi(value)
// 	if err == nil {
// 		return r,
// 	}
// }

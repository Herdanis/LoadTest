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

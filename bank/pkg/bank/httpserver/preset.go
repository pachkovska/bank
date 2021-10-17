package httpserver

import "net/http"

const (
	// URIPrefix ...
	URIPrefix = "/api/v1"

	// URIPrefixbank ...
	URIPrefixBank = URIPrefix + "/bank"

	// URIPathClientGetBalance ...
	URIPathClientGetBalance = URIPrefixBank + "/balance"

	// HTTPMethodCreatebank ...
	HTTPMethodGetBalance = http.MethodPost
)

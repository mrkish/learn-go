package http_server

import "net/http"

func ListenAndServe(addr string, handler Handler) error {
	return nil
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func PlayerServer() {}

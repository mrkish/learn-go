package http_server

func ListenAndServe(addr string, handler Handler) error {
	return nil
}

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func PlayerServer() {}

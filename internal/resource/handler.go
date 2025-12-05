package resource

import "net/http"

type ResourceHandler struct {
	Service ResourceService
}

func GetResource(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("Hello, world!\n"))
}

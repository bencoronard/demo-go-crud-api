package config

import (
	"net/http"

	xhttp "github.com/bencoronard/demo-go-common-libs/http"
)

type container struct {
	srv *http.Server
}

func NewContainer(h *http.Server) xhttp.Container {
	return &container{srv: h}
}

func (c *container) ServeHTTP() error {
	return c.srv.ListenAndServe()
}

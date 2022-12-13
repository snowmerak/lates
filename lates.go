package lates

import (
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
)

type Lates struct {
	router *httprouter.Router
	server *http.Server
}

func New() *Lates {
	return &Lates{
		router: httprouter.New(),
	}
}

func (l *Lates) SetHttpServer(server *http.Server) {
	l.server = server
}

func (l *Lates) ListenAndServe(addr string) error {
	if l.server != nil {
		if addr != "" {
			l.server.Addr = addr
		}
		return l.server.ListenAndServe()
	}
	return http.ListenAndServe(addr, l.router)
}

func (l *Lates) ListenAndServe2(addr string) error {
	if l.server == nil {
		return errors.New("server must be set")
	}

	if addr != "" {
		l.server.Addr = addr
	}

	if err := http2.ConfigureServer(l.server, nil); err != nil {
		return err
	}

	return l.server.ListenAndServe()
}

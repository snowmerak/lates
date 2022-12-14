package lates

import (
	"errors"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
	"github.com/snowmerak/lates/made"
)

func (l *Lates) Get(path string, handler func(c *Context) templ.Component) {
	l.AddRouter("GET", path, handler)
}

func (l *Lates) Post(path string, handler func(c *Context) templ.Component) {
	l.AddRouter("POST", path, handler)
}

func (l *Lates) Put(path string, handler func(c *Context) templ.Component) {
	l.AddRouter("PUT", path, handler)
}

func (l *Lates) Delete(path string, handler func(c *Context) templ.Component) {
	l.AddRouter("DELETE", path, handler)
}

func (l *Lates) Patch(path string, handler func(c *Context) templ.Component) {
	l.AddRouter("PATCH", path, handler)
}

func (l *Lates) AddRouter(method string, path string, handler func(c *Context) templ.Component) error {
	method = strings.ToUpper(method)
	switch method {
	case "GET":
	case "POST":
	case "PUT":
	case "DELETE":
	case "PATCH":
	default:
		return errors.New("invalid method")
	}

	l.router.Handle(method, path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		c := &Context{
			r: r,
			p: p,
			w: w,
		}
		comp := handler(c)
		comp = made.Layout(comp, l.cssPathes...)
		templ.Handler(comp).ServeHTTP(w, r)
	})

	return nil
}

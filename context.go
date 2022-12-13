package lates

import (
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	r *http.Request
	p httprouter.Params
	w http.ResponseWriter
}

func (c *Context) GetFormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.r.FormFile(key)
}

func (c *Context) GetMultipartFile(name string, maxMemoryBytes int64) ([]*multipart.FileHeader, error) {
	if c.r.MultipartForm == nil {
		err := c.r.ParseMultipartForm(maxMemoryBytes)
		if err != nil {
			return nil, err
		}
	}
	return c.r.MultipartForm.File[name], nil
}

func (c *Context) SaveFile(name, path string) error {
	file, _, err := c.r.FormFile(name)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	return err
}

func (c *Context) SaveMultipartFile(name string, maxMemoryBytes int64, path string) error {
	headers, err := c.GetMultipartFile(name, maxMemoryBytes)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, header := range headers {
		file, err := header.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := io.Copy(f, file); err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) GetFormData(key string) string {
	return c.r.FormValue(key)
}

func (c *Context) GetURLQuery(key string) string {
	return c.r.URL.Query().Get(key)
}

func (c *Context) GetPathVariable(key string) string {
	return c.p.ByName(key)
}

func (c *Context) GetBody() ([]byte, error) {
	return io.ReadAll(c.r.Body)
}

func (c *Context) GetBodyReader() io.ReadCloser {
	return c.r.Body
}

func (c *Context) GetCookie(key string) string {
	return c.r.Header.Get("Cookie")
}

func (c *Context) GetRemoteAddress() string {
	return c.r.RemoteAddr
}

func (c *Context) GetRemoteIpPort() (string, string, error) {
	return net.SplitHostPort(c.r.RemoteAddr)
}

func (c *Context) SetStatusCode(statusCode int) {
	c.w.WriteHeader(statusCode)
}

func (c *Context) Redirect(url string, statusCode int) {
	http.Redirect(c.w, c.r, url, statusCode)
}

# Lates

[riːt]

`lates` is a simple tool for `templ`.

## install

### install github.com/a-h/templ

first, visit [`templ`](https://github.com/a-h/templ), and install `templ`.

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

and, install visual studio extension `templ` from the marketplace.

`https://marketplace.visualstudio.com/items?itemName=a-h.templ`

### go get github.com/snowmerak/lates

```bash
go get github.com/snowmerak/lates
```

## usage

### how to use templ

read the `readme.md` of `templ`.

[link](https://github.com/a-h/templ/blob/main/README.md)

### lates example

```go
// templates/index/helloworld.templ
package index

templ HelloWorld() {
    <h1>Hello, world!</h1>
}
```

```go
// templates/index/param.templ
package index

templ Param(param string) {
    <h1>{param}</h1>
}
```

run `templ generate`.

```go
// main.go
package main

import (
	"example/templates/index"

	"github.com/a-h/templ"
	"github.com/snowmerak/lates"
)

func main() {
	late := lates.New()

	late.Get("/", func(c *lates.Context) templ.Component {
		return index.HelloWorld()
	})

	late.Get("/:param", func(c *lates.Context) templ.Component {
		return index.Param(c.GetPathVariable("param"))
	})

	if err := late.ListenAndServe(":8080"); err != nil {
		panic(err)
	}
}

```

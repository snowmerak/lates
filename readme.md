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
package main

import "github.com/snowmerak/lates"

func main() {
    late := lates.New()

    late.Get("/", helloworld.Get)

    if err := late.ListenAndServe(":8080"); err != nil {
        panic(err)
    }
}
```

package main

import (
  "code.google.com/p/goweb/goweb"
  "fmt"
)

func main() {

  goweb.MapFunc("/", func(c *goweb.Context) {
    fmt.Fprintf(c.ResponseWriter, '', )
  })

  goweb.ListenAndServe(":80")
}

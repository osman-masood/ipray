package main

import (
  "code.google.com/p/goweb/goweb"
  "fmt"
  "io/ioutil"
)

func main() {
  pageRead, error := ioutil.ReadFile("ipray.html")
  if error != nil {
    panic(error.String())
  }

  goweb.MapFunc("/", func(c *goweb.Context) {
    fmt.Fprintf(c.ResponseWriter, string(pageRead))
  })

  goweb.ListenAndServe(":80")
}

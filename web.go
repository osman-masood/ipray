package main

import (
  "code.google.com/p/goweb/goweb"
  "fmt"
  "io/ioutil"
)

func main() {
  pageRead, err := ioutil.ReadFile("ipray.html")
  if err != nil {
    panic(err.String())
  }
  pageReadString := string(pageRead)

  goweb.MapFunc("/", func(c *goweb.Context) {
    fmt.Fprintf(c.ResponseWriter, pageReadString)
  })

  goweb.ListenAndServe(":80")
}

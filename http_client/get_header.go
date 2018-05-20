package main

import (
  "log"
  "net/http"
)

func main() {
  resp, err := http.Get("http://localhost:18888")
  if err != nil {
    panic(err)
  }
  log.Println("status:", resp.Status)
  log.Println("Headers", resp.Header)
}

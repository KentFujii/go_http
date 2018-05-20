package main

import (
  "log"
  "net/http"
  "os"
  // "strings"
)

func main() {
  file, err := os.Open("readme.md")
  if err != nil {
    panic(err)
  }
  resp, err := http.Post("http://localhost:18888", "text/plain", file)
  // あるいは
  // reader := strings.NewReader("テキスト")
  // resp, err := http.Post("http://localhost:18888", "text/plain", reader)
  if err != nil {
    panic(err)
  }
  log.Println("Status:", resp.Status)
}

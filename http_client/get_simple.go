package main

import (
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  resp, _ := http.Get("http://localhost:18888")
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  log.Println(string(body))
  // ステータスコード 文字列で"200 OK"
  log.Println("Status:", resp.Status)
  // ステータスコード 数値で200
  log.Println("StatusCode:", resp.StatusCode)
  // ヘッダー
  log.Println("Headers:", resp.Header)
  // ヘッダーの一部を指定して取得
  log.Println("Headers:", resp.Header.Get("Content-Length"))
}

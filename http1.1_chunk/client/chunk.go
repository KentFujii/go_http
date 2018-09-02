package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18886/chunked")
	if err != nil {
	}
	defer resp.Body.Close()
	// サーバーから随時送信されるコンテンツの区切りを確認することでチャンクを扱かえる
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handlerChunkedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		// 一回書き込むごとにFlush()を呼び出すことで、クライアントはループ毎に結果を受け取れる
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/chunked", handlerChunkedResponse)
	log.Printf("Start http listening :18886")
	httpServer.Addr = ":18886"
	log.Println(httpServer.ListenAndServe())
}

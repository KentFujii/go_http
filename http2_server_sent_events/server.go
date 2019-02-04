package main

import (
	"fmt"
	"math/big"
	"net/http"
	"time"
)

func handlerSSE(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}
	closeNotify := w.(http.CloseNotifier).CloseNotify()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var num int64 = 1
	for id := 1; id <= 100; id++ {
		// 通信が切れても終了
		select {
		case <-closeNotify:
			fmt.Println("Connection closed from client")
			return
		default:
			// do nothing
		}
		for {
			num++
			// 確率論的に素数を求める
			if big.NewInt(num).ProbablyPrime(20) {
				fmt.Println(num)
				fmt.Fprintf(w, "data: {\"id\": %d, \"number\": %d}\n\n", id, num)
				flusher.Flush()
				time.Sleep(time.Second)
				break
			}
		}
		time.Sleep(time.Second)
	}
	// 100個超えたら送信終了
	fmt.Println("Connection close from server")
}

func main() {
	http.HandleFunc("/", handlerSSE)
	fmt.Println("start htp listening :18884")
	err := http.ListenAndServe(":18884", nil)
	fmt.Println(err)
}

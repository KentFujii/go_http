package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculator int

// RPCで外部から呼ばれるメソッド
func (c *Calculator) Multiply(args Args, result :int) {
	log.Printf("Multiply called: %d, %d\n", args.A, args.B)
	*result = args.A * args.B
	return nil
}

// 外部から呼ばれる時の引数
type Args struct {
	A, B int
}

func main() {
}

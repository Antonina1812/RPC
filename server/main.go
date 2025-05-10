package main

import (
	"RPC/shared"
	"log"
	"net"
	"net/rpc"
)

func main() {
	mathService := new(shared.MathService)

	rpc.Register(mathService)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Println("RPC server started on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}

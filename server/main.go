package main

import (
	"RPC/shared"
	"net"
	"net/rpc"
)

func main() {
	mathService := new(shared.MathService)

	// Регистрируем сервис
	err := rpc.RegisterName("MathService", mathService)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	println("Server is running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

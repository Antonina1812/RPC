package main

import (
	"RPC/shared"
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer client.Close()

	args := &shared.MultiplyArgs{X: 6, Y: 7}
	var reply shared.MultiplyReply

	err = client.Call("MathService.Multiply", args, &reply)
	if err != nil {
		fmt.Println("RPC call error:", err)
		return
	}

	fmt.Printf("RPC Result: %d * %d = %d\n", args.X, args.Y, reply.Result)
}

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

	var x, y int
	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)

	args := &shared.MultiplyArgs{X: x, Y: y}
	reply := &shared.MultiplyReply{}

	err = client.Call("MathService.Multiply", args, reply)
	if err != nil {
		fmt.Println("RPC call error:", err)
		return
	}

	fmt.Printf("Result: %d * %d = %d\n", args.X, args.Y, reply.Result)
}

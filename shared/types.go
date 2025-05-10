package shared

import "log"

type MultiplyArgs struct {
	X, Y int
}

type MultiplyReply struct {
	Result int
}

type MathService struct{}

func (m *MathService) Multiply(args *MultiplyArgs, reply *MultiplyReply) error {
	reply.Result = args.X * args.Y
	log.Printf("Multiply: %d * %d = %d", args.X, args.Y, reply.Result)
	return nil
}

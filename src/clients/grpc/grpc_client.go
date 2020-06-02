package grpc

import (
	"google.golang.org/grpc"
	"log"
)

func NewGRPCClient(url string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
		return nil, err
	}

	return conn, nil
}

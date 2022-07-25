package main

import (
	"context"
	"fmt"
	"github.com/emiksk/earthly-example/grpc/pb"
	"google.golang.org/grpc"
	"net"
	"os"
)

const portNumber = 8080

func main() {
	server := grpc.NewServer()

	pb.RegisterEchoServiceServer(server, &Server{})

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		fmt.Println("failed to listen:", err)
		os.Exit(1)
	}

	server.Serve(l)
}

type Server struct {
	pb.UnimplementedEchoServiceServer
}

var _ pb.EchoServiceServer = &Server{}

func (s *Server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: "Reply: " + req.Message,
	}, nil
}

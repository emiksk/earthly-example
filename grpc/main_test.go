package main

import (
	"context"
	"fmt"
	"github.com/emiksk/earthly-example/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

// gRPCサーバを立ち上げて，クライアントで Echo メソッドを実行するテスト
func Test(t *testing.T) {
	go main()

	conn, _ := grpc.Dial(fmt.Sprintf("localhost:%d", portNumber), grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewEchoServiceClient(conn)

	message := "Hello"
	reply, err := client.Echo(context.Background(), &pb.EchoRequest{Message: message})
	if err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}

	expected := "Reply: " + message
	if reply.Message != expected {
		t.Fatalf("Unexpected reply: expected reply is %v, but actual reply is %v", expected, reply.Message)
	}
}

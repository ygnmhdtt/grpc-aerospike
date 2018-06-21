package main

import (
	"context"
	"log"
	"net"
	// pb "github.com/ygnmhdtt/grpc-aerospike/grpc-aerospike-kvs/grpc-aerospike-kvs.pb.go"
)

const (
	port = ":50051"
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// GetMemo returns memo by name
func GetMemo(ctx context.Context, in *GetMemoRequest, opts ...grpc.CallOption) (*MemoResponse, error) {
}

// PutMemo creates new memo
func PutMemo(ctx context.Context, in *PutMemoRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
}

// DeleteMemo deletes memo by name
func DeleteMemo(ctx context.Context, in *DeleteMemoRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterKvsServer(s, NewKvsServer())
	s.Serve(lis)
}

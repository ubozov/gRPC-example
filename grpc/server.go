package server

import (
	"context"
	"net"

	"github.com/ubozov/grpc-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

// Start starts the grpc server
func Start() error {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	return srv.Serve(listener)
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b

	return &proto.Response{Result: result}, nil
}

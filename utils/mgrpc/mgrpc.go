package mgrpc

import (
	"context"
	"net"

	"casorder/utils/logging"

	grpc "google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var LOG = logging.GetLogger()

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	LOG.Info("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type GrpcServer struct {
	host string
	port string
}

func New(host string, port string) *GrpcServer {
	gs := GrpcServer{
		host: host,
		port: port,
	}

	return &gs
}

func (gs *GrpcServer) Start() error {
	grpcAddr := net.JoinHostPort(gs.host, gs.port)
	LOG.Info("Starting Grpc server: ", grpcAddr)

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})
	grpcServer.Serve(listener)
	if err := grpcServer.Serve(listener); err != nil {
		LOG.Error("failed to serve: ", err)
		return err
	}
	return nil
}

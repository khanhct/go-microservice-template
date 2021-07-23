package mgrpc

import (
	"context"
	"net"
	"time"

	"casorder/utils/logging"

	grpc "google.golang.org/grpc"
	pb "casorder/taskmanager/grpc/build/helloworld"
)

var LOG = logging.GetLogger()

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	LOG.Info("Received: ", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type GrpcServer struct {
	host string
	port string
	grpcServer *grpc.Server
}

func New(host string, port string) GrpcServer {
	gs := GrpcServer{
		host: host,
		port: port,
	}
	gs.grpcServer = grpc.NewServer()

	return gs
}

func (gs *GrpcServer) Start() error {
	grpcAddr := net.JoinHostPort(gs.host, gs.port)
	LOG.Info("Starting Grpc server: ", grpcAddr)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return err
	}

	// Register servicers
	pb.RegisterGreeterServer(gs.grpcServer, &server{})

	gs.grpcServer.Serve(listener)
	if err := gs.grpcServer.Serve(listener); err != nil {
		LOG.Error("failed to serve: ", err)
		return err
	}
	return nil
}


func TestClient(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		LOG.Error("did not connect: ", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "hello KhanhCT"})
	if err != nil {
		LOG.Error("could not greet: ", err)
	}
	LOG.Info("Greeting: ", r.GetMessage())
}
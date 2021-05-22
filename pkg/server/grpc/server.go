package grpc

import (
	"context"
	"net"

	"github.com/ringsaturn/go-web-template/api"
	"github.com/ringsaturn/go-web-template/pkg/config"
	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedHelloServer
	grpcServer *grpc.Server
}

func NewServer(conf *config.Config) (*Server, error) {
	server := &Server{
		grpcServer: grpc.NewServer(),
	}
	api.RegisterHelloServer(server.grpcServer, &Server{})
	return server, nil
}

func (srv *Server) Ping(ctx context.Context, req *api.HelloReq) (*api.HelloResp, error) {
	return &api.HelloResp{}, nil
}

func (srv *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", "localhost:1111")
	if err != nil {
		return err
	}
	return srv.grpcServer.Serve(lis)
}

func (srv *Server) Shutdown(ctx context.Context) error {
	srv.grpcServer.GracefulStop()
	return nil
}

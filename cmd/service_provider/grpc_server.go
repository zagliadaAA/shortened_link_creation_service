package service_provider

import (
	pb "shortened_link_creation_service/protos/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (sp *ServiceProvider) GetGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterShortenerURLServer(grpcServer, sp.GetLinkController())

	reflection.Register(grpcServer)

	return grpcServer
}

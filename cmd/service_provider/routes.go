package service_provider

import (
	pb "shortened_link_creation_service/protos/gen/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	// роутер ручек
	mux := http.NewServeMux()
	// link
	mux.HandleFunc("POST /createShortURL", sp.GetLinkController().CreateShortURL)
	mux.HandleFunc("GET /returnURL", sp.GetLinkController().ReturnURL)

	return mux
}
*/

func (sp *ServiceProvider) GetGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterShortenerURLServer(grpcServer, sp.GetLinkController())

	reflection.Register(grpcServer)

	return grpcServer
}

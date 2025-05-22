package grpc

import (
	"context"

	pb "shortened_link_creation_service/protos/gen/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Controller) ReturnURL(ctx context.Context, req *pb.ReturnURLRequest) (*pb.ReturnURLResponse, error) {
	link, err := s.linkUseCase.GetLinkByShortURL(req.ShortUrl)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "не удалось найти оригинальный URL: %v", err)
	}

	return &pb.ReturnURLResponse{
		Url: link.URL,
	}, nil
}

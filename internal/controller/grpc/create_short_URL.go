package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "shortened_link_creation_service/protos/gen/proto"
)

func (s *Controller) CreateShortURL(ctx context.Context, req *pb.CreateShortURLRequest) (*pb.CreateShortURLResponse, error) {
	// Проверяем, существует ли уже ссылка
	link, err := s.linkUseCase.GetLinkByURL(req.Url)
	if err != nil {
		// Если не существует, создаем новую
		link, err = s.linkUseCase.CreateShortURL(req.Url)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "не удалось создать короткий URL: %v", err)
		}
	}

	return &pb.CreateShortURLResponse{
		ShortUrl: link.ShortURL,
	}, nil
}

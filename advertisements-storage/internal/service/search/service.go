package searchservice

import (
	"context"
	"errors"
	"strings"

	types "advertisement-storage/pkg"
	pb "advertisement-storage/pkg/pb"
)

type searchService struct {
	repo types.SearchRepository
}

func NewSearchService(repo types.SearchRepository) *searchService {
	return &searchService{
		repo: repo,
	}
}

func (s *searchService) SearchAdvertisementByTitle(ctx context.Context, req *pb.SearchAdvertisementByTitleRequest) (*pb.SearchAdvertisementByTitleResponse, error) {
	if len(strings.TrimSpace(req.Title)) == 0 {
		return &pb.SearchAdvertisementByTitleResponse{
			Advertisement: []*pb.Advertisement{},
		}, nil
	}

	if len(req.Title) > 200 {
		return nil, errors.New("search query is too long")
	}

	title := strings.TrimSpace(req.Title)

	return s.repo.Search(ctx, &pb.SearchAdvertisementByTitleRequest{
		Title: title,
	})
}

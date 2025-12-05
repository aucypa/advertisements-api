package advertisementsservice

import (
	types "advertisement-storage/pkg"
	pb "advertisement-storage/pkg/pb"
	"context"
	"errors"
	"strings"
)

type advertisementsService struct {
	adRepo   types.AdvertisementsRepository
	catRepo  types.CategoryRepository
	userRepo types.UserRepository
}

func NewAdvertisementsService(adRepo types.AdvertisementsRepository, catRepo types.CategoryRepository, userRepo types.UserRepository) *advertisementsService {
	return &advertisementsService{
		adRepo:   adRepo,
		catRepo:  catRepo,
		userRepo: userRepo,
	}
}

func (s *advertisementsService) CreateAdvertisement(ctx context.Context, req *pb.CreateAdvertisementRequest) (*pb.CreateAdvertisementResponse, error) {
	if len(req.Title) == 0 {
		return nil, errors.New("title is required")
	}

	if len(req.Description) == 0 {
		return nil, errors.New("description is required")
	}

	if req.Price == 0 {
		return nil, errors.New("price is required")
	}

	if len(req.Currency) == 0 {
		return nil, errors.New("currency is required")
	}

	if req.CategoryId == 0 {
		return nil, errors.New("categoryId is required")
	}

	if req.UserId == 0 {
		return nil, errors.New("userId is required")
	}

	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)
	req.Currency = strings.TrimSpace(strings.ToUpper(req.Currency))

	_, err := s.catRepo.GetByID(ctx, &pb.GetCategoryByIdRequest{Id: req.CategoryId})
	if err != nil {
		return nil, err
	}

	_, err = s.userRepo.GetByID(ctx, &pb.GetUserByIDRequest{Id: req.UserId})
	if err != nil {
		return nil, err
	}

	return s.adRepo.Create(ctx, req)
}

func (s *advertisementsService) GetAdvertisementById(ctx context.Context, req *pb.GetAdvertisementByIdRequest) (*pb.GetAdvertisementByIdResponse, error) {
	if req.Id <= 0 {
		return nil, errors.New("invalid advertisement ID")
	}

	return s.adRepo.GetByID(ctx, req)
}

func (s *advertisementsService) GetAllAdvertisements(ctx context.Context, req *pb.GetAllAdvertisementsRequest) (*pb.GetAllAdvertisementsResponse, error) {
	return s.adRepo.GetAll(ctx, req)
}

func (s *advertisementsService) UpdateAdvertisementById(ctx context.Context, req *pb.UpdateAdvertisementByIdRequest) (*pb.UpdateAdvertisementByIdResponse, error) {
	if req.Id <= 0 {
		return nil, errors.New("invalid advertisement ID")
	}

	advertisement, err := s.adRepo.GetByID(ctx, &pb.GetAdvertisementByIdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	if req.CategoryId != 0 && req.CategoryId != advertisement.Advertisement.CategoryId {
		_, err := s.catRepo.GetByID(ctx, &pb.GetCategoryByIdRequest{Id: req.CategoryId})
		if err != nil {
			return nil, err
		}
	}

	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)
	req.Currency = strings.TrimSpace(strings.ToUpper(req.Currency))

	return s.adRepo.Update(ctx, req)
}

func (s *advertisementsService) DeleteAdvertisementByID(ctx context.Context, req *pb.DeleteAdvertisementByIDRequest) (*pb.DeleteAdvertisementByIDResponse, error) {
	if req.Id <= 0 {
		return nil, errors.New("invalid advertisement ID")
	}

	_, err := s.adRepo.GetByID(ctx, &pb.GetAdvertisementByIdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	_, err = s.adRepo.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAdvertisementByIDResponse{}, nil
}

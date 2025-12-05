package categoryservice

import (
	"context"
	"errors"
	"strings"

	types "advertisement-storage/pkg"
	pb "advertisement-storage/pkg/pb"
)

type categoryService struct {
	repo types.CategoryRepository
}

func NewCategoryService(repo types.CategoryRepository) *categoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	if len(strings.TrimSpace(req.Name)) == 0 {
		return nil, errors.New("category name is required")
	}

	if len(strings.TrimSpace(req.Alias)) == 0 {
		return nil, errors.New("category alias is required")
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Alias = strings.ToLower(strings.TrimSpace(req.Alias))

	existing, err := s.repo.GetByAlias(ctx, &pb.GetCategoryByAliasRequest{Alias: req.Alias})
	if err == nil && existing != nil {
		return nil, err
	}

	return s.repo.Create(ctx, req)
}

func (s *categoryService) GetCategoryById(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error) {
	if req.Id <= 0 {
		return nil, errors.New("invalid category ID")
	}
	return s.repo.GetByID(ctx, req)
}

func (s *categoryService) GetCategoryByAlias(ctx context.Context, req *pb.GetCategoryByAliasRequest) (*pb.GetCategoryByAliasResponse, error) {
	if strings.TrimSpace(req.Alias) == "" {
		return nil, errors.New("alias is required")
	}

	req.Alias = strings.ToLower(strings.TrimSpace(req.Alias))

	return s.repo.GetByAlias(ctx, req)
}

func (s *categoryService) GetAllCategory(ctx context.Context, req *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error) {
	return s.repo.GetAll(ctx, req)
}

func (s *categoryService) UpdateCategoryById(ctx context.Context, req *pb.UpdateCategoryByIdRequest) (*pb.UpdateCategoryByIdResponse, error) {
	if req.Id <= 0 {
		return nil, errors.New("invalid category ID")
	}

	if len(req.Alias) != 0 {
		req.Alias = strings.ToLower(strings.TrimSpace(req.Alias))

		existing, err := s.repo.GetByAlias(ctx, &pb.GetCategoryByAliasRequest{Alias: req.Alias})
		if err == nil && existing != nil && existing.Category.Id != req.Id {
			return nil, err
		}
	}

	if len(req.Name) != 0 {
		req.Name = strings.TrimSpace(req.Name)
	}

	return s.repo.Update(ctx, req)
}

func (s *categoryService) DeleteCategoryById(ctx context.Context, req *pb.DeleteCategoryByIDRequest) (*pb.DeleteCategoryByIDResponse, error) {
	if req.Id <= 0 {
		return nil, errors.New("invalid category ID")
	}

	return s.repo.Delete(ctx, req)
}

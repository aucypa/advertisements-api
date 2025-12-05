package types

import (
	"context"
	"database/sql"

	pb "advertisement-storage/pkg/pb"
)

type Database interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type CategoryRepository interface {
	Create(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error)
	GetByID(ctx context.Context, req *pb.GetCategoryByIdRequest) (*pb.GetCategoryByIdResponse, error)
	GetAll(ctx context.Context, req *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error)
	GetByAlias(ctx context.Context, req *pb.GetCategoryByAliasRequest) (*pb.GetCategoryByAliasResponse, error)
	Update(ctx context.Context, req *pb.UpdateCategoryByIdRequest) (*pb.UpdateCategoryByIdResponse, error)
	Delete(ctx context.Context, req *pb.DeleteCategoryByIDRequest) (*pb.DeleteCategoryByIDResponse, error)
}

type SearchRepository interface {
	Search(ctx context.Context, req *pb.SearchAdvertisementByTitleRequest) (*pb.SearchAdvertisementByTitleResponse, error)
}

type AdvertisementsRepository interface {
	Create(ctx context.Context, req *pb.CreateAdvertisementRequest) (*pb.CreateAdvertisementResponse, error)
	GetByID(ctx context.Context, req *pb.GetAdvertisementByIdRequest) (*pb.GetAdvertisementByIdResponse, error)
	GetAll(ctx context.Context, req *pb.GetAllAdvertisementsRequest) (*pb.GetAllAdvertisementsResponse, error)
	Update(ctx context.Context, req *pb.UpdateAdvertisementByIdRequest) (*pb.UpdateAdvertisementByIdResponse, error)
	Delete(ctx context.Context, req *pb.DeleteAdvertisementByIDRequest) (*pb.DeleteAdvertisementByIDResponse, error)
}

type UserRepository interface {
	Create(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	GetByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIdResponse, error)
	GetByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error)
	GetAll(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error)
	Update(ctx context.Context, req *pb.UpdateUserByIdRequest) (*pb.UpdateUserByIdResponse, error)
	Delete(ctx context.Context, req *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error)
}

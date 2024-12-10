package storage

import (
	"context"
	"project/models"
)

type StorageI interface {
	Users() UsersRepoI
}

type UsersRepoI interface {
	CreateUser(ctx context.Context, req *models.CreateUser) (*models.User, error)
	GetByIdUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetListUser(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	UpdateUser(ctx context.Context, req *models.UpdateUser) (int64, error)
	DeleteUser(ctx context.Context, req *models.UserPrimaryKey) error
}

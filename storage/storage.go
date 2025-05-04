package storage

import (
	"context"
	"minimedium/models"
)

type StorageI interface {
	Users() UsersRepoI
	Posts() PostsRepoI
	Login() LoginRepoI
	Follow() FollowRepoI
}

type UsersRepoI interface {
	CreateUser(ctx context.Context, req *models.CreateUser) (*models.User, error)
	GetByIdUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetListUser(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	UpdateUser(ctx context.Context, req *models.UpdateUser) (int64, error)
	DeleteUser(ctx context.Context, req *models.UserPrimaryKey) error
}

type PostsRepoI interface {
	CreatePost(ctx context.Context, req *models.CreatePost) (*models.Post, error)
	GetByIdPost(ctx context.Context, req *models.PostPrimaryKey) (*models.Post, error)
	UpdatePost(ctx context.Context, req *models.UpdatePost) (int64, error)
	GetListPost(ctx context.Context, req *models.GetListPostRequest) (*models.GetListPostResponse, error)
	DeletePost(ctx context.Context, req *models.PostPrimaryKey) error
}
type LoginRepoI interface {
	SignIn(ctx context.Context, req *models.SignIn) (*models.SignInResponse, error)
}
type FollowRepoI interface {
	Follow(ctx context.Context, req *models.Follow) error

}


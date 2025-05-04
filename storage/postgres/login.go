package postgres

import (
	"context"
	"database/sql"
	"minimedium/models"

	"github.com/jackc/pgx/v4/pgxpool"
)


type login struct {
	db *pgxpool.Pool
}

func NewSignRepo(db *pgxpool.Pool) *login {
	return &login{
		db: db,
	}
}

func (s *login) SignIn(ctx context.Context, req *models.SignIn) (*models.SignInResponse, error) {

	var (
		query = `
		
			SELECT 
				"id",
				"name",
				"user_name",
				"bio",
				"media_url",
				"password",
				"created_at",
				"updated_at"
			FROM "users"
			WHERE "user_name" = $1 AND "password" = $2
		`
	)
	var (
		id         sql.NullString
		name       sql.NullString
		user_name  sql.NullString
		bio        sql.NullString
		media_url  sql.NullString
		password   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)
	err := s.db.QueryRow(ctx, query, req.UserName, req.Password).Scan(
		&id,
		&name,
		&user_name,
		&bio,
		&media_url,
		&password,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &models.SignInResponse{
		Id:        id.String,
		Name:      name.String,
		UserName:  user_name.String,
		Bio:       bio.String,
		MediaUrl:  media_url.String,
		Password:  password.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"project/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (p *postRepo) CreatePost(ctx context.Context, req *models.CreatePost) (*models.Post, error) {
	var (
		postId = uuid.New().String()

		query = `
		
		INSERT INTO "post" (
			"id",
			"title",
			"body",
			"user_id",
			"updated_at"
		) VALUES ($1, $2, $3, $4, NOW())
		`
	)
	_, err := p.db.Exec(ctx,query,
		postId,
		req.Title,
		req.Body,
		req.UserId,
	)

	if err != nil {
		return nil, err
	}
	return p.GetByIdPost(ctx, &models.PostPrimaryKey{Id: postId})
}

func (p *postRepo) GetByIdPost(ctx context.Context, req *models.PostPrimaryKey) (*models.Post, error) {

	var (
		query = `
		
		SELECT 
			"id",
			"title",
			"body",
			"user_id",
			"created_at",
			"updated_at"
		FROM post
		WHERE "id" = $1
		`
	)

	var (
		id         sql.NullString
		title      sql.NullString
		body       sql.NullString
		userId     sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	err := p.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&title,
		&body,
		&userId,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &models.Post{
		Id:        id.String,
		Title:     title.String,
		Body:      body.String,
		UserId:    userId.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (p *postRepo) GetListPost(ctx context.Context, req *models.GetListPostRequest) (*models.GetListPostResponse, error) {
	var (
		resp       models.GetListPostResponse
		where      = " WHERE TRUE"
		page_limit = " OFFSET 0"
		limit      = " LIMIT 10"
		sort       = " ORDER BY created_at DESC"
	)
	if req.Page > 0 {
		page_limit = fmt.Sprintf(" OFFSET %d ", req.Page)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d ", req.Limit)
	}
	if len(req.Search) > 0 {
		where += " AND body ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
		SELECT 
			COUNT(*) OVER(),
			"id",
			"title",
			"body",
			"user_id",
			"created_at",
			"updated_at"
		FROM post	`

	query += where + sort + page_limit + limit
	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			Id        sql.NullString
			Title     sql.NullString
			Body      sql.NullString
			UserId    sql.NullString
			CreatedAt sql.NullString
			UpdatedAt sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&Id,
			&Title,
			&Body,
			&UserId,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Posts = append(resp.Posts, &models.Post{
			Id:        Id.String,
			Title:     Title.String,
			Body:      Body.String,
			UserId:    UserId.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		})

	}
	return &resp, nil
}

func (p *postRepo) DeletePost(ctx context.Context, req *models.PostPrimaryKey) error {
	_, err := p.db.Exec(ctx, "DELETE FROM post WHERE id = $1", req.Id)
	return err
}

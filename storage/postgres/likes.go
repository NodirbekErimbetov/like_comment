package postgres

import (
	"context"
	"fmt"
	"minimedium/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type likeRepo struct {
	db *pgxpool.Pool
}

func NewLikeRepo(db *pgxpool.Pool) *likeRepo {
	return &likeRepo{
		db: db,
	}
}

func (l *likeRepo) Like(ctx context.Context, req *models.Like) (int64, error) {

	var (
		query = `
				INSERT INTO "likes" (
					"post_id",
					"count"
				) VALUES($1,$2)

		`
	)
	fmt.Println(query)
	resp, err := l.db.Exec(ctx, query,
		req.PostId,
		req.Count,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected := resp.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

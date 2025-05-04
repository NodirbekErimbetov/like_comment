package postgres

import (
	"context"
	"minimedium/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type followRepo struct {
	db *pgxpool.Pool
}

func NewFollowRepo(db *pgxpool.Pool) *followRepo {
	return &followRepo{
		db: db,
	}
}
func (u *followRepo) Follow(ctx context.Context, req *models.Follow) error {

	var (
		query = `
			INSERT INTO "follow"(
				"follower_id",
				followed_id"
			)	VALUES ($1,$2)

		`
	)
	_, err := u.db.Exec(
		ctx,
		query,
		req.FollowerId,
		req.FollowedId,
	)
	if err != nil {
		return nil
	}
	return err
}

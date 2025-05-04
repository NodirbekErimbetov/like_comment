package postgres

import (
	"context"
	"fmt"
	"log"
	"minimedium/config"
	"minimedium/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db   *pgxpool.Pool
	user storage.UsersRepoI
	post storage.PostsRepoI
	sign storage.LoginRepoI
	follow storage.FollowRepoI

}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("dbname = %s password = %s port = %s user = %s host =%s  sslmode=disable",
			cfg.PostgresDatabase,
			cfg.PostgresPassword,
			cfg.PostgresPort,
			cfg.PostgresUser,
			cfg.PostgresHost,
		),
	)
	if err != nil {
		return nil, err
	}
	pgxpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Error while connecting to database !")
	}
	return &Store{
		db: pgxpool,
	}, nil
}

func (s *Store) Users() storage.UsersRepoI {
	if s.user == nil {
		s.user = NewUsersRepo(s.db)
	}
	return s.user
}

func (s *Store) Posts() storage.PostsRepoI {
	if s.post == nil {
		s.post = NewPostRepo(s.db)
	}
	return s.post
}

func (s *Store) Login() storage.LoginRepoI {
	if s.sign == nil {
		s.sign = NewSignRepo(s.db)
	}
	return s.sign
}

func (s *Store) Follow() storage.FollowRepoI{
	if s.follow == nil{
		s.follow = NewFollowRepo(s.db)
	}
	return s.follow
}



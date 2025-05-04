package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"minimedium/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUsersRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db}
}

func (u *userRepo) CreateUser(ctx context.Context, req *models.CreateUser) (*models.User, error) {
	var (
		userId = uuid.New().String()
		query  = `
		INSERT INTO "users"(
					"id",
					"name",
					"user_name",
					"bio",
					"media_url",
					"password",
					"updated_at"		
					) VALUES ($1, $2, $3, $4, $5,$6, NOW())
				`
	)

	_, err := u.db.Exec(ctx, query,
		userId,
		req.Name,
		req.UserName,
		req.Bio,
		req.MediaUrl,
		req.Password,
	)

	if err != nil {
		return nil, err
	}
	return u.GetByIdUser(ctx, &models.UserPrimaryKey{Id: userId})
}
func (u *userRepo) GetByIdUser(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

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
		WHERE id = $1 OR user_name = $1
		`
	)
	fmt.Println(query)
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

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
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

	return &models.User{
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

func (u *userRepo) GetListUser(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp       models.GetListUserResponse
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
		where += " AND first_name ILIKE" + " '%" + req.Search + "%'" + "OR last_name ILIKE" + " '%" + req.Search + "%'"
	}

	var query = `
	SELECT
			COUNT(*) OVER(),
			"id",
			"name",	
			"user_name",	
			"bio",	
			"media_url",	
			"password",	
			"created_at",
			"updated_at"
	FROM users
`
	query += where + sort + page_limit + limit
	rows, err := u.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
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

		err = rows.Scan(
			&resp.Count,
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

		resp.Users = append(resp.Users, &models.User{
			Id:        id.String,
			Name:      name.String,
			UserName:  user_name.String,
			Bio:       bio.String,
			MediaUrl:  media_url.String,
			Password:  password.String,
			CreatedAt: created_at.String,
			UpdatedAt: updated_at.String,
		})
	}
	return &resp, nil
}
func (u *userRepo) UpdateUser(ctx context.Context, req *models.UpdateUser) (int64, error) {

	query := `	
	UPDATE users
		SET 
			name =$2,
			user_name =$3,
			bio= $4,
			media_url = $5,
			password =$6,
			updated_at = NOW()
	WHERE id = $1
	`
	result, err := u.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.UserName,
		req.Bio,
		req.MediaUrl,
		req.Password,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (u *userRepo) DeleteUser(ctx context.Context, req *models.UserPrimaryKey) error {
	_, err := u.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	return err
}

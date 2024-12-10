package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"project/models"

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
					"first_name",
					"last_name",
					"email",
					"password",
					"updated_at"		
					) VALUES ($1, $2, $3, $4, $5, NOW())
				`
	)

	_, err := u.db.Exec(ctx, query,
		userId,
		req.FirstName,
		req.LastName,
		req.Email,
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
			"first_name",	
			"last_name",	
			"email",	
			"password",	
			"created_at",	
			"updated_at"
		FROM "users"
		WHERE "id" = $1
		`
	)
	var (
		id         sql.NullString
		first_name sql.NullString
		last_name  sql.NullString
		email      sql.NullString
		password   sql.NullString
		created_at sql.NullString
		updated_at sql.NullString
	)

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&first_name,
		&last_name,
		&email,
		&password,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		FirstName: first_name.String,
		LastName:  last_name.String,
		Email:     email.String,
		Password:  password.String,
		CreatedAt: created_at.String,
		UpdatedAt: updated_at.String,
	}, nil
}

func (u *userRepo) GetListUser(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {

	var (
		resp       models.GetListUserResponse
		where      = "WHERE TRUE"
		page_limit = "OFFSET 1 "
		limit      = "LIMIT 10 "
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
		"first_name",
		"last_name",
		"email",
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
			Id        sql.NullString
			FirstName sql.NullString
			LastName  sql.NullString
			Email     sql.NullString
			Password  sql.NullString
			CreatedAt sql.NullString
			UpdatedAt sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&Id,
			&FirstName,
			&LastName,
			&Email,
			&Password,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:        Id.String,
			FirstName: FirstName.String,
			LastName:  LastName.String,
			Email:     Email.String,
			Password:  Password.String,
			CreatedAt: CreatedAt.String,
			UpdatedAt: UpdatedAt.String,
		})
	}
	return &resp, nil
}
func (u *userRepo) UpdateUser(ctx context.Context, req *models.UpdateUser) (int64, error) {

	query := `	
	UPDATE users
		SET 
			first_name =$2,
			last_name =$3,
			password =$4,
			updated_at = NOW()
	WHERE id = $1
	`
	result, err := u.db.Exec(ctx, query,
		req.Id,
		req.FirstName,
		req.LastName,
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

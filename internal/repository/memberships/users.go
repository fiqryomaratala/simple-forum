package memberships

import (
	"context"

	"github.com/fiqryomaratala/simple-forum/internal/model/memberships"
)

//Membuat Repository untuk GetUser, untuk validasi apakah user sudah exist atau belum.

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, email, password, username, password, created_at, updated_at, created_by, updated_by
	FROM users WHERE email = $1 OR username = $2`
	row := r.db.QueryRowContext(ctx, query, email, username)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Password, &response.Username,
		&response.CreatedBy, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

//Membuat Repository untuk CreateUser

func (r *repository) CreateUser(ctx context.Context, model memberships.UserModel) error {
	query := `INSER INTO users (email, password, username, created_at, updated_at, created_by, 
	updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.ExecContext(ctx, query, model.Email, model.Password, model.Username,
		model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

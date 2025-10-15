package memberships

import (
	"context"

	"github.com/fiqryomaratala/simple-forum/internal/model/memberships"
)

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

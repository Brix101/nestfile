package repository

import (
	"context"

	"github.com/Brix101/nestfile/internal/domain"
)

type sqlUserRepository struct {
	conn Connection
}

func NewSqlUser(conn Connection) domain.UserRepository {
	return &sqlUserRepository{conn: conn}
}

func (p *sqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.User, error) {
	rows, err := p.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs := []domain.User{}
	for rows.Next() {

		var usr domain.User
		if err := rows.Scan(
			&usr.ID,
			&usr.Username,
			&usr.Password,
			&usr.CreatedAt,
			&usr.UpdatedAt,
		); err != nil {
			return nil, err
		}
		usrs = append(usrs, usr)
	}

	return usrs, nil
}

func (p *sqlUserRepository) GetByID(ctx context.Context, id int64) (domain.User, error) {
	query := `
		SELECT
			id,
			username,
			password,
			created_at,
			updated_at
		FROM
			users
		WHERE
			id = $1`

	usr, err := p.fetch(ctx, query, id)
	if err != nil {
		return domain.User{}, err
	}

	if len(usr) == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	return usr[0], nil
}

func (p *sqlUserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {
	query := `
		SELECT
			id,
			username,
			password,
			created_at,
			updated_at
		FROM
			users
		WHERE
			username = $1`

	usr, err := p.fetch(ctx, query, username)
	if err != nil {
		return domain.User{}, err
	}

	if len(usr) == 0 {
		return domain.User{}, domain.ErrNotFound
	}

	return usr[0], nil
}

func (p *sqlUserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := `
		SELECT
			id,
			username,
			password,
			created_at,
			updated_at
		FROM
			users`

	usrs, err := p.fetch(ctx, query)
	if err != nil {
		return []domain.User{}, err
	}

	return usrs, nil
}

func (p *sqlUserRepository) Update(ctx context.Context, usr *domain.User) (*domain.User, error) {
	query := `
		UPDATE users
		SET
			username = $2,
			email = $3,
			updated_at = NOW()
		WHERE
			id = $1
		RETURNING updated_at`

	row := p.conn.QueryRowContext(
		ctx,
		query,
		usr.ID,
		usr.Username,
		usr.Password,
	)

	if err := row.Scan(&usr.UpdatedAt); err != nil {
		return nil, err
	}

	return usr, nil
}

func (p *sqlUserRepository) Create(ctx context.Context, usr *domain.User) (*domain.User, error) {
	query := `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at`

	if err := p.conn.QueryRowContext(
		ctx,
		query,
		usr.Username,
		usr.Password,
	).Scan(
		&usr.ID,
		&usr.CreatedAt,
		&usr.UpdatedAt); err != nil {
		return nil, err
	}
	return usr, nil
}

func (p *sqlUserRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE users
		WHERE id = $1`

	result, err := p.conn.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

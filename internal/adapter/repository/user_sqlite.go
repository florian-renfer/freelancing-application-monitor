package repository

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserSQL struct {
	db SQL
}

func NewUserSQL(db SQL) UserSQL {
	return UserSQL{
		db: db,
	}
}

func (u UserSQL) Create(ctx context.Context, user domain.User) (domain.User, error) {
	var query = `
		INSERT INTO 
			users (id, email, password, firstname, lastname, date_of_birth, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8)
	`

	if err := u.db.ExecuteContext(
		ctx,
		query,
		user.ID(),
		user.Email(),
		user.Password(),
		user.Firstname(),
		user.Lastname(),
		user.DateOfBirth(),
		user.CreatedAt(),
		user.UpdatedAt(),
	); err != nil {
		return domain.User{}, errors.Wrap(err, "error creating user")
	}

	return user, nil
}

func (u UserSQL) FindAll(ctx context.Context) ([]domain.User, error) {
	var query = "SELECT id, email, firstname, lastname, date_of_birth, created_at, updated_at FROM users"

	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return []domain.User{}, errors.Wrap(err, "error listing users")
	}

	var users = make([]domain.User, 0)
	for rows.Next() {
		var (
			ID          uuid.UUID
			Email       string
			Firstname   string
			Lastname    string
			DateOfBirth time.Time
			CreatedAt   time.Time
			UpdatedAt   time.Time
		)
		if err = rows.Scan(&ID, &Email, &Firstname, &Lastname, &DateOfBirth, &CreatedAt, &UpdatedAt); err != nil {
			return []domain.User{}, errors.Wrap(err, "error listing users")
		}

		users = append(users, domain.NewUser(
			ID,
			Email,
			"",
			Firstname,
			Lastname,
			DateOfBirth,
			CreatedAt,
			UpdatedAt,
		))
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (u UserSQL) DeleteById(ctx context.Context, id uuid.UUID) error {
	var query = "DELETE FROM users WHERE id = $1"

	if err := u.db.ExecuteContext(
		ctx,
		query,
		id,
	); err != nil {
		return errors.Wrap(err, "error deleting user")
	}

	return nil
}

package repository

import (
	"context"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
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

func (a UserSQL) Create(ctx context.Context, user domain.User) (domain.User, error) {
	var query = `
		INSERT INTO 
			users (id, email, password, firstname, lastname, date_of_birth, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8)
	`

	if err := a.db.ExecuteContext(
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

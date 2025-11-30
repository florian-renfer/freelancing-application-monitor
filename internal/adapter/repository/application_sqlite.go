package repository

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ApplicationSQL struct {
	db SQL
}

func NewApplicationSQL(db SQL) ApplicationSQL {
	return ApplicationSQL{
		db: db,
	}
}

func (a ApplicationSQL) Create(ctx context.Context, application domain.Application) (domain.Application, error) {
	var query = `
		INSERT INTO 
			applications (id, title, description, url, state, applied_at, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8)
	`

	if err := a.db.ExecuteContext(
		ctx,
		query,
		application.Id(),
		application.Title(),
		application.Description(),
		application.Url(),
		application.State().String(),
		application.AppliedAt(),
		application.CreatedAt(),
		application.UpdatedAt(),
	); err != nil {
		return domain.Application{}, errors.Wrap(err, "error creating application")
	}

	return application, nil
}

func (a ApplicationSQL) FindAll(ctx context.Context) ([]domain.Application, error) {
	var query = "SELECT * FROM applications"

	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return []domain.Application{}, errors.Wrap(err, "error listing applications")
	}

	var applications = make([]domain.Application, 0)
	for rows.Next() {
		var (
			id          uuid.UUID
			title       string
			description string
			url         string
			state       domain.ApplicationState
			appliedAt   time.Time
			createdAt   time.Time
			updatedAt   time.Time
		)

		if err = rows.Scan(&id, &title, &description, &url, &state, &appliedAt, &createdAt, &updatedAt); err != nil {
			return []domain.Application{}, errors.Wrap(err, "error listing applications")
		}

		applications = append(applications, domain.NewApplication(
			id,
			title,
			description,
			url,
			state,
			appliedAt,
			createdAt,
			updatedAt,
		))
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return []domain.Application{}, err
	}

	return applications, nil
}

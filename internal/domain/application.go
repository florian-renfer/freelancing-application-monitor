package domain

import (
	"context"
	"net/url"
	"time"

	"github.com/google/uuid"
)

type (
	Application struct {
		id          uuid.UUID
		title       string
		description string
		url         url.URL
		appliedAt   time.Time
		createdAt   time.Time
	}

	ApplicationRepository interface {
		Create(context.Context, Application) (Application, error)
		FindAll(context.Context) ([]Application, error)
	}
)

func NewApplication(id uuid.UUID, title, description string, url url.URL, appliedAt, createdAt time.Time) Application {
	return Application{
		id:          id,
		title:       title,
		description: description,
		url:         url,
		appliedAt:   time.Now(),
		createdAt:   time.Now(),
	}
}

func (a Application) Id() uuid.UUID {
	return a.id
}

func (a Application) Title() string {
	return a.title
}

func (a Application) Description() string {
	return a.description
}

func (a Application) Url() url.URL {
	return a.url
}

func (a Application) AppliedAt() time.Time {
	return a.appliedAt
}

func (a Application) CreatedAt() time.Time {
	return a.createdAt
}

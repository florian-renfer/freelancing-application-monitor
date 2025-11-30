package domain

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	ApplicationState int

	Application struct {
		id          uuid.UUID
		title       string
		description string
		url         string
		state       ApplicationState
		appliedAt   time.Time
		createdAt   time.Time
		updatedAt   time.Time
	}

	ApplicationRepository interface {
		Create(context.Context, Application) (Application, error)
		FindAll(context.Context) ([]Application, error)
	}
)

const (
	DRAFT ApplicationState = iota
	APPLIED
	OFFERED
	ACCEPTED
	REJECTED
	WITHDRAWN
	CLOSED
)

var (
	errInvalidApplicationState = errors.New("invalid application state")
	applicationStateStrings    = [...]string{"DRAFT", "APPLIED", "OFFERED", "ACCEPTED", "REJECTED", "WITHDRAWN", "CLOSED"}
)

func (s ApplicationState) String() string {
	return [...]string{"DRAFT", "APPLIED", "OFFERED", "ACCEPTED", "REJECTED", "WITHDRAWN", "CLOSED"}[s]
}

func (s *ApplicationState) Scan(value any) error {
	switch v := value.(type) {
	case int64:
		*s = ApplicationState(v)
		return nil
	case []byte:
		return s.Scan(string(v))
	case string:
		for i, str := range applicationStateStrings {
			if str == v {
				*s = ApplicationState(i)
				return nil
			}
		}
		return errInvalidApplicationState
	default:
		return errInvalidApplicationState
	}
}

func (s *ApplicationState) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	for i, v := range applicationStateStrings {
		if v == str {
			*s = ApplicationState(i)
			return nil
		}
	}
	return errInvalidApplicationState
}

func (s ApplicationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func NewApplication(id uuid.UUID, title, description string, url string, state ApplicationState, appliedAt, createdAt, updatedAt time.Time) Application {
	return Application{
		id:          id,
		title:       title,
		description: description,
		url:         url,
		state:       state,
		appliedAt:   appliedAt,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
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

func (a Application) Url() string {
	return a.url
}

func (a Application) State() ApplicationState {
	return a.state
}

func (a Application) AppliedAt() time.Time {
	return a.appliedAt
}

func (a Application) CreatedAt() time.Time {
	return a.createdAt
}

func (a Application) UpdatedAt() time.Time {
	return a.updatedAt
}

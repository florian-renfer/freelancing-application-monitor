package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type (
	// CreateApplicationUseCase input port
	CreateApplicationUseCase interface {
		Execute(context.Context, CreateApplicationInput) (CreateApplicationOutput, error)
	}

	// CreateApplicationInput input data
	CreateApplicationInput struct {
		Title       string                  `json:"title" validate:"required,min=10,max=255"`
		Description string                  `json:"description" validate:"required,min=10,max=3000"`
		Url         string                  `json:"url" validate:"required,url,startswith=https://,max=512"`
		State       domain.ApplicationState `json:"state" validate:"required"`
		AppliedAt   time.Time               `json:"applied_at" validate:"required,lt"`
	}

	// CreateApplicationPresenter output port
	CreateApplicationPresenter interface {
		Output(domain.Application) CreateApplicationOutput
	}

	// CreateApplicationOutput output data
	CreateApplicationOutput struct {
		Id          uuid.UUID               `json:"id"`
		Title       string                  `json:"title"`
		Description string                  `json:"description"`
		Url         string                  `json:"url"`
		State       domain.ApplicationState `json:"state"`
		AppliedAt   time.Time               `json:"applied_at"`
		CreatedAt   time.Time               `json:"created_at"`
		UpdatedAt   time.Time               `json:"updated_at"`
	}

	createApplicationInteractor struct {
		repo       domain.ApplicationRepository
		presenter  CreateApplicationPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateApplicationInteractor(
	repo domain.ApplicationRepository,
	presenter CreateApplicationPresenter,
	t time.Duration,
) CreateApplicationUseCase {
	return createApplicationInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a createApplicationInteractor) Execute(ctx context.Context, input CreateApplicationInput) (CreateApplicationOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	var application = domain.NewApplication(
		uuid.New(),
		input.Title,
		input.Description,
		input.Url,
		input.State,
		input.AppliedAt,
		time.Now(),
		time.Now(),
	)

	application, err := a.repo.Create(ctx, application)
	if err != nil {
		return a.presenter.Output(domain.Application{}), err
	}

	return a.presenter.Output(application), nil
}

package usecase

import (
	"context"
	"net/url"
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
		Title       string  `json:"title" validate:"required"`
		Description string  `json:"description" validate:"required"`
		Url         url.URL `json:"url" validate:"required"`
	}

	// CreateApplicationPresenter output port
	CreateApplicationPresenter interface {
		Output(domain.Application) CreateApplicationOutput
	}

	// CreateApplicationOutput output data
	CreateApplicationOutput struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		CPF       string  `json:"cpf"`
		Balance   float64 `json:"balance"`
		CreatedAt string  `json:"created_at"`
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
		time.Now(),
		time.Now(),
	)

	application, err := a.repo.Create(ctx, application)
	if err != nil {
		return a.presenter.Output(domain.Application{}), err
	}

	return a.presenter.Output(application), nil
}

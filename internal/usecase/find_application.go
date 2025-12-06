package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type (
	// FindApplicationUseCase input port
	FindApplicationUseCase interface {
		Execute(context.Context, FindApplicationInput) (FindApplicationOutput, error)
	}

	// FindApplicationInput input data
	FindApplicationInput uuid.UUID

	// FindApplicationPresenter output port
	FindApplicationPresenter interface {
		Output(domain.Application) FindApplicationOutput
	}

	// FindApplicationOutput output data
	FindApplicationOutput struct {
		Id          uuid.UUID               `json:"id"`
		Title       string                  `json:"title"`
		Description string                  `json:"description"`
		Url         string                  `json:"url"`
		State       domain.ApplicationState `json:"state"`
		AppliedAt   time.Time               `json:"applied_at"`
		CreatedAt   time.Time               `json:"created_at"`
		UpdatedAt   time.Time               `json:"updated_at"`
	}

	findApplicationInteractor struct {
		repo       domain.ApplicationRepository
		presenter  FindApplicationPresenter
		ctxTimeout time.Duration
	}
)

func NewFindApplicationInteractor(repo domain.ApplicationRepository, presenter FindApplicationPresenter, timeout time.Duration) FindApplicationUseCase {
	return &findApplicationInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func (c findApplicationInteractor) Execute(ctx context.Context, input FindApplicationInput) (FindApplicationOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	application, err := c.repo.FindById(ctx, uuid.UUID(input))
	if err != nil {
		return FindApplicationOutput{}, err
	}

	return c.presenter.Output(application), nil
}

package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type (
	// FindAllApplicationUseCase input port
	FindAllApplicationUseCase interface {
		Execute(context.Context) ([]FindAllApplicationOutput, error)
	}

	// FindAllApplicationPresenter output port
	FindAllApplicationPresenter interface {
		Output([]domain.Application) []FindAllApplicationOutput
	}

	// FindAllApplicationOutput output data
	FindAllApplicationOutput struct {
		Id          uuid.UUID               `json:"id"`
		Title       string                  `json:"title"`
		Description string                  `json:"description"`
		Url         string                  `json:"url"`
		State       domain.ApplicationState `json:"state"`
		AppliedAt   time.Time               `json:"applied_at"`
		CreatedAt   time.Time               `json:"created_at"`
		UpdatedAt   time.Time               `json:"updated_at"`
	}

	// FindAllApplicationUseCase implementation
	findAllApplicationInteractor struct {
		repo       domain.ApplicationRepository
		presenter  FindAllApplicationPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllApplicationInteractor(repo domain.ApplicationRepository, presenter FindAllApplicationPresenter, timeout time.Duration) FindAllApplicationUseCase {
	return &findAllApplicationInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func (c findAllApplicationInteractor) Execute(ctx context.Context) ([]FindAllApplicationOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	applications, err := c.repo.FindAll(ctx)
	if err != nil {
		return c.presenter.Output([]domain.Application{}), err
	}

	return c.presenter.Output(applications), nil
}

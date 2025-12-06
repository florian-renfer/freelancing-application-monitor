package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type (
	// FindAllUserUseCase input port
	FindAllUserUseCase interface {
		Execute(context.Context) ([]FindAllUserOutput, error)
	}

	// FindAllUserPresenter output port
	FindAllUserPresenter interface {
		Output([]domain.User) []FindAllUserOutput
	}

	// FindAllUserOutput output data
	FindAllUserOutput struct {
		ID          uuid.UUID `json:"id"`
		Email       string    `json:"email"`
		Firstname   string    `json:"firstname"`
		Lastname    string    `json:"lastname"`
		DateOfBirth time.Time `json:"date_of_birth"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// FindAllUserUseCase implementation
	findAllUserInteractor struct {
		repo       domain.UserRepository
		presenter  FindAllUserPresenter
		ctxTimeout time.Duration
	}
)

func NewFindAllUserInteractor(repo domain.UserRepository, presenter FindAllUserPresenter, timeout time.Duration) FindAllUserUseCase {
	return &findAllUserInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func (c findAllUserInteractor) Execute(ctx context.Context) ([]FindAllUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	users, err := c.repo.FindAll(ctx)
	if err != nil {
		return c.presenter.Output([]domain.User{}), err
	}

	return c.presenter.Output(users), nil
}

package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type (
	// DeleteUserUseCase input port
	DeleteUserUseCase interface {
		Execute(context.Context, DeleteUserInput) error
	}

	// DeleteUserInput input data
	DeleteUserInput uuid.UUID

	// DeleteUserUseCase implementation
	deleteUserInteractor struct {
		repo       domain.UserRepository
		ctxTimeout time.Duration
	}
)

func NewDeleteUserInteractor(repo domain.UserRepository, timeout time.Duration) DeleteUserUseCase {
	return &deleteUserInteractor{
		repo:       repo,
		ctxTimeout: timeout,
	}
}

func (c deleteUserInteractor) Execute(ctx context.Context, input DeleteUserInput) error {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	err := c.repo.DeleteById(ctx, uuid.UUID(input))
	if err != nil {
		return err
	}

	return nil
}

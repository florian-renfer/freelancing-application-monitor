package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/google/uuid"
)

type (
	// CreateUserUseCase input port
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	// CreateUserInput input data
	CreateUserInput struct {
		Email       string    `json:"email" validate:"required"`
		Password    string    `json:"password" validate:"required"`
		Firstname   string    `json:"firstname" validate:"required"`
		Lastname    string    `json:"lastname" validate:"required"`
		DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
	}

	// CreateUserPresenter output port
	CreateUserPresenter interface {
		Output(domain.User) CreateUserOutput
	}

	// CreateUserOutput output data
	CreateUserOutput struct {
		ID          uuid.UUID `json:"id"`
		Email       string    `json:"title"`
		Firstname   string    `json:"firstname"`
		Lastname    string    `json:"lastname"`
		DateOfBirth time.Time `json:"date_of_birth"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	// CreateUserUseCase implementation
	createUserInteractor struct {
		repo       domain.UserRepository
		presenter  CreateUserPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateUserInteractor(repo domain.UserRepository, presenter CreateUserPresenter, timeout time.Duration) CreateUserUseCase {
	return &createUserInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func (c createUserInteractor) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	var user = domain.NewUser(
		uuid.New(),
		input.Email,
		input.Password,
		input.Firstname,
		input.Lastname,
		input.DateOfBirth,
		time.Now(),
		time.Now(),
	)

	user, err := c.repo.Create(ctx, user)
	if err != nil {
		return c.presenter.Output(domain.User{}), err
	}

	return c.presenter.Output(user), nil
}

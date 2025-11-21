package presenter

import (
	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type createUserPresenter struct{}

func NewCreateUserPresenter() usecase.CreateUserPresenter {
	return createUserPresenter{}
}

func (a createUserPresenter) Output(user domain.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID:          user.ID(),
		Email:       user.Email(),
		Firstname:   user.Firstname(),
		Lastname:    user.Lastname(),
		DateOfBirth: user.DateOfBirth(),
		CreatedAt:   user.CreatedAt(),
		UpdatedAt:   user.UpdatedAt(),
	}
}

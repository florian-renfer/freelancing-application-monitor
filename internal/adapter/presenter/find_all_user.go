package presenter

import (
	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type findAllUserPresenter struct{}

func NewFindAllUserPresenter() usecase.FindAllUserPresenter {
	return findAllUserPresenter{}
}

func (f findAllUserPresenter) Output(users []domain.User) []usecase.FindAllUserOutput {
	var o = make([]usecase.FindAllUserOutput, 0)

	for _, user := range users {
		o = append(o, usecase.FindAllUserOutput{
			ID:          user.ID(),
			Email:       user.Email(),
			Firstname:   user.Firstname(),
			Lastname:    user.Lastname(),
			DateOfBirth: user.DateOfBirth(),
			CreatedAt:   user.CreatedAt(),
			UpdatedAt:   user.UpdatedAt(),
		})
	}

	return o
}

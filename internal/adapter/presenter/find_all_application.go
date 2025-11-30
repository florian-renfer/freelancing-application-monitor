package presenter

import (
	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type findAllApplicationPresenter struct{}

func NewFindAllApplicationPresenter() usecase.FindAllApplicationPresenter {
	return findAllApplicationPresenter{}
}

func (a findAllApplicationPresenter) Output(applications []domain.Application) []usecase.FindAllApplicationOutput {
	var out = make([]usecase.FindAllApplicationOutput, 0)
	for _, application := range applications {
		out = append(out, usecase.FindAllApplicationOutput{
			Id:          application.Id(),
			Title:       application.Title(),
			Description: application.Description(),
			Url:         application.Url(),
			State:       application.State(),
			AppliedAt:   application.AppliedAt(),
			CreatedAt:   application.CreatedAt(),
			UpdatedAt:   application.UpdatedAt(),
		})
	}
	return out
}

package presenter

import (
	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type createApplicationPresenter struct{}

func NewCreateApplicationPresenter() usecase.CreateApplicationPresenter {
	return createApplicationPresenter{}
}

func (a createApplicationPresenter) Output(application domain.Application) usecase.CreateApplicationOutput {
	return usecase.CreateApplicationOutput{
		Id:          application.Id(),
		Title:       application.Title(),
		Description: application.Description(),
		Url:         application.Url(),
		State:       application.State(),
		AppliedAt:   application.AppliedAt(),
		CreatedAt:   application.CreatedAt(),
		UpdatedAt:   application.UpdatedAt(),
	}
}

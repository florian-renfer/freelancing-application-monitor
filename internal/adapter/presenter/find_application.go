package presenter

import (
	"github.com/florian-renfer/freelancing-application-monitor/internal/domain"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type findApplicationPresenter struct{}

func NewFindApplicationPresenter() usecase.FindApplicationPresenter {
	return findApplicationPresenter{}
}

func (a findApplicationPresenter) Output(application domain.Application) usecase.FindApplicationOutput {
	var out = usecase.FindApplicationOutput{
		Id:          application.Id(),
		Title:       application.Title(),
		Description: application.Description(),
		Url:         application.Url(),
		State:       application.State(),
		AppliedAt:   application.AppliedAt(),
		CreatedAt:   application.CreatedAt(),
		UpdatedAt:   application.UpdatedAt(),
	}
	return out
}

package action

import (
	"net/http"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/logging"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/response"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type FindAllApplicationAction struct {
	uc  usecase.FindAllApplicationUseCase
	log logger.Logger
}

func NewFindAllApplicationAction(uc usecase.FindAllApplicationUseCase, log logger.Logger) FindAllApplicationAction {
	return FindAllApplicationAction{
		uc:  uc,
		log: log,
	}
}

func (a FindAllApplicationAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_application"

	output, err := a.uc.Execute(r.Context())
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error finding all applications")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success finding all applications")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

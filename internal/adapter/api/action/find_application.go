package action

import (
	"net/http"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/logging"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/response"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type FindApplicationAction struct {
	uc  usecase.FindApplicationUseCase
	log logger.Logger
}

func NewFindApplicationAction(uc usecase.FindApplicationUseCase, log logger.Logger) FindApplicationAction {
	return FindApplicationAction{
		uc:  uc,
		log: log,
	}
}

func (a FindApplicationAction) Execute(w http.ResponseWriter, r *http.Request, id usecase.FindApplicationInput) {
	const logKey = "find_application"

	// FIXME: 404 should be returned on invalid uuid or nonexisting data
	application, err := a.uc.Execute(r.Context(), id)
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error finding application")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success finding application")

	response.NewSuccess(application, http.StatusOK).Send(w)
}

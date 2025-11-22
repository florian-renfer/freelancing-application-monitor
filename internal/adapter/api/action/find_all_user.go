package action

import (
	"net/http"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/logging"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/response"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type FindAllUserAction struct {
	uc  usecase.FindAllUserUseCase
	log logger.Logger
}

func NewFindAllUserAction(uc usecase.FindAllUserUseCase, log logger.Logger) FindAllUserAction {
	return FindAllUserAction{
		uc:  uc,
		log: log,
	}
}

func (a FindAllUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_user"

	output, err := a.uc.Execute(r.Context())
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error findin all users")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success finding all users")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

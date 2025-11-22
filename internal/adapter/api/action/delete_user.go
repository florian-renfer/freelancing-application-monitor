package action

import (
	"net/http"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/logging"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/response"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type DeleteUserAction struct {
	uc  usecase.DeleteUserUseCase
	log logger.Logger
}

func NewDeleteUserAction(uc usecase.DeleteUserUseCase, log logger.Logger) DeleteUserAction {
	return DeleteUserAction{
		uc:  uc,
		log: log,
	}
}

func (a DeleteUserAction) Execute(w http.ResponseWriter, r *http.Request, id usecase.DeleteUserInput) {
	const logKey = "delete_user"

	err := a.uc.Execute(r.Context(), id)
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error deleting user")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success finding all users")

	response.NewSuccess(nil, http.StatusNoContent).Send(w)
}

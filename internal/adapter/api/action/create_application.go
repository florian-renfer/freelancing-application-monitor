package action

import (
	"encoding/json"
	"net/http"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/logging"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/response"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type CreateApplicationAction struct {
	uc  usecase.CreateApplicationUseCase
	log logger.Logger
}

func NewCreateApplicationAction(uc usecase.CreateApplicationUseCase, log logger.Logger) CreateApplicationAction {
	return CreateApplicationAction{
		uc:  uc,
		log: log,
	}
}

func (a CreateApplicationAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_application"

	var input usecase.CreateApplicationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error when decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	// TODO: add validation

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when creating a new application")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success creating application")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

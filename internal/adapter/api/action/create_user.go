package action

import (
	"encoding/json"
	"net/http"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/logging"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/response"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"
)

type CreateUserAction struct {
	uc  usecase.CreateUserUseCase
	log logger.Logger
}

func NewCreateUserAction(uc usecase.CreateUserUseCase, log logger.Logger) CreateUserAction {
	return CreateUserAction{
		uc:  uc,
		log: log,
	}
}

func (a CreateUserAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_user"

	var input usecase.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	// TODO: add validation
	// if errs := a.validateInput(input); len(errs) > 0 {
	// 	logging.NewError(
	// 		a.log,
	// 		response.ErrInvalidInput,
	// 		logKey,
	// 		http.StatusBadRequest,
	// 	).Log("invalid input")
	//
	// 	response.NewErrorMessage(errs, http.StatusBadRequest).Send(w)
	// 	return
	// }

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error creating user")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success creating user")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

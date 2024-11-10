package user

import (
	"encoding/json"
	"net/http"

	"github.com/VsenseTechnologies/skf_plc_http_server/application/usecase/user"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/response"
	"github.com/gorilla/mux"
)

type Handler struct {
	dbRepo repository.DataBaseRepository
}

func NewHandler(dbRepo repository.DataBaseRepository) *Handler {
	return &Handler{
		dbRepo,
	}
}

func (h *Handler) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var request request.UserLogin

	if error := json.NewDecoder(r.Body).Decode(&request); error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	userLoginUseCase := user.InitUserLoginUseCase(h.dbRepo)

	error, errorStatus, token := userLoginUseCase.Execute(&request)

	if error != nil {
		response := &response.StatusMessage{
			Message: error.Error(),
		}
		if errorStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &response.Token{
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetDriersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]

	getDriersUseCase := user.InitGetDriersUseCase(h.dbRepo)

	error, errorStatus, driers := getDriersUseCase.Execute(userId)

	if error != nil {
		response := &response.StatusMessage{
			Message: error.Error(),
		}
		if errorStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &response.Driers{
		Driers: driers,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetRecipeStepCountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	drierId := vars["drierId"]

	getRecipeSteCountUseCase := user.InitGetRecipeStepCountUseCase(h.dbRepo)

	error, errorStatus, recipeStepCount := getRecipeSteCountUseCase.Execute(drierId)

	if error != nil {
		response := &response.StatusMessage{
			Message: error.Error(),
		}
		if errorStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &response.RecipeStepCount{
		RecipeStepCount: recipeStepCount,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetDrierStatusesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plcId := vars["plcId"]
	drierId := vars["drierId"]

	getDrierStatusesUseCase := user.InitGetDrierStatusesUseCase(h.dbRepo)

	err, errStatus, drierStatuses := getDrierStatusesUseCase.Execute(plcId, drierId)

	if err != nil {
		response := &response.StatusMessage{
			Message: err.Error(),
		}

		if errStatus == 1 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &response.DrierStatuses{
		DrierStatuses: drierStatuses,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

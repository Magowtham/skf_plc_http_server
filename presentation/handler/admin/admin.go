package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/VsenseTechnologies/skf_plc_http_server/application/usecase/admin"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/request"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/response"
	"github.com/gorilla/mux"
)

type Handler struct {
	dbRepo    repository.DataBaseRepository
	cacheRepo repository.CacheRepository
	smtpRepo  repository.SmtpClientRepository
}

func NewHandler(dbRepo repository.DataBaseRepository, cacheRepo repository.CacheRepository, smtpRepo repository.SmtpClientRepository) *Handler {
	return &Handler{
		dbRepo,
		cacheRepo,
		smtpRepo,
	}
}

func (h *Handler) DatabaseInitializeHandler(w http.ResponseWriter, r *http.Request) {
	initializeDataBaseUseCase := admin.InitInitializeDataBaseUseCase(h.dbRepo)
	error, errorStatus := initializeDataBaseUseCase.Execute()

	if error != nil {
		response := &response.StatusMessage{
			Message: "failed to initialize database",
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

	response := &response.StatusMessage{
		Message: "database initialized successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var request request.Admin

	error := json.NewDecoder(r.Body).Decode(&request)

	if error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	signUpUseCase := admin.InitCreateAdminUseCase(h.dbRepo)

	error, errorStatus := signUpUseCase.Execute(&request)

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

	response := &response.StatusMessage{
		Message: "admin created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteAdminHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	adminId := vars["adminId"]

	adminDeleteUseCase := admin.InitDeleteAdminUaseCase(h.dbRepo)

	error, errorStatus := adminDeleteUseCase.Execute(adminId)

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

	response := &response.StatusMessage{
		Message: "admin deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) AdminLoginHandler(w http.ResponseWriter, r *http.Request) {
	var request request.AdminLogin

	error := json.NewDecoder(r.Body).Decode(&request)

	if error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	adminLoginUseCase := admin.InitAdminLoginUseCase(h.dbRepo)

	error, errorStatus, token := adminLoginUseCase.Execute(request)

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

	response := &response.StatusMessage{
		Message: "login successfull",
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Domain:   "skfplc.http.vsensetech.in",
		Expires:  time.Now().Add(time.Hour * 24 * 365),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var request request.User

	error := json.NewDecoder(r.Body).Decode(&request)

	if error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	createUserUseCase := admin.InitCreateUserUseCase(h.dbRepo)

	error, errorStatus := createUserUseCase.Execute(&request)

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

	response := &response.StatusMessage{
		Message: "user created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]

	deleteUserUseCase := admin.InitDeleteUserCase(h.dbRepo, h.cacheRepo)

	error, errorStatus := deleteUserUseCase.Execute(userId)

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

	response := &response.StatusMessage{
		Message: "user deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	getAllUsersUseCase := admin.InitGetAllUsersUseCase(h.dbRepo)
	error, _, users := getAllUsersUseCase.Execute()

	if error != nil {
		response := &response.StatusMessage{
			Message: error.Error(),
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := &response.AllUsers{
		Users: users,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreatePlcHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	var request request.Plc

	error := json.NewDecoder(r.Body).Decode(&request)

	if error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	createPlcUseCase := admin.InitCreatePlcUseCase(h.dbRepo)

	error, errorStatus := createPlcUseCase.Execute(userId, &request)

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

	response := &response.StatusMessage{
		Message: "plc created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeletePlcHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]

	deletePlcUseCase := admin.InitDeletePlcUseCase(h.dbRepo, h.cacheRepo)

	if error, errorStatus := deletePlcUseCase.Execute(plcId); error != nil {
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

	response := &response.StatusMessage{
		Message: "plc deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetPlcsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]

	getPlcsUseCase := admin.InitGetPlcsUseCase(h.dbRepo)

	error, errorStatus, plcs := getPlcsUseCase.Execute(userId)

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

	response := &response.Plcs{
		Plcs: plcs,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreateDrierHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]

	var request request.Drier

	if error := json.NewDecoder(r.Body).Decode(&request); error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(response)
		return
	}

	createDrierUseCase := admin.InitCreateDrierUseCase(h.dbRepo, h.cacheRepo)

	if error, errorStatus := createDrierUseCase.Execute(plcId, &request); error != nil {

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

	response := &response.StatusMessage{
		Message: "drier created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetDriersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]

	getDriersUseCase := admin.InitGetDriersUseCase(h.dbRepo)

	error, errorStatus, driers := getDriersUseCase.Execute(plcId)

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

func (h *Handler) DeleteDrierHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]
	drierId := vars["drierId"]

	deleteDrierUseCase := admin.InitDeleteDrierUseCase(h.dbRepo, h.cacheRepo)

	error, errorStatus := deleteDrierUseCase.Execute(plcId, drierId)

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

	response := &response.StatusMessage{
		Message: "drier deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreateRegisterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]
	drierId := vars["drierId"]

	var request request.Register

	if error := json.NewDecoder(r.Body).Decode(&request); error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	createRegisterUseCase := admin.InitCreateRegisterUseCase(h.dbRepo, h.cacheRepo)

	if error, errorStatus := createRegisterUseCase.Execute(plcId, drierId, &request); error != nil {
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

	response := &response.StatusMessage{
		Message: "register created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetRegistersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]
	drierId := vars["drierId"]

	getRegistersUseCase := admin.InitGetRegisterUseCase(h.dbRepo)

	error, errorStatus, registers := getRegistersUseCase.Execute(plcId, drierId)

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

	response := &response.Registers{
		Registers: registers,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteRegisterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]

	drierId := vars["drierId"]

	regAddress := vars["regAddress"]

	regTypeName := vars["regTypeName"]

	deleteRegisterUseCase := admin.InitDeleteRegisterUseCase(h.dbRepo, h.cacheRepo)

	error, errorStatus := deleteRegisterUseCase.Execute(plcId, drierId, regAddress, regTypeName)

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

	response := &response.StatusMessage{
		Message: "register deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) CreateRegisterTypeHandler(w http.ResponseWriter, r *http.Request) {
	var request request.RegisterType

	if error := json.NewDecoder(r.Body).Decode(&request); error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	createRegTypeUseCase := admin.InitCreateRegTypeUseCase(h.dbRepo)

	if error, errorStatus := createRegTypeUseCase.Execute(&request); error != nil {
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

	response := &response.StatusMessage{
		Message: "register type created successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) DeleteRegTypeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	regTypeName := vars["regTypeName"]

	deleteRegTypeUseCase := admin.InitDeleteRegTypeUseCase(h.dbRepo)

	if error, errorStatus := deleteRegTypeUseCase.Execute(regTypeName); error != nil {
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

	response := &response.StatusMessage{
		Message: "register type deleted successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GiveUserAccessHandler(w http.ResponseWriter, r *http.Request) {
	var request request.UserAccess

	if error := json.NewDecoder(r.Body).Decode(&request); error != nil {
		response := &response.StatusMessage{
			Message: "invalid json format",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	giveUserAccessUseCase := admin.InitGiveUserAccessUseCase(h.dbRepo, h.smtpRepo)

	if error, errorStatus := giveUserAccessUseCase.Execute(&request); error != nil {
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

	response := &response.StatusMessage{
		Message: "email sent successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetRegisterTypesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	plcId := vars["plcId"]
	drierId := vars["drierId"]

	getRegisterTypesUseCase := admin.InitGetRegisterTypesUseCase(h.dbRepo)

	error, errorStatus, regTypes := getRegisterTypesUseCase.Execute(plcId, drierId)

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

	response := &response.RegTypes{
		RegTypes: regTypes,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

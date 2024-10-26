package route

import (
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/handler/admin"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/handler/user"
	"github.com/gorilla/mux"
)

func Router(dbRepo repository.DataBaseRepository, smtpRepo repository.SmtpClientRepository) *mux.Router {

	adminHandler := admin.NewHandler(dbRepo, smtpRepo)
	userHandler := user.NewHandler(dbRepo)

	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	rootRouter := router.PathPrefix("/root").Subrouter()
	adminRouter := router.PathPrefix("/admin").Subrouter()
	userRouter := router.PathPrefix("/user").Subrouter()

	// adminRouter.Use(middleware.AdminAuthenticationMiddleWare)

	loginRouter.HandleFunc("/admin", adminHandler.AdminLoginHandler).Methods("POST")
	loginRouter.HandleFunc("/user", userHandler.UserLoginHandler).Methods("POST")

	rootRouter.HandleFunc("/create/admin", adminHandler.SignUpHandler).Methods("POST")

	adminRouter.HandleFunc("/database/init", adminHandler.DatabaseInitializeHandler).Methods("GET")
	adminRouter.HandleFunc("/delete/admin/{adminId}", adminHandler.DeleteAdminHandler).Methods("DELETE")
	adminRouter.HandleFunc("/create/user", adminHandler.CreateUserHandler).Methods("POST")
	adminRouter.HandleFunc("/delete/user/{userId}", adminHandler.DeleteUserHandler).Methods("DELETE")
	adminRouter.HandleFunc("/users", adminHandler.GetAllUsersHandler).Methods("GET")
	adminRouter.HandleFunc("/create/plc/{userId}", adminHandler.CreatePlcHandler).Methods("POST")
	adminRouter.HandleFunc("/delete/plc/{plcId}", adminHandler.DeletePlcHandler).Methods("DELETE")
	adminRouter.HandleFunc("/plcs/{userId}", adminHandler.GetPlcsHandler).Methods("GET")
	adminRouter.HandleFunc("/create/drier/{plcId}", adminHandler.CreateDrierHandler).Methods("POST")
	adminRouter.HandleFunc("/driers/{plcId}", adminHandler.GetDriersHandler).Methods("GET")
	adminRouter.HandleFunc("/delete/drier/{drierId}", adminHandler.DeleteDrierHandler).Methods("DELETE")
	adminRouter.HandleFunc("/create/register/{plcId}/{drierId}", adminHandler.CreateRegisterHandler).Methods("POST")
	adminRouter.HandleFunc("/registers/{plcId}/{drierId}", adminHandler.GetRegistersHandler).Methods("GET")
	adminRouter.HandleFunc("/delete/register/{plcId}/{drierId}/{regAddress}/{regTypeName}", adminHandler.DeleteRegisterHandler).Methods("DELETE")
	adminRouter.HandleFunc("/create/register_type", adminHandler.CreateRegisterTypeHandler).Methods("POST")
	adminRouter.HandleFunc("/delete/register_type/{regTypeName}", adminHandler.DeleteRegTypeHandler).Methods("DELETE")
	adminRouter.HandleFunc("/give/user/access", adminHandler.GiveUserAccessHandler).Methods("POST")
	adminRouter.HandleFunc("/register/types/{plcId}/{drierId}", adminHandler.GetRegisterTypesHandler).Methods("GET")

	userRouter.HandleFunc("/driers/{userId}", userHandler.GetDriersHandler).Methods("GET")
	userRouter.HandleFunc("/recipe/step/count/{drierId}", userHandler.GetRecipeStepCountHandler).Methods("GET")

	return router
}

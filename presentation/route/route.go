package route

import (
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/handler/admin"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/handler/user"
	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/middleware"
	"github.com/gorilla/mux"
)

func Router(dbRepo repository.DataBaseRepository, cacheRepo repository.CacheRepository, smtpRepo repository.SmtpClientRepository) *mux.Router {

	adminHandler := admin.NewHandler(dbRepo, cacheRepo, smtpRepo)
	userHandler := user.NewHandler(dbRepo)

	router := mux.NewRouter()

	loginRouter := router.PathPrefix("/login").Subrouter()
	rootRouter := router.PathPrefix("/root").Subrouter()
	adminRouter := router.PathPrefix("/admin").Subrouter()
	userRouter := router.PathPrefix("/user").Subrouter()

	router.Use(middleware.CorsMiddleWare)
	// adminRouter.Use(middleware.AdminAuthenticationMiddleWare)

	loginRouter.HandleFunc("/admin", adminHandler.AdminLoginHandler).Methods("POST")
	loginRouter.HandleFunc("/user", userHandler.UserLoginHandler).Methods("POST")

	rootRouter.HandleFunc("/create/admin", adminHandler.SignUpHandler).Methods("POST")
	rootRouter.HandleFunc("/delete/admin/{adminId}", adminHandler.DeleteAdminHandler).Methods("DELETE,OPTIONS")

	adminRouter.HandleFunc("/database/init", adminHandler.DatabaseInitializeHandler).Methods("GET")
	adminRouter.HandleFunc("/create/user", adminHandler.CreateUserHandler).Methods("POST")
	adminRouter.HandleFunc("/delete/user/{userId}", adminHandler.DeleteUserHandler).Methods("DELETE,OPTIONS")
	adminRouter.HandleFunc("/users", adminHandler.GetAllUsersHandler).Methods("GET")
	adminRouter.HandleFunc("/create/plc/{userId}", adminHandler.CreatePlcHandler).Methods("POST")
	adminRouter.HandleFunc("/delete/plc/{plcId}", adminHandler.DeletePlcHandler).Methods("DELETE,OPTIONS")
	adminRouter.HandleFunc("/plcs/{userId}", adminHandler.GetPlcsHandler).Methods("GET")
	adminRouter.HandleFunc("/create/drier/{plcId}", adminHandler.CreateDrierHandler).Methods("POST")
	adminRouter.HandleFunc("/driers/{plcId}", adminHandler.GetDriersHandler).Methods("GET")
	adminRouter.HandleFunc("/delete/drier/{plcId}/{drierId}", adminHandler.DeleteDrierHandler).Methods("DELETE,OPTIONS")
	adminRouter.HandleFunc("/create/register/{plcId}/{drierId}", adminHandler.CreateRegisterHandler).Methods("POST")
	adminRouter.HandleFunc("/registers/{plcId}/{drierId}", adminHandler.GetRegistersHandler).Methods("GET")
	adminRouter.HandleFunc("/delete/register/{plcId}/{drierId}/{regAddress}/{regTypeName}", adminHandler.DeleteRegisterHandler).Methods("DELETE,OPTIONS")
	adminRouter.HandleFunc("/create/register_type", adminHandler.CreateRegisterTypeHandler).Methods("POST")
	adminRouter.HandleFunc("/delete/register_type/{regTypeName}", adminHandler.DeleteRegTypeHandler).Methods("DELETE,OPTIONS")
	adminRouter.HandleFunc("/give/user/access", adminHandler.GiveUserAccessHandler).Methods("POST")
	adminRouter.HandleFunc("/register/types/{plcId}/{drierId}", adminHandler.GetRegisterTypesHandler).Methods("GET")

	userRouter.HandleFunc("/driers/{userId}", userHandler.GetDriersHandler).Methods("GET")
	userRouter.HandleFunc("/recipe/step/count/{drierId}", userHandler.GetRecipeStepCountHandler).Methods("GET")
	userRouter.HandleFunc("/drier/statuses/{plcId}/{drierId}", userHandler.GetDrierStatusesHandler).Methods("GET")

	return router
}

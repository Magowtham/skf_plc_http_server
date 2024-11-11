package repository

import "github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"

type DataBaseRepository interface {
	Init() error

	CheckAdminEmailExists(email string) (bool, error)

	CreateAdmin(admin *entity.Admin) error

	DeleteAdmin(adminId string) error

	CheckAdminIdExists(adminId string) (bool, error)

	GetAdminByEmail(email string) (*entity.Admin, error)

	CreateUser(user *entity.User) error

	CheckUserEmailExists(email string) (bool, error)

	CheckUserIdExists(userId string) (bool, error)

	DeleteUser(userId string) error

	GetUserById(userId string) (*entity.User, error)

	GetUserByEmail(email string) (*entity.User, error)

	GetAllUsers() ([]*entity.User, error)

	CheckPlcIdExists(plcId string) (bool, error)

	CreatePlc(plc *entity.Plc) error

	DeletePlc(plcId string) error

	GetPlcsByUserId(userId string) ([]*entity.Plc, error)

	CreateDrier(drier *entity.Drier) error

	GetDriersByUserId(userId string) ([]*entity.Drier, error)

	GetDriersByPlcId(plcId string) ([]*entity.Drier, error)

	CheckDrierIdExists(drierId string) (bool, error)

	DeleteDrier(drierId string) error

	CheckRegisterAddressAndRegisterTypeExists(plcId string, registerAddress string, registerType string) (bool, error)

	CheckRegisterAddressExists(plcId string, registerAddress string) (bool, error)

	CheckRegisterTypeExists(plcId string, drierId string, registerType string) (bool, error)

	UpdateDrierRecipeStepCountAndCreateRegister(plcId string, register *entity.Register) error

	CreateRegister(plcId string, register *entity.Register) error

	GetRegisterAddressesByDrierId(plcId string, drierId string) ([]string, error)

	GetAllRegisterAddress(plcId string) ([]string, error)

	GetRegistersByDrierId(plcId string, drierId string) ([]*entity.Register, error)

	UpdateDrierRecipeStepCountAndDeleteRegisterByRegAddress(plcId string, drierId string, registerAddress string) error

	DeleteRegisterByRegAddress(plcId string, registerAddress string) error

	CheckRegTypeNameExistsInRegTypes(regTypeName string) (bool, error)

	CreateRegType(regType *entity.RegisterType) error

	DeleteRegType(regTypeName string) error

	GetAllRegisterTypes() ([]*entity.RegisterType, error)

	GetRegisterTypesFromPlcByDrierId(plcId string, drierId string) ([]string, error)

	GetRecipeStepCount(drierId string) (int, error)

	GetRegisterValueByRegisterTypeAndDrierId(plcId string, drierId string, regType string) (string, error)

	CreateUserFeedback(userId string, feedback string) error
}

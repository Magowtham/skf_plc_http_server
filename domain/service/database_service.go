package service

import (
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
)

type DataBaseService struct {
	DataBaseRepo repository.DataBaseRepository
}

func NewDataBaseService(repo repository.DataBaseRepository) *DataBaseService {
	return &DataBaseService{
		DataBaseRepo: repo,
	}
}

func (service *DataBaseService) InitializeDataBase() error {
	return service.DataBaseRepo.Init()
}

func (service *DataBaseService) CheckAdminEmailExists(email string) (bool, error) {
	return service.DataBaseRepo.CheckAdminEmailExists(email)
}

func (service *DataBaseService) CheckAdminIdExists(adminId string) (bool, error) {
	return service.DataBaseRepo.CheckAdminIdExists(adminId)
}

func (service *DataBaseService) CreateAdmin(admin entity.Admin) error {
	return service.DataBaseRepo.CreateAdmin(admin)
}

func (service *DataBaseService) DeleteAdmin(adminId string) error {
	return service.DataBaseRepo.DeleteAdmin(adminId)
}

func (service *DataBaseService) GetAdminByEmail(email string) (entity.Admin, error) {
	return service.DataBaseRepo.GetAdminByEmail(email)
}

func (service *DataBaseService) CreateUser(user *entity.User) error {
	return service.DataBaseRepo.CreateUser(user)
}

func (service *DataBaseService) CheckUserIdExists(userId string) (bool, error) {
	return service.DataBaseRepo.CheckUserIdExists(userId)
}

func (service *DataBaseService) CheckUserEmailExists(email string) (bool, error) {
	return service.DataBaseRepo.CheckUserEmailExists(email)
}

func (service *DataBaseService) DeleteUser(userId string) error {
	return service.DataBaseRepo.DeleteUser(userId)
}

func (service *DataBaseService) GetUserById(userId string) (entity.User, error) {
	return service.DataBaseRepo.GetUserById(userId)
}

func (service *DataBaseService) GetUserByEmail(email string) (entity.User, error) {
	return service.DataBaseRepo.GetUserByEmail(email)
}

func (service *DataBaseService) GetAllUsers() ([]entity.User, error) {
	return service.DataBaseRepo.GetAllUsers()
}

func (service *DataBaseService) CheckPlcIdExists(plcId string) (bool, error) {
	return service.DataBaseRepo.CheckPlcIdExists(plcId)
}

func (service *DataBaseService) CreatePlc(plc entity.Plc) error {
	return service.DataBaseRepo.CreatePlc(plc)
}

func (service *DataBaseService) DeletePlc(plcId string) error {
	return service.DataBaseRepo.DeletePlc(plcId)
}

func (service *DataBaseService) GetPlcsByUserId(userId string) ([]entity.Plc, error) {
	return service.DataBaseRepo.GetPlcsByUserId(userId)
}

func (service *DataBaseService) CreateDrier(drier *entity.Drier) error {
	return service.DataBaseRepo.CreateDrier(drier)
}

func (service *DataBaseService) GetDriersByUserId(userId string) ([]entity.Drier, error) {
	return service.DataBaseRepo.GetDriersByUserId(userId)
}

func (service *DataBaseService) GetDriersByPlcId(plcId string) ([]entity.Drier, error) {
	return service.DataBaseRepo.GetDriersByPlcId(plcId)
}

func (service *DataBaseService) GetAllRegisterAddress(plcId string) ([]string, error) {
	return service.DataBaseRepo.GetAllRegisterAddress(plcId)
}
func (service *DataBaseService) CheckDrierIdExists(drierId string) (bool, error) {
	return service.DataBaseRepo.CheckDrierIdExists(drierId)
}

func (service *DataBaseService) DeleteDrier(drierId string) error {
	return service.DataBaseRepo.DeleteDrier(drierId)
}

func (service *DataBaseService) CheckRegisterAddressExists(plcId string, registerAddress string) (bool, error) {
	return service.DataBaseRepo.CheckRegisterAddressExists(plcId, registerAddress)
}

func (service *DataBaseService) CheckRegisterAddressAndRegisterTypeExists(plcId string, registerAddress string, registerType string) (bool, error) {
	return service.DataBaseRepo.CheckRegisterAddressAndRegisterTypeExists(plcId, registerAddress, registerType)
}

func (service *DataBaseService) CheckRegisterTypeExists(plcId string, drierId string, registerType string) (bool, error) {
	return service.DataBaseRepo.CheckRegisterTypeExists(plcId, drierId, registerType)
}

func (service *DataBaseService) UpdateDrierRecipeStepCountAndCreateRegister(plcId string, register *entity.Register) error {
	return service.DataBaseRepo.UpdateDrierRecipeStepCountAndCreateRegister(plcId, register)
}

func (service *DataBaseService) CreateRegister(plcId string, register *entity.Register) error {
	return service.DataBaseRepo.CreateRegister(plcId, register)
}

func (service *DataBaseService) GetRegisterAddressesByDrierId(plcId string, drierId string) ([]string, error) {
	return service.DataBaseRepo.GetRegisterAddressesByDrierId(plcId, drierId)
}

func (service *DataBaseService) GetRegistersByDrierId(plcId string, drierId string) ([]entity.Register, error) {
	return service.DataBaseRepo.GetRegistersByDrierId(plcId, drierId)
}

func (service *DataBaseService) UpdateDrierRecipeStepCountAndDeleteRegisterByRegAddress(plcId string, drierId string, registerAddress string) error {
	return service.DataBaseRepo.UpdateDrierRecipeStepCountAndDeleteRegisterByRegAddress(plcId, drierId, registerAddress)
}

func (service *DataBaseService) DeleteRegisterByRegAddress(plcId string, registerAddress string) error {
	return service.DataBaseRepo.DeleteRegisterByRegAddress(plcId, registerAddress)
}

func (service *DataBaseService) CheckRegTypeNameExistsInRegTypes(regTypeName string) (bool, error) {
	return service.DataBaseRepo.CheckRegTypeNameExistsInRegTypes(regTypeName)
}

func (service *DataBaseService) CreateRegType(regType *entity.RegisterType) error {
	return service.DataBaseRepo.CreateRegType(regType)
}

func (service *DataBaseService) DeleteRegType(regTypeName string) error {
	return service.DataBaseRepo.DeleteRegType(regTypeName)
}

func (service *DataBaseService) GetAllRegisterTypes() ([]entity.RegisterType, error) {
	return service.DataBaseRepo.GetAllRegisterTypes()
}

func (service *DataBaseService) GetRegisterTypesFromPlcByDrierId(plcId string, drierId string) ([]string, error) {
	return service.DataBaseRepo.GetRegisterTypesFromPlcByDrierId(plcId, drierId)
}

func (service *DataBaseService) GetRecipeStepCount(drierId string) (int, error) {
	return service.DataBaseRepo.GetRecipeStepCount(drierId)
}

func (service *DataBaseService) GetRegisterValueByRegisterTypeAndDrierId(plcId string, drierId string, regType string) (string, error) {
	return service.DataBaseRepo.GetRegisterValueByRegisterTypeAndDrierId(plcId, drierId, regType)
}

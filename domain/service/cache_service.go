package service

import (
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/VsenseTechnologies/skf_plc_http_server/domain/repository"
)

type CacheService struct {
	repo repository.CacheRepository
}

func NewCacheService(repo repository.CacheRepository) *CacheService {
	return &CacheService{
		repo,
	}
}

func (service *CacheService) CreateDrier(drierId string) error {
	return service.repo.CreateDrier(drierId)
}

func (service *CacheService) DeleteDrier(drierId string) error {
	return service.repo.DeleteDrier(drierId)
}

func (service *CacheService) CreateRegister(plcId string, reg *entity.Register) error {
	return service.repo.CreateRegister(plcId, reg)
}

func (service *CacheService) DeleteRegister(plcId string, regAddress string) error {
	return service.repo.DeleteRegister(plcId, regAddress)
}

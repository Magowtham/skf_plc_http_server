package repository

import "github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"

type CacheRepository interface {
	CreateDrier(drierId string) error
	DeleteDrier(drierId string) error
	CreateRegister(plcId string, reg *entity.Register) error
	DeleteRegister(plcId string, regAddress string) error
}

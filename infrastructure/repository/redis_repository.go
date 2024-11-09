package repository

import (
	"context"
	"fmt"

	"github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client,
	}
}

func (repo *RedisRepository) CreateDrier(drierId string) error {
	key1 := fmt.Sprintf("rcp_stp_ct_%s", drierId)
	key2 := fmt.Sprintf("rcp_stp_tp_%s", drierId)
	key3 := fmt.Sprintf("pid_%s", drierId)
	_, err := repo.client.MSet(context.Background(), key1, "0", key2, "0", key3, "0").Result()
	return err
}

func (repo *RedisRepository) DeleteDrier(drierId string) error {
	key1 := fmt.Sprintf("rcp_stp_ct_%s", drierId)
	key2 := fmt.Sprintf("rcp_stp_tp_%s", drierId)
	_, err := repo.client.Del(context.Background(), key1, key2).Result()
	return err
}

func (repo *RedisRepository) CreateRegister(plcId string, reg *entity.Register) error {
	key1 := fmt.Sprintf("dr_id_%s_%s", plcId, reg.RegAddress)
	key2 := fmt.Sprintf("rg_ty_%s_%s", plcId, reg.RegAddress)
	key3 := fmt.Sprintf("rg_vl_%s_%s", plcId, reg.RegAddress)

	_, err := repo.client.MSet(context.Background(), key1, reg.DrierId, key2, reg.RegType, key3, reg.Value).Result()

	return err
}

func (repo *RedisRepository) DeleteRegister(plcId string, regAddress string) error {
	key1 := fmt.Sprintf("dr_id_%s_%s", plcId, regAddress)
	key2 := fmt.Sprintf("rg_ty_%s_%s", plcId, regAddress)
	key3 := fmt.Sprintf("rg_vl_%s_%s", plcId, regAddress)
	_, err := repo.client.Del(context.Background(), key1, key2, key3).Result()
	return err
}

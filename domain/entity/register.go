package entity

import "time"

type Register struct {
	RegAddress          string    `json:"reg_address"`
	RegType             string    `json:"reg_type"`
	Label               string    `json:"label"`
	DrierId             string    `json:"drier_id"`
	Value               string    `json:"value"`
	LastUpdateTimestamp time.Time `json:"last_update_timestamp"`
}

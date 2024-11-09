package response

import "github.com/VsenseTechnologies/skf_plc_http_server/domain/entity"

type StatusMessage struct {
	Message string `json:"message"`
}

type AllUsers struct {
	Users []entity.User `json:"users"`
}

type Plcs struct {
	Plcs []entity.Plc `json:"plcs"`
}

type Driers struct {
	Driers []entity.Drier `json:"driers"`
}

type Registers struct {
	Registers []entity.Register `json:"registers"`
}

type Token struct {
	Token string `json:"token"`
}

type RegTypes struct {
	RegTypes []entity.RegisterType `json:"reg_types"`
}

type RecipeStepCount struct {
	RecipeStepCount int `json:"recipe_step_count"`
}

type DrierStatuses struct {
	DrierStatuses []*entity.DrierStatus `json:"statuses"`
}

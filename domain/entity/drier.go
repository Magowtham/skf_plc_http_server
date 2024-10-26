package entity

type Drier struct {
	DrierId         string `json:"drier_id"`
	PlcId           string `json:"plc_id"`
	RecipeStepCount string `json:"recipe_step_count"`
	Label           string `json:"label"`
}

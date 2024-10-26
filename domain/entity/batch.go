package entity

type Batch struct {
	DrierId    string `json:"drier_id"`
	RecipeStep string `json:"recipe_step"`
	Time       string `json:"time"`
	Temp       string `json:"temp"`
	Pid        string `json:"pid"`
}

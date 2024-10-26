package entity

type Admin struct {
	AdminId  string `json:"admin_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

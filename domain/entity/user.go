package entity

type User struct {
	UserId   string `json:"user_id"`
	Email    string `json:"email"`
	Label    string `json:"label"`
	Password string `json:"-"`
}

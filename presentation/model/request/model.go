package request

type Admin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Label    string `json:"label"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Plc struct {
	PlcId string `json:"plc_id"`
	Label string `json:"label"`
}

type Drier struct {
	Label string `json:"label"`
}

type Register struct {
	RegAddress string `json:"reg_address"`
	RegType    string `json:"reg_type"`
	Label      string `json:"label"`
}

type RegisterType struct {
	Type  string `json:"type"`
	Label string `json:"label"`
}

type UserAccess struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserFeedback struct {
	Feedback string `json:"feedback"`
}

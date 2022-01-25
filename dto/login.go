package dto

type LoginForm struct {
	Pwd string `json:"pwd" form:"pwd" example:"密码"`
}

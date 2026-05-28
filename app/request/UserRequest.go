package request

type UserLoginRequest struct {
	UserName string `json:"username" form:"user_name" validate:"required,max=255"`
	Password string `json:"password" form:"password" validate:"required"`
}

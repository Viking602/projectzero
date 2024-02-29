package types

type UserRegister struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

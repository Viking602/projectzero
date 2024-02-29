package types

type UserRegister struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	UserName string `json:"username"`
	Token    string `json:"token"`
}

type UserInfoRequest struct {
	UserName string `form:"username" binding:"required"`
}

type UserInfoResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	UserType int    `json:"userType"`
	Status   int    `json:"status"`
	DeleteAt int64  `json:"deleteAt"`
	CreateAt int64  `json:"createAt"`
	UpdateAt int64  `json:"updateAt"`
}

type UserUpdateRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	UserType int    `json:"userType"`
	Status   int    `json:"status"`
}

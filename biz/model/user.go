package model

type RegisterRequest struct {
	NickName string `json:"nick_name" form:"nick_name"`

	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"password" form:"password"`

	Key string `json:"key" form:"key"` // 现阶段前端验证。这个 key 作为对称密钥，可用来加密金钱
}

type RegisterResponse struct {
}

type LoginRequest struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"password" form:"password"`
}

type LoginResponse struct {
	Token    string        `json:"token"`
	UserInfo BasicUserInfo `json:"user_info"`
}

type BasicUserInfo struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

type UpdateUserRequest struct {
	NickName string `json:"nick_name"`
}

type UpdateUserResponse struct {
}

type UploadAvatarRequest struct {
	Avatar string `json:"avatar"`
}

type UploadAvatarResponse struct {
	UserInfo BasicUserInfo `json:"user_info"`
}

type SendEmailRequest struct {
	Email         string        `json:"email"`
	Password      string        `json:"password"`
	OperationType SendEmailType `json:"operation_type"`
}

type ValidateEmailRequest struct {
	EmailToken string `form:"email_token" json:"email_token"`
}

type ValidateEmailResponse struct {
}

type ShowMoneyRequest struct {
	Key string `form:"key" json:"key"`
}

type ShowMoneyResponse struct {
	UserId   uint   `form:"user_id" json:"user_id"`
	UserName string `form:"user_name" json:"user_name"`
	Money    string `form:"money" json:"money"`
}

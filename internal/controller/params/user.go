package params

import "github.com/FIY-pc/user-manager/internal/model"

type LoginReq struct {
	Email    string `json:"Email" form:"email" validate:"email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RegisterReq struct {
	Email    string `json:"Email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Nickname string `json:"Name" form:"nickname"`
}

type Login200Resp struct {
	Code string       `json:"code"`
	Msg  string       `json:"msg"`
	Data Login200Data `json:"data"`
}

type Login200Data struct {
	Token      string `json:"token"`
	Permission int    `json:"permission"`
}

type Register200Resp struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	Nickname string `json:"nickname"`
}

type UserInfoResp struct {
	Code string     `json:"code"`
	Data model.User `json:"data"`
}

package controller

import (
	"github.com/FIY-pc/user-manager/internal/config"
	"github.com/FIY-pc/user-manager/internal/controller/params"
	"github.com/FIY-pc/user-manager/internal/model"
	"github.com/FIY-pc/user-manager/internal/tools"
	"github.com/FIY-pc/user-manager/internal/util"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func Login(c echo.Context) error {
	req := params.LoginReq{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, params.CommonErrorResp{
			Code: "400",
			Msg:  "invalid params",
		})
	}

	// 检查用户是否存在
	user, err := model.GetUser(req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, params.CommonErrorResp{
			Code: "400",
			Msg:  "user not found",
		})
	}

	// 检查密码是否正确
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return c.JSON(http.StatusBadRequest, params.CommonErrorResp{
			Code: "400",
			Msg:  "password is incorrect",
		})
	}

	// 生成token
	claims := util.JwtClaims{
		ID:         user.ID,
		Permission: user.Permission,
		Exp:        time.Now().Unix() + config.Config.Jwt.Exp,
	}
	token, err := util.GenerateToken(claims)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "generate token failed",
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, params.Login200Resp{
		Code: "200",
		Msg:  "login success",
		Data: params.Login200Data{
			Token:      token,
			Permission: user.Permission,
		},
	})
}

func Register(c echo.Context) error {
	req := params.RegisterReq{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
		Nickname: c.FormValue("nickname"),
	}
	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, params.CommonErrorResp{
			Code:  "400",
			Msg:   "invalid params",
			Error: "email or password is empty",
		})
	}

	// 检查用户是否存在
	_, err := model.GetUser(req.Email)
	if err == nil {
		return c.JSON(http.StatusBadRequest, params.CommonErrorResp{
			Code:  "400",
			Msg:   "user is already exists",
			Error: "",
		})
	}
	// 创建用户
	user := model.User{
		Email:      req.Email,
		Password:   req.Password,
		Nickname:   req.Nickname,
		Permission: model.PermissionUser,
	}
	if user.Nickname == "" {
		user.Nickname = tools.GenerateRandName()
	}
	// 对密码进行哈希加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "hash password failed",
			Error: err.Error(),
		})
	}
	user.Password = string(hashPassword)
	_, err = model.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "create user failed",
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, params.Register200Resp{
		Code:     "200",
		Msg:      "register success",
		Nickname: user.Nickname,
	})
}

func GetUser(c echo.Context) error {
	user, err := model.GetUser(c.QueryParam("email"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "get user failed",
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, params.UserInfoResp{
		Code: "200",
		Data: user,
	})
}

func CreateUser(c echo.Context) error {
	var err error
	user := model.User{}
	user.Email = c.QueryParam("email")
	user.Nickname = c.QueryParam("nickname")
	user.Password = c.QueryParam("password")
	user.Permission, err = strconv.Atoi(c.QueryParam("permission"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, params.CommonErrorResp{
			Code:  "400",
			Msg:   "invalid params",
			Error: err.Error(),
		})
	}
	if user.Nickname == "" {
		user.Nickname = tools.GenerateRandName()
	}
	// 对密码进行哈希加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "hash password failed",
			Error: err.Error(),
		})
	}
	user.Password = string(hashPassword)
	user, err = model.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "create user failed",
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, params.UserInfoResp{
		Code: "200",
		Data: user,
	})
}

func UpdateUser(c echo.Context) error {
	user, err := model.GetUser(c.QueryParam("email"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "get user failed",
			Error: err.Error(),
		})
	}
	if c.QueryParam("nickname") != "" {
		user.Nickname = c.QueryParam("nickname")
	}
	if c.QueryParam("password") != "" {
		user.Password = c.QueryParam("password")
	}
	if c.QueryParam("permission") != "" {
		user.Permission, err = strconv.Atoi(c.QueryParam("permission"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, params.CommonErrorResp{
				Code:  "400",
				Msg:   "invalid params",
				Error: err.Error(),
			})
		}
	}
	// 对密码进行哈希加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "hash password failed",
			Error: err.Error(),
		})
	}
	user.Password = string(hashPassword)
	user, err = model.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "update user failed",
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, params.UserInfoResp{
		Code: "200",
		Data: user,
	})
}

func DeleteUser(c echo.Context) error {
	user, err := model.DeleteUser(c.QueryParam("email"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, params.CommonErrorResp{
			Code:  "500",
			Msg:   "delete user failed",
			Error: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, params.UserInfoResp{
		Code: "200",
		Data: user,
	})
}

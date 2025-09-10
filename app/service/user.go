package service

import (
	"NewProject/app/biz"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/wapper"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserBiz *biz.UserBiz
}

func NewUserService(userBiz *biz.UserBiz) *UserService {
	return &UserService{
		UserBiz: userBiz,
	}
}

// 注册

func (s *UserService) Register(c *gin.Context) {
	var (
		registerReq scheme.UserRegisterReq
		err         error
		newUser     *models.User
		errCode     wapper.ErrorCode
	)
	err = c.ShouldBindJSON(&registerReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}

	newUser, errCode = s.UserBiz.Register(registerReq)
	if errCode != wapper.Success {
		wapper.ResError(c, errCode)
		return
	}
	wapper.ResSuccess(c, newUser)
}

// 登录

func (s *UserService) Login(c *gin.Context) {
	var loginReq scheme.UserLoginReq
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}

	var errCode wapper.ErrorCode
	loginRes, errCode := s.UserBiz.Login(loginReq)
	if errCode != wapper.Success {
		wapper.ResError(c, errCode)
		return
	}
	wapper.ResSuccess(c, loginRes)
}

// 获取用户列表

func (s *UserService) GetUserList(c *gin.Context) {
	var (
		userListReq  scheme.UserListReq
		err          error
		userListResp scheme.UserListResp
		errCode      wapper.ErrorCode
	)
	err = c.ShouldBindJSON(&userListReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	userListResp, errCode = s.UserBiz.GetUserList(userListReq)
	if errCode != wapper.Success {
		wapper.ResError(c, errCode)
		return
	}
	wapper.ResSuccess(c, userListResp)
}

// 用户查询
func (s *UserService) GetUserInfo(c *gin.Context) {
	var getUserInfoReq scheme.GetUserInfoReq
	if err := c.ShouldBindJSON(&getUserInfoReq); err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	var errCode wapper.ErrorCode
	userInfo, errCode := s.UserBiz.GetUserInfo(getUserInfoReq)
	if errCode != wapper.Success {
		wapper.ResError(c, errCode)
		return
	}
	wapper.ResSuccess(c, userInfo)

}

// 更新用户信息--修改用户名，密码和邮箱和修改状态（传了值的修改，没传的不改，修改完密码要重新加密）

func (s *UserService) UpdateUser(c *gin.Context) {
	var (
		err           error
		updateUserReq scheme.UserUpdateReq
		newUpdateInfo models.User
		errCode       wapper.ErrorCode
	)
	err = c.ShouldBindJSON(&updateUserReq)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	newUpdateInfo, errCode = s.UserBiz.UpdateUser(updateUserReq)
	if errCode != wapper.Success {
		wapper.ResError(c, errCode)
		return
	}
	wapper.ResSuccess(c, newUpdateInfo)
}

// 删除用户
func (s *UserService) DelUser(c *gin.Context) {
	var (
		userId scheme.DelUserReq
		err    error
	)
	err = c.ShouldBindJSON(&userId)
	if err != nil {
		wapper.ResError(c, wapper.ParameterBindingFailed)
		return
	}
	errCode := s.UserBiz.DelUser(userId)
	if errCode != wapper.Success {
		wapper.ResError(c, errCode)
		return
	}
	wapper.ResSuccess(c, "删除成功")
}

//增加了查询用户拥有的资源列表功能，并把路径信息单独摘出来和查询到的资源信息一起输出；

package biz

import (
	"NewProject/app/data"
	"NewProject/app/scheme"
	"NewProject/models"
	"NewProject/pkg/util"
	"NewProject/pkg/wapper"
	"errors"
	"gorm.io/gorm"
)

type UserBiz struct {
	UserData *data.UserData
}

func NewUserBiz(userData *data.UserData) *UserBiz {
	return &UserBiz{
		UserData: userData,
	}
}

func (b *UserBiz) Register(registerReq scheme.UserRegisterReq) (*models.User, wapper.ErrorCode) {
	//把要用的变量定义在一起
	var (
		err      error
		userData = models.User{
			Email:        registerReq.Email,
			Username:     registerReq.Username,
			PasswordHash: registerReq.PasswordHash,
			IsActive:     scheme.UserActive,
		}
		userInfo       models.User
		hashedPassword string
		registerInfo   models.User
	)

	//验证邮箱注册格式是否正确(还没实现)
	//验证用户名是否存在
	userInfo, err = b.UserData.GetInfoByUsername(registerReq.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, wapper.DataNotFound
	}
	if userInfo.Username == registerReq.Username {
		return nil, wapper.UsernameExisted
	}
	//密码加密
	hashedPassword, err = util.HashPassword(registerReq.PasswordHash)
	if err != nil {
		return nil, wapper.PasswordEncryptionFailed
	}
	userData.PasswordHash = hashedPassword

	registerInfo, err = b.UserData.Register(&userData)
	if err != nil {
		return nil, wapper.RegisterFailed
	}
	return &registerInfo, wapper.Success
}

// 登录
func (b *UserBiz) Login(loginReq scheme.UserLoginReq) (models.User, wapper.ErrorCode) {
	var (
		err       error
		userInfo  models.User
		loginData = models.User{
			PasswordHash: loginReq.PasswordHash,
		}
	)
	if loginReq.LoginName != "" {
		userInfo, err = b.UserData.GetInfoByEmailOrUsername(loginReq.LoginName)
		if err != nil {
			return models.User{}, wapper.DataNotFound
		}
	} else {
		return models.User{}, wapper.UsernameOrEmailIsNull
	}
	if !util.CheckPassword(userInfo.PasswordHash, loginData.PasswordHash) {
		return models.User{}, wapper.PasswordError
	}
	//检查用户是否存在
	if userInfo.Username == "" || userInfo.Email == "" {
		return models.User{}, wapper.UserNotFound
	}

	if userInfo.IsActive == 0 {
		return models.User{}, wapper.NotBeenActivated
	}
	err = b.UserData.Login(loginReq.LoginName)
	if err != nil {
		return models.User{}, wapper.LoginFailed
	}
	return userInfo, wapper.Success
}

// 获取用户列表
func (b *UserBiz) GetUserList(userListReq scheme.UserListReq) (scheme.UserListResp, wapper.ErrorCode) {
	userListResp, err := b.UserData.GetUserList(userListReq)
	if err != nil {
		return scheme.UserListResp{}, wapper.GetUserListFailed
	}

	return userListResp, wapper.Success
}

// 用户查询
func (b *UserBiz) GetUserInfo(userId scheme.GetUserInfoReq) (models.User, wapper.ErrorCode) {
	userInfo, err := b.UserData.GetUserInfo(userId.Id)
	if err != nil {
		return models.User{}, wapper.GetUserInfoFailed
	}
	return userInfo, wapper.Success
}

// 更新用户信息--修改用户名，密码和邮箱和修改状态（传了值的修改，没传的不改，修改完密码要重新加密）
// 只判断username或者email唯一字段是否重复
// 重复=自己的
// 不重复=不是自己的--不可修改
// null=可以修改
// 其余字段为空则不会被修改（gorm特性）

func (b *UserBiz) UpdateUser(updateReq scheme.UserUpdateReq) (models.User, wapper.ErrorCode) {
	var (
		err         error
		updatedData models.User
		//先把请求结构体的值，赋值给返回结构体，并声名成models.user
		userUpdateResp     models.User
		hashedPasswordHash string
		byEmailInfo        models.User
		byUsernameInfo     models.User
	)
	//判断username是否为空，若不为空， 判断是否重复
	if updateReq.Username != "" {
		byUsernameInfo, err = b.UserData.GetInfoByUsername(updateReq.Username)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, wapper.DataNotFound
		}
		if byUsernameInfo.Id > 0 && byUsernameInfo.Id != updateReq.Id {
			return models.User{}, wapper.UsernameExisted
		}
	}
	//判断email是否为空，若不为空， 判断是否重复
	if updateReq.Email != "" {
		byEmailInfo, err = b.UserData.GetInfoByEmail(updateReq.Email)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, wapper.DataNotFound
		}
		if byEmailInfo.Id > 0 && byEmailInfo.Id != updateReq.Id {
			return models.User{}, wapper.EmailExisted
		}
	}
	//若要修改密码（密码不为空），将密码重新加密
	if updateReq.PasswordHash != "" {
		hashedPasswordHash, err = util.HashPassword(updateReq.PasswordHash)
		if err != nil {
			return models.User{}, wapper.PasswordEncryptionFailed
		}
	}
	//把要修改值的所有值赋给models.user
	updatedData = models.User{
		Id:           updateReq.Id,
		Username:     updateReq.Username,
		Email:        updateReq.Email,
		PasswordHash: hashedPasswordHash,
		IsActive:     updateReq.IsActive,
	}
	//把修改过后的model当做参数传给UpdateUser，然后调用数据库修改用户信息，并返回修改后的用户信息
	userUpdateResp, err = b.UserData.UpdateUser(updatedData)
	if err != nil {
		return models.User{}, wapper.DataNotFound
	}
	return userUpdateResp, wapper.Success
}

// 删除用户
func (b *UserBiz) DelUser(userId scheme.DelUserReq) wapper.ErrorCode {
	err := b.UserData.DelUser(userId)
	if err != nil {
		return wapper.DelUserFailed
	}
	return wapper.Success
}

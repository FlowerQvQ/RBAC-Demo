package data

import (
	"NewProject/app/scheme"
	"NewProject/models"
)

type UserData struct {
	DB *Data
}

func NewUserData(data *Data) *UserData {
	return &UserData{
		DB: data,
	}
}

// 用户注册

func (d *UserData) Register(user *models.User) (models.User, error) {
	var (
		err error
	)
	err = d.DB.DBClient.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

// 用户登录

func (d *UserData) Login(emailOrUsername string) error {
	var user = models.User{}

	err := d.DB.DBClient.Model(&models.User{}).Where("username = ? OR email = ?", emailOrUsername, emailOrUsername).First(&user).Error
	if err != nil {
		return err
	}
	return err
}

// 用户列表
func (d *UserData) GetUserList(userListReq scheme.UserListReq) (scheme.UserListResp, error) {
	var (
		err          error
		userList     []models.User
		userListInfo = d.DB.DBClient.Model(&models.User{})
		offset       int   //计算起始位置
		totalRecords int64 //总记录数
		total        int
	)
	//用不同的字段来分别查询数据
	if userListReq.QueryName != "" {
		userListInfo = userListInfo.Where("email LIKE ? OR username LIKE ? ", "%"+userListReq.QueryName+"%", "%"+userListReq.QueryName+"%")
	}
	if userListReq.IsActive != 0 {
		userListInfo = userListInfo.Where("is_active = ?", userListReq.IsActive)
	}
	//获取满足条件的总记录数，计算满足条件的记录数并保存
	err = userListInfo.Count(&totalRecords).Error
	if err != nil {
		return scheme.UserListResp{}, err
	}
	//计算总页数 / 总记录数
	if totalRecords > 0 {
		total = int((totalRecords + int64(userListReq.Limit) - 1) / int64(userListReq.Limit))
	}

	//分页处理
	if userListReq.Page > 0 && userListReq.Limit > 0 {
		//为什么要用 (Page - 1) * Limit
		//因为大多数系统中，页码是从 1 开始计数的，而数据库的偏移量是从 0 开始的。所以需要减 1 来对齐这两个体系。
		offset = (userListReq.Page - 1) * userListReq.Limit //先跳过前10条记录（即第一页），然后取出接下来的10条（即第二页）
		userListInfo = userListInfo.Offset(offset).Limit(userListReq.Limit)
	}

	err = userListInfo.Find(&userList).Error
	if err != nil {
		return scheme.UserListResp{}, err
	}
	userListResp := scheme.UserListResp{
		UserList: userList,
		Total:    total,
		Pagination: scheme.Pagination{
			Limit: userListReq.Limit,
			Page:  userListReq.Page,
		},
	}
	return userListResp, nil
}

// 用户查询
func (d *UserData) GetUserInfo(userId int64) (models.User, error) {
	var user models.User
	err := d.DB.DBClient.Model(&models.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 更新用户信息--修改用户名，密码和邮箱和修改状态（传了值的修改，没传的不改，修改完密码要重新加密）

func (d *UserData) UpdateUser(updateUserInfo models.User) (models.User, error) {
	var (
		err error
	)
	err = d.DB.DBClient.Model(&models.User{}).Where("id = ? ", updateUserInfo.Id).Updates(&updateUserInfo).Error
	if err != nil {
		return models.User{}, err
	}
	//先更新，再查询输出更新好的全部字段
	err = d.DB.DBClient.Model(&models.User{}).Where("id = ?", updateUserInfo.Id).First(&updateUserInfo).Error
	return updateUserInfo, nil
}

// 通过用户名获取数据库信息

func (d *UserData) GetInfoByUsername(username string) (models.User, error) {
	var (
		err  error
		user models.User
	)
	err = d.DB.DBClient.Where("username = ?", username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 通过邮箱获取数据库信息

func (d *UserData) GetInfoByEmail(email string) (models.User, error) {
	var (
		err  error
		user models.User
	)
	err = d.DB.DBClient.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 通过邮箱/用户名获取数据库信息
func (d *UserData) GetInfoByEmailOrUsername(emailOrUsername string) (models.User, error) {
	var user = models.User{}
	err := d.DB.DBClient.Where("username = ? OR email = ? ", emailOrUsername, emailOrUsername).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 删除用户
func (d *UserData) DelUser(userId scheme.DelUserReq) error {

	err := d.DB.DBClient.Model(&models.User{}).Where("id = ?", userId.Id).Update("is_active", scheme.UserInActive).Error
	if err != nil {
		return err
	}
	return nil
}

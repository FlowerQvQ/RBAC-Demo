package util

import "golang.org/x/crypto/bcrypt"

// 密码加密
// bcrypt.GenerateFromPassword 内部会： 自动生成一个随机的 salt（通常为 16 字节）
// 使用该 salt 和给定的 cost 参数（默认是 bcrypt.DefaultCost，通常是 10）对密码进行多次加密计算
// 返回一个包含 salt 和哈希值的字符串
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

const (
	PwdMinLength = 8
	PwdMaxLength = 16
)

// 密码长度校验
func CheckPasswordLen(password string) bool {
	pwdLen := len(password)
	return pwdLen < PwdMinLength || pwdLen > PwdMaxLength
}

// 验证密码
//输入明文密码和数据库中保存的哈希值，
//CompareHashAndPassword 会从哈希字符串中自动提取出 salt 和 cost 参数
//然后使用相同的 cost 对输入的明文密码加 salt 进行哈希运算
//比较两个哈希是否一致

func CheckPassword(hash, password string) bool {
	//password 用户输入的明文密码
	//hash 数据库中存储的 bcrypt 加密后的哈希字符串
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package model

import (
	"github.com/ShiinaOrez/ginny/util"
)

// 检查数据库中是否有相同用户名的用户
func CheckUserByUsername(username string) bool {
	userID := 0
	err := DB.Self.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	return err == nil
}

// 检查用户名密码是否匹配
func CheckPasswordValidate(username, password string) bool {
	var passwordHash, salt string
	err := DB.Self.QueryRow("SELECT password_hash, salt FROM users WHERE username = ?", username).Scan(&passwordHash, &salt)
	return util.CheckPassword(password, passwordHash, salt) && err == nil
}

// 新建一个User
func CreateUser(username, passwordHash, salt string) {
	DB.Self.Query("INSERT INTO users (username, password_hash, salt) VALUES (?, ?, ?)", username, passwordHash, salt)
}

// 更新密码字段
func UpdatePassword(username, password string) error {
	newPasswordHash, newSalt, err := util.PasswordHash(password)
	if err != nil {
		return err
	}
	_, err = DB.Self.Query("UPDATE users SET password_hash = ?, salt = ? WHERE username = ?", newPasswordHash, newSalt, username)
	return err
}

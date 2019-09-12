package model

// 检查数据库中是否有相同用户名的用户
func CheckUserByUsername(username string) bool {
    userID := 0
    err := DB.Self.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
    return err == nil
}

// 检查用户名密码是否匹配
func CheckPasswordValidate(username, password string) bool {
    var record string
	err := DB.Self.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&record)
	return password == record && err == nil
}

func CreateUser(username, password string) {
    DB.Self.Query("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
}

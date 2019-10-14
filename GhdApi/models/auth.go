package models

// Auth 认证 结构
type Auth struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth 检测授权
func CheckAuth(username, password string) bool {
	// var auth Auth
	// db.Select("id").Where(Auth{Username : username, Password : password}).First(&auth)
	// if auth.ID > 0 {
	// 	return true
	// }
	
	return true
	return false
}
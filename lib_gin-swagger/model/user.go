package model

type User struct {
	Username string `form:"username" json:"username" binding:"required_default"`
	Password string `form:"password" json:"password" binding:"required_default"`
}

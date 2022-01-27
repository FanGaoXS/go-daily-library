package controller

import (
	"github.com/gin-gonic/gin"
	"lib_gin_swagger/model"
	"lib_gin_swagger/vo/response"
	"net/http"
)

// @BasePath /api/v1

// UserLogin
// user 		login godoc
// @Summary 	user login
// @Description user login
// @Tags 		user
// @Param		user	body	model.User		true	"用户信息"
// @Success 	200 		{object} model.User			"用户信息"
// @Router 		/user/login [POST]
func UserLogin(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.New().SetMessage("user login").SetData(user))
}

// UserLogout
// user 		logout godoc
// @Summary 	user logout
// @Description user logout
// @Tags 		user
// @Param		username	query	 	string		true	"用户名"
// @Success 	200 		{boolean} 	bool		"退出是否成功"
// @Router 		/user/logout [GET]
func UserLogout(c *gin.Context) {
	username := c.Query("username")
	flag := false
	if username == "test" {
		flag = true
	}
	c.JSON(http.StatusOK, response.New().SetMessage(username+" logout").SetData(flag))
}

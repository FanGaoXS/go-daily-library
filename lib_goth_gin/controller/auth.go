package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"net/http"
)

func BeginAuth(c *gin.Context) {
	// 前往验证（会自动从request中解析provider的值，然后跳转到对应的provider供应商进行验证）
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func AuthCallBack(c *gin.Context) {
	provider := c.Param("provider")
	// 会自动解析来自provider供应商的回调返回的user
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	} // 验证失败
	c.JSON(http.StatusOK, gin.H{
		provider + " user": user,
	})
}

func Logout(c *gin.Context) {
	provider := c.Param("provider")
	// 会自动从request解析provider的值，然后对对应的provider供应商进行logout
	err := gothic.Logout(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": provider + " logout error" + err.Error(),
		})
		return
	} // 退出失败
	c.JSON(http.StatusOK, gin.H{
		"message": provider + " logout success",
	})
}

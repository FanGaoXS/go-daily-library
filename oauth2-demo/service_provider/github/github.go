package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ClientId     = "3b81f5ce0886aff6ac2a"
	ClientSecret = "720d66f2f5a7f806d3807efdbdfd6ecd4f4d9f44"
)

type UserInfo struct {
	HtmlUrl   string    `json:"html_url"`
	AvatarUrl string    `json:"avatar_url"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func Authorize(c *gin.Context) {
	u, _ := url.Parse("https://github.com/login/oauth/authorize")
	values := u.Query()
	values.Add("client_id", ClientId)
	values.Add("redirect_uri", "http://localhost:8090/github/oauth/callback")
	u.RawQuery = values.Encode()

	c.Redirect(http.StatusMovedPermanently, u.String())
	return
}

// Callback
// 1 来自服务提供商返回的授权码
// 2 使用clientId、clientSecret、授权码code向服务商请求令牌
// 3 向服务提供商发起POST请求以获取token
func Callback(c *gin.Context) {
	code := c.Query("code")
	u, _ := url.Parse("https://github.com/login/oauth/access_token")
	values := u.Query()
	values.Add("client_id", ClientId)
	values.Add("client_secret", ClientSecret)
	values.Add("code", code)
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Add("accept", "application/json")
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	obj := make(map[string]interface{})
	json.Unmarshal(bytes, &obj)
	c.JSON(http.StatusOK, obj)
	return
}

// Userinfo
// 使用token获取userinfo
func Userinfo(c *gin.Context) {
	token := c.Query("token")
	u, _ := url.Parse("https://api.github.com/user")

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "token "+token)
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	info := new(UserInfo)
	json.Unmarshal(bytes, &info)
	c.JSON(http.StatusOK, info)
	return
}

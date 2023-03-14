package gitlab

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
	ClientId     = "193ef592ac3ca29aef70eaeacc40047b2810c73b7d9120d52fc47273cd94b34b"
	ClientSecret = "d96fceea906cd5d1aa1e347361358cd5cae0bc0fd225101b586c366fd5e395b5"
)

type UserInfo struct {
	WebUrl    string    `json:"web_url"`
	AvatarUrl string    `json:"avatar_url"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func Authorize(c *gin.Context) {
	u, _ := url.Parse("https://gitlab.com/oauth/authorize")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("redirect_uri", "http://localhost:8090/gitlab/oauth/callback")
	values.Set("response_type", "code")
	values.Set("state", "123")
	values.Add("scope", "api")
	values.Add("scope", "read_api")
	values.Add("scope", "read_user")
	values.Add("scope", "profile")
	u.RawQuery = values.Encode()

	c.Redirect(http.StatusMovedPermanently, u.String())
	return
}

func Callback(c *gin.Context) {
	code := c.Query("code")
	u, _ := url.Parse("https://gitlab.com/oauth/token")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("client_secret", ClientSecret)
	values.Set("code", code)
	values.Set("grant_type", "authorization_code")
	values.Set("redirect_uri", "http://localhost:8090/gitlab/oauth/callback")
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	req.Header.Set("accept", "application/json")
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	obj := make(map[string]interface{})
	json.Unmarshal(bytes, &obj)
	c.JSON(http.StatusOK, obj)
	return
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.Query("refresh_token")
	u, _ := url.Parse("https://gitlab.com/oauth/token")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("client_secret", ClientSecret)
	values.Set("refresh_token", refreshToken)
	values.Set("grant_type", "refresh_token")
	values.Set("redirect_uri", "http://localhost:8090/gitlab/oauth/callback")
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodPost, u.String(), nil)
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	obj := make(map[string]interface{})
	json.Unmarshal(bytes, &obj)
	c.JSON(http.StatusOK, obj)
	return
}

func Userinfo(c *gin.Context) {
	token := c.Query("token")
	u, _ := url.Parse("https://gitlab.com/api/v4/user")

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("Authorization", "Bearer "+token)
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	info := new(UserInfo)
	json.Unmarshal(bytes, &info)
	c.JSON(http.StatusOK, info)
	return
}

package gitee

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
	ClientId     = "894c85cea5b36cd629a3cf0867778692acbad55877487e111f4683312bc512cb"
	ClientSecret = "dd3ad7b413b5026edc08fbdf1999baacdaccab7665466c2df28aeb245c568cae"
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
	u, _ := url.Parse("https://gitee.com/oauth/authorize")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("redirect_uri", "http://localhost:8090/gitee/oauth/callback")
	values.Set("response_type", "code")
	u.RawQuery = values.Encode()

	c.Redirect(http.StatusMovedPermanently, u.String())
	return
}

func Callback(c *gin.Context) {
	code := c.Query("code")
	u, _ := url.Parse("https://gitee.com/oauth/token")
	values := u.Query()
	values.Set("client_id", ClientId)
	values.Set("client_secret", ClientSecret)
	values.Set("code", code)
	values.Set("grant_type", "authorization_code")
	values.Set("redirect_uri", "http://localhost:8090/gitee/oauth/callback")
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
	u, _ := url.Parse("https://gitee.com/oauth/token")
	values := u.Query()
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", refreshToken)
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

	u, _ := url.Parse("https://gitee.com/api/v5/user")
	values := u.Query()
	values.Set("access_token", token)
	u.RawQuery = values.Encode()

	client := http.DefaultClient
	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("accept", "application/json")
	res, _ := client.Do(req)
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bytes))

	info := new(UserInfo)
	json.Unmarshal(bytes, &info)
	c.JSON(http.StatusOK, info)
	return
}

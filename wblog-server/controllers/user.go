package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gookit/color"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"wblog/utils"

	"github.com/alimoeeny/gooauth2"
	"github.com/cihub/seelog"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"wblog/helpers"
	"wblog/models"
	"wblog/system"
)

type NewJwtClaims struct {
	*models.User
	jwt.StandardClaims
}

var SecretKey = []byte("9hUxqaGelNnCZaCW")

type GithubUserInfo struct {
	AvatarURL         string      `json:"avatar_url"`
	Bio               interface{} `json:"bio"`
	Blog              string      `json:"blog"`
	Company           interface{} `json:"company"`
	CreatedAt         string      `json:"created_at"`
	Email             interface{} `json:"email"`
	EventsURL         string      `json:"events_url"`
	Followers         int         `json:"followers"`
	FollowersURL      string      `json:"followers_url"`
	Following         int         `json:"following"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	GravatarID        string      `json:"gravatar_id"`
	Hireable          interface{} `json:"hireable"`
	HTMLURL           string      `json:"html_url"`
	ID                int         `json:"id"`
	Location          interface{} `json:"location"`
	Login             string      `json:"login"`
	Name              interface{} `json:"name"`
	OrganizationsURL  string      `json:"organizations_url"`
	PublicGists       int         `json:"public_gists"`
	PublicRepos       int         `json:"public_repos"`
	ReceivedEventsURL string      `json:"received_events_url"`
	ReposURL          string      `json:"repos_url"`
	SiteAdmin         bool        `json:"site_admin"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	Type              string      `json:"type"`
	UpdatedAt         string      `json:"updated_at"`
	URL               string      `json:"url"`
}
type SignInInfo struct {
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

// SignupPost 注册post
func SignupPost(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	var user = models.User{}
	if err1 := c.ShouldBind(&user); err1 != nil {
		color.Redln(user)
		color.Redln("绑定user失败")

	}

	if len(user.Email) == 0 || len(user.Password) == 0 {
		res["message"] = "email or password cannot be null"
		return
	}
	user.Password = helpers.Md5(user.Email + user.Password)
	err = user.Insert()
	if err != nil {
		res["message"] = "创建失败"
		return
	}
	res["succeed"] = true
}
func LogoutGet(c *gin.Context) {
	c.JSON(202, gin.H{"error": "失败了"})
}
func Login(c *gin.Context) {
	type SignIn struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	result := &models.Result{
		Code:    200,
		Message: "登录成功n",
		Data:    nil,
	}
	var (
		err    error
		user   *models.User
		signin SignIn
	)
	if c.ShouldBind(&signin) != nil {
		color.Redln("绑定signin失败")
	}

	username := signin.Username
	password := signin.Password
	color.Redln("html登录了")
	fmt.Print(username)
	fmt.Println(password)
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "username or password cannot be null",
		})
		return
	}

	user, err = models.GetUserByUsername(username)
	if err != nil || user.Password != helpers.Md5(username+password) {
		c.JSON(http.StatusOK, gin.H{
			"message": "invalid username or password",
		})
		return
	}
	if user.LockState {
		c.JSON(http.StatusOK, gin.H{
			"message": "Your account have been locked",
		})
		return
	}

	if user.IsAdmin {
		expiresTime := time.Now().Unix() + int64(60*60*24)
		//claims := jwt.StandardClaims{
		//	Audience:  user.Username,          // 受众
		//	ExpiresAt: expiresTime,            // 失效时间
		//	Id:        string(rune(user.Uid)), // 编号
		//	IssuedAt:  time.Now().Unix(),      // 签发时间
		//	Issuer:    sqlU.Username,            // 签发人
		//	NotBefore: time.Now().Unix(),      // 生效时间
		//	Subject:   "login",                // 主题
		//}
		stdClaims := jwt.StandardClaims{

			Audience:  "啊啊啊",             // 受众
			ExpiresAt: expiresTime,       // 失效时间
			Id:        "id",              // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "sqlU.Username",   // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}
		newClaims := NewJwtClaims{
			User:           user,
			StandardClaims: stdClaims,
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
		if token, err := tokenClaims.SignedString(SecretKey); err == nil {
			result.Message = "登录成功"
			result.Data = token
			result.Code = http.StatusOK
			c.JSON(result.Code, result)
		} else {
			result.Message = "登录失败，请重新登陆"
			result.Code = http.StatusOK
			c.JSON(result.Code, gin.H{
				"result": result,
			})
		}

	} else {
		c.JSON(http.StatusMovedPermanently, gin.H{
			"err": "重定向",
		})
	}
}

func Oauth2Callback(c *gin.Context) {
	var (
		userInfo *GithubUserInfo
		user     *models.User
	)
	code := c.Query("code")
	//state := c.Query("state")
	//
	//// validate state
	//session := sessions.Default(c)
	//if len(state) == 0 || state != session.Get(SESSION_GITHUB_STATE) {
	//	c.Abort()
	//	return
	//}
	// remove state from session
	//session.Delete(SESSION_GITHUB_STATE)
	//session.Save()

	// exchange accesstoken by code
	token, err := exchangeTokenByCode(code)
	if err != nil {
		seelog.Error(err)
		c.Redirect(http.StatusMovedPermanently, "/signin")
		return
	}

	//get github userinfo by accesstoken
	userInfo, err = getGithubUserInfoByAccessToken(token)
	if err != nil {
		seelog.Error(err)
		c.Redirect(http.StatusMovedPermanently, "/signin")
		return
	}

	sessionUser, exists := c.Get(CONTEXT_USER_KEY)
	if exists { // 已登录
		user, _ = sessionUser.(*models.User)
		_, err1 := models.IsGithubIdExists(userInfo.Login, user.ID)
		if err1 != nil { // 未绑定
			if user.IsAdmin {
				user.GithubLoginId = userInfo.Login
			}
			user.AvatarUrl = userInfo.AvatarURL
			user.GithubUrl = userInfo.HTMLURL
			err = user.UpdateGithubUserInfo()
		} else {
			err = errors.New("this github loginId has bound another account.")
		}
	} else {
		user = &models.User{
			GithubLoginId: userInfo.Login,
			AvatarUrl:     userInfo.AvatarURL,
			GithubUrl:     userInfo.HTMLURL,
		}
		user, err = user.FirstOrCreate()
		if err == nil {
			if user.LockState {
				err = errors.New("Your account have been locked.")
				HandleMessage(c, "Your account have been locked.")
				return
			}
		}
	}

	if err == nil {
		s := sessions.Default(c)
		s.Clear()
		s.Set(SESSION_KEY, user.ID)
		s.Save()
		if user.IsAdmin {
			c.Redirect(http.StatusMovedPermanently, "/admin/index")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
		return
	}
}

func exchangeTokenByCode(code string) (accessToken string, err error) {
	var (
		transport *oauth.Transport
		token     *oauth.Token
	)
	transport = &oauth.Transport{Config: &oauth.Config{
		ClientId:     system.GetConfiguration().GithubClientId,
		ClientSecret: system.GetConfiguration().GithubClientSecret,
		RedirectURL:  system.GetConfiguration().GithubRedirectURL,
		TokenURL:     system.GetConfiguration().GithubTokenUrl,
		Scope:        system.GetConfiguration().GithubScope,
	}}
	token, err = transport.Exchange(code)
	if err != nil {
		return
	}
	accessToken = token.AccessToken
	// cache token
	tokenCache := oauth.CacheFile("./request.token")
	if err := tokenCache.PutToken(token); err != nil {
		seelog.Error(err)
	}
	return
}

func getGithubUserInfoByAccessToken(token string) (*GithubUserInfo, error) {
	var (
		resp *http.Response
		body []byte
		err  error
	)
	resp, err = http.Get(fmt.Sprintf("https://api.github.com/user?access_token=%s", token))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var userInfo GithubUserInfo
	err = json.Unmarshal(body, &userInfo)
	return &userInfo, err
}

func ProfileGet(c *gin.Context) {
	userCtx, exists := c.Get("user")
	color.Redln(userCtx)
	//查询用户组及该组的功能权限
	user, _ := userCtx.(models.User)
	if exists {
		utils.JSON(c, http.StatusOK, "success", gin.H{
			"user":     user,
			"comments": models.MustListUnreadComment(),
		})
	}
}

func ProfileUpdate(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)

	avatarUrl := c.PostForm("avatarUrl")
	nickName := c.PostForm("nickName")
	sessionUser, _ := c.Get(CONTEXT_USER_KEY)
	user, ok := sessionUser.(*models.User)
	if !ok {
		res["message"] = "server interval error"
		return
	}
	err = user.UpdateProfile(avatarUrl, nickName)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
	res["user"] = models.User{AvatarUrl: avatarUrl, NickName: nickName}
	utils.JSON(c, 200, "success", res)
}

func BindEmail(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer utils.JSON(c, 200, "success", res)
	email := c.PostForm("email")
	sessionUser, _ := c.Get(CONTEXT_USER_KEY)
	user, ok := sessionUser.(*models.User)
	if !ok {
		res["message"] = "server interval error"
		return
	}
	if len(user.Email) > 0 {
		res["message"] = "email have bound"
		return
	}
	_, err = models.GetUserByUsername(email)
	if err == nil {
		res["message"] = "email have be registered"
		return
	}
	err = user.UpdateEmail(email)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func UnbindEmail(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	sessionUser, _ := c.Get(CONTEXT_USER_KEY)
	user, ok := sessionUser.(*models.User)
	if !ok {
		res["message"] = "server interval error"
		return
	}
	if user.Email == "" {
		res["message"] = "email haven't bound"
		return
	}
	err = user.UpdateEmail("")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func UnbindGithub(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	sessionUser, _ := c.Get(CONTEXT_USER_KEY)
	user, ok := sessionUser.(*models.User)
	if !ok {
		res["message"] = "server interval error"
		return
	}
	if user.GithubLoginId == "" {
		res["message"] = "github haven't bound"
		return
	}
	user.GithubLoginId = ""
	err = user.UpdateGithubUserInfo()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func UserIndex(c *gin.Context) {
	users, _ := models.ListUsers()
	user, _ := c.Get(CONTEXT_USER_KEY)
	utils.JSON(c, http.StatusOK, "success", gin.H{
		"users":    users,
		"user":     user,
		"comments": models.MustListUnreadComment(),
	})
}

func UserLock(c *gin.Context) {
	var (
		err  error
		_id  uint64
		res  = gin.H{}
		user *models.User
	)
	defer writeJSON(c, res)
	id := c.Param("id")
	_id, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	user, err = models.GetUser(uint(_id))
	if err != nil {
		res["message"] = err.Error()
		return
	}
	user.LockState = !user.LockState
	err = user.Lock()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

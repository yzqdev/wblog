package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gookit/color"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"wblog-server/helpers"
	"wblog-server/models"
	"wblog-server/system"
)

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

var SecretKey = []byte("9hUxqaGelNnCZaCW")

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegUser struct {
	LoginUser
	Email string `json:"email"`
}
type NewJwtClaims struct {
	UserId string
	jwt.StandardClaims
}

func LogoutGet(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.Redirect(http.StatusSeeOther, "/signin")
}
func InitPage(c *gin.Context) {
	models.CreateAdmin()
	c.JSON(200, "success")
}
func SignupPost(c *gin.Context) {
	var (
		err error
	)
	regUser := &RegUser{}
	if err = c.ShouldBindJSON(regUser); err != nil {
		color.Redln("解析json失败")
	}
	user := &models.User{
		Uid:      xid.New().String(),
		Email:    regUser.Email,
		Username: regUser.Username,
		Password: regUser.Password,
		IsAdmin:  true,
	}
	if len(user.Username) == 0 || len(user.Password) == 0 {
		helpers.JSON(c, http.StatusOK, "username or password cannot be null", false)
		return
	}
	user.Password = helpers.Md5(user.Username + user.Password)
	err = user.Insert()
	if err != nil {
		helpers.JSON(c, http.StatusOK, "邮箱已存在", false)
		return
	}
	helpers.JSON(c, http.StatusOK, "注册成功", true)

}

func SigninPost(c *gin.Context) {
	var (
		err  error
		user *models.User
	)
	loginUser := &LoginUser{}
	if err = c.ShouldBindJSON(loginUser); err != nil {
		color.Redln("解析json失败")
	}
	username := loginUser.Username
	password := loginUser.Password
	if username == "" || password == "" {
		helpers.JSON(c, http.StatusBadRequest, "username or password cannot be null", false)
		return
	}
	color.Redln(helpers.Md5(username + password))
	user, err = models.GetUserByUsername(username)
	if err != nil || user.Password != helpers.Md5(username+password) {
		helpers.JSON(c, http.StatusUnauthorized, "invalid username or password", false)
		return
	}
	if user.LockState {
		helpers.JSON(c, http.StatusForbidden, "Your account have been locked", false)
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
			Issuer:    "yzqdev",          // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}
		newClaims := NewJwtClaims{
			UserId:         user.Uid,
			StandardClaims: stdClaims,
		}
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
		if token, err := tokenClaims.SignedString(SecretKey); err == nil {
			helpers.JSON(c, http.StatusOK, "true", map[string]interface{}{
				"isAdmin": true,
				"token":   token,
				"path":    "/admin/index",
			})
		} else {

			helpers.JSON(c, http.StatusOK, "登录失败，请重新登陆", false)
		}

	} else {
		helpers.JSON(c, http.StatusOK, "true", map[string]interface{}{
			"isAdmin": false,
			"path":    "/",
		})
	}
}

func Oauth2Callback(c *gin.Context) {
	var (
		userInfo *GithubUserInfo
		user     *models.User
	)
	code := c.Query("code")
	state := c.Query("state")

	// validate state
	session := sessions.Default(c)
	if len(state) == 0 || state != session.Get(helpers.SESSION_GITHUB_STATE) {
		c.Abort()
		return
	}
	// remove state from session
	session.Delete(helpers.SESSION_GITHUB_STATE)
	session.Save()

	// exchange accesstoken by code
	token, err := exchangeTokenByCode(code)
	if err != nil {
		system.BlogLog.Error("exchange token error", zap.Error(err))
		c.Redirect(http.StatusMovedPermanently, "/signin")
		return
	}

	//get github userinfo by accesstoken
	userInfo, err = getGithubUserInfoByAccessToken(token)
	if err != nil {
		system.BlogLog.Error("get github use error", zap.Error(err))
		c.Redirect(http.StatusMovedPermanently, "/signin")
		return
	}

	sessionUser, exists := c.Get(helpers.CONTEXT_USER_KEY)
	if exists { // 已登录
		user, _ = sessionUser.(*models.User)
		_, err1 := models.IsGithubIdExists(userInfo.Login, user.Id)
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
				helpers.HandleMessage(c, "Your account have been locked.")
				return
			}
		}
	}

	if err == nil {
		s := sessions.Default(c)
		s.Clear()
		s.Set(helpers.SESSION_KEY, user.Id)
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
		transport *helpers.Transport
		token     *helpers.Token
	)
	transport = &helpers.Transport{Config: &helpers.Config{
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
	tokenCache := helpers.CacheFile("./request.token")
	if err := tokenCache.PutToken(token); err != nil {
		system.BlogLog.Error("request token err", zap.Error(err))
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
	sessionUser, exists := c.Get(helpers.CONTEXT_USER_KEY)
	userId := sessionUser.(string)
	sqlUser, _ := models.GetUserByUid(userId)
	if exists {
		helpers.JSON(c, http.StatusOK, "admin/profile.html", gin.H{
			"user":     sqlUser,
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
	nickName := c.PostForm("nickname")
	sessionUser, _ := c.Get(helpers.CONTEXT_USER_KEY)
	userId, ok := sessionUser.(string)
	sqlUser, _ := models.GetUserByUid(userId)
	if !ok {
		res["message"] = "server interval error"
		return
	}
	err = sqlUser.UpdateProfile(avatarUrl, nickName)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
	res["user"] = sqlUser
	defer helpers.JSON(c, http.StatusOK, "success", res)
}

func BindEmail(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)
	email := c.PostForm("email")
	sessionUser, _ := c.Get(helpers.CONTEXT_USER_KEY)
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
	defer helpers.WriteJson(c, res)
	sessionUser, _ := c.Get(helpers.CONTEXT_USER_KEY)
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
	defer helpers.WriteJson(c, res)
	sessionUser, _ := c.Get(helpers.CONTEXT_USER_KEY)
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
func UserInfo(c *gin.Context) {
	userContext, exist := c.Get("userId")
	if !exist {
		color.Danger.Println("失败了")
	}
	//查询用户组及该组的功能权限
	userId, ok := userContext.(string) //这个是类型推断,判断接口是什么类型
	if !ok {

		color.Danger.Println("断言失败")
	}
	sqlUser, _ := models.GetUserByUid(userId)
	helpers.JSON(c, http.StatusOK, "success", sqlUser)
}
func UserIndex(c *gin.Context) {
	users, _ := models.ListUsers()
	helpers.JSON(c, http.StatusOK, "admin/user.html", gin.H{
		"users":    users,
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
	defer helpers.WriteJson(c, res)
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

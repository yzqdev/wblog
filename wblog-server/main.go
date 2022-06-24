package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"wblog-server/controllers"
	"wblog-server/helpers"
	"wblog-server/models"
	"wblog-server/system"
)

func main() {

	// 初始化zap日志库
	system.BlogLog = helpers.Zap()
	zap.ReplaceGlobals(system.BlogLog)
	_, err := models.InitDB()
	if err != nil {
		return
	}

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	setTemplate(router)
	setSessions(router)
	router.Use(SharedData())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,

		MaxAge: 12 * time.Hour,
	}))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	//Periodic tasks
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().Do(helpers.CreateXMLSitemap)
	s.Every(7).Days().Do(controllers.Backup)
	s.StartAsync()

	//router.Static("/static", "static")
	//router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	InitRouter(router)
	InitRouterV2(router)
	router.NoRoute(helpers.Handle404)

	router.Run(system.GetConfiguration().Addr)
}

func setTemplate(engine *gin.Engine) {

	funcMap := template.FuncMap{
		"dateFormat": helpers.DateFormat,
		"substring":  helpers.Substring,
		"isOdd":      helpers.IsOdd,
		"isEven":     helpers.IsEven,
		"truncate":   helpers.Truncate,
		"add":        helpers.Add,
		"minus":      helpers.Minus,
		"listtag":    helpers.ListTag,
	}

	engine.SetFuncMap(funcMap)

	engine.LoadHTMLGlob("views/**/*")
}

//setSessions initializes sessions & csrf middlewares
func setSessions(router *gin.Engine) {
	config := system.GetConfiguration()
	//https://github.com/gin-gonic/contrib/tree/master/sessions
	store := cookie.NewStore([]byte(config.SessionSecret))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}) //Also set Secure: true if using SSL, you should though
	router.Use(sessions.Sessions("gin-session", store))
	//https://github.com/utrack/gin-csrf
	/*router.Use(csrf.Middleware(csrf.Options{
		Secret: config.SessionSecret,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))*/
}

//+++++++++++++ middlewares +++++++++++++++++++++++

//SharedData fills in common data, such as user info, etc...
func SharedData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if uID := session.Get(helpers.SESSION_KEY); uID != nil {
			user, err := models.GetUser(uID)
			if err == nil {
				c.Set(helpers.CONTEXT_USER_KEY, user)
			}
		}
		if system.GetConfiguration().Reg {
			c.Set("Reg", true)
		}
		c.Next()
	}
}

//AuthRequired grants access to authenticated users, requires SharedData middleware
func AdminScopeRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(helpers.CONTEXT_USER_KEY); user != nil {
			if u, ok := user.(*models.User); ok && u.IsAdmin {
				c.Next()
				return
			}
		}
		system.BlogLog.Warn("User not authorized to visit  " + c.Request.RequestURI)
		helpers.JSON(c, http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden!",
		})
		c.Abort()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(helpers.CONTEXT_USER_KEY); user != nil {
			if _, ok := user.(*models.User); ok {
				c.Next()
				return
			}
		}
		system.BlogLog.Warn("User not authorized to visit  " + c.Request.RequestURI)
		helpers.JSON(c, http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden!",
		})
		c.Abort()
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		system.BlogLog.Error("warning", zap.Error(err))
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//func getCurrentDirectory() string {
//	return ""
//}

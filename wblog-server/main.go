package main

import (
	"flag"
	"github.com/claudiu/gocron"
	"github.com/gin-contrib/sessions"
	"html/template"
	"net/http"

	"path/filepath"

	"os"
	"strings"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"wblog/controllers"
	"wblog/helpers"
	"wblog/models"
	"wblog/system"
)

func main() {

	configFilePath := flag.String("C", "conf/conf.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")
	flag.Parse()

	logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		seelog.Critical("err parsing seelog config file", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	if err := system.LoadConfiguration(*configFilePath); err != nil {
		seelog.Critical("err parsing config log file", err)
		return
	}

	db, err := models.InitDB()
	if err != nil {
		seelog.Critical("err open databases", err)
		return
	}
	defer db.Close()
	//设置模式
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//setSessions(router)
	//router.Use(SharedData())
	router.Static("/static", "./static")
	//Periodic tasks

	setTemplate(router)

	gocron.Every(1).Day().Do(controllers.CreateXMLSitemap)
	gocron.Every(7).Days().Do(controllers.Backup)
	gocron.Start()
	InitRoute(router)

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
	engine.LoadHTMLGlob("./views/**/*")
}

//setSessions initializes sessions & csrf middlewares
func setSessions(router *gin.Engine) {
	//config := system.GetConfiguration()
	//https://github.com/gin-gonic/contrib/tree/master/sessions
	//store := sessions.NewCookieStore([]byte(config.SessionSecret))
	//store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}) //Also set Secure: true if using SSL, you should though
	//router.Use(sessions.Sessions("gin-session", store))
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
		if uID := session.Get(controllers.SESSION_KEY); uID != nil {
			user, err := models.GetUser(uID)
			if err == nil {
				c.Set(controllers.CONTEXT_USER_KEY, user)
			}
		}
		if system.GetConfiguration().SignupEnabled {
			c.Set("SignupEnabled", true)
		}
		c.Next()
	}
}

//AuthRequired grants access to authenticated users, requires SharedData middleware
func AdminScopeRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(controllers.CONTEXT_USER_KEY); user != nil {
			if u, ok := user.(*models.User); ok && u.IsAdmin {
				c.Next()
				return
			}
		}
		seelog.Warnf("User not authorized to visit %s", c.Request.RequestURI)
		c.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden!",
		})
		c.Abort()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(controllers.CONTEXT_USER_KEY); user != nil {
			if _, ok := user.(*models.User); ok {
				c.Next()
				return
			}
		}
		seelog.Warnf("User not authorized to visit %s", c.Request.RequestURI)
		c.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden!",
		})
		c.Abort()
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		seelog.Critical(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//func getCurrentDirectory() string {
//	return ""
//}

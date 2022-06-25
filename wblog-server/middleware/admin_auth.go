package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"wblog-server/helpers"
	"wblog-server/models"
	"wblog-server/system"
)

//AuthRequired grants access to authenticated users, requires SharedData middleware
func AdminScopeRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get(helpers.CONTEXT_USER_KEY)
		color.Redln(`userid`, user)
		if user != nil {
			u, ok := user.(string)

			sqlUser, _ := models.GetUserByUid(u)
			if ok && sqlUser.AdminRole {
				c.Next()
				return
			}
		}
		system.BlogLog.Warn("User not authorized to visit  " + c.Request.RequestURI)
		helpers.JSON(c, http.StatusForbidden, "需要管理员权限!", gin.H{
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

package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/steambap/captcha"
	"wblog-server/helpers"
)

func CaptchaGet(context *gin.Context) {
	session := sessions.Default(context)
	data, _ := captcha.NewMathExpr(150, 50)
	session.Delete(helpers.SESSION_CAPTCHA)
	session.Set(helpers.SESSION_CAPTCHA, data.Text)
	session.Save()
	data.WriteImage(context.Writer)
}

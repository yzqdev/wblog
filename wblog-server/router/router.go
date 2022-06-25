package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"wblog-server/controllers"
	v2 "wblog-server/controllers/v2"
	"wblog-server/middleware"
	"wblog-server/system"
)

func InitRouter(router *gin.Engine) {

	router.GET("/", controllers.IndexGet)
	router.GET("/index", controllers.IndexGet)
	router.GET("/rss", controllers.RssGet)

	if system.GetConfiguration().Reg {
		color.Redln("允许注册!")
		router.POST("/signup", controllers.SignupPost)
	}
	// user signin and logout
	router.POST("/signin", controllers.SigninPost)
	router.GET("/logout", controllers.LogoutGet)
	router.GET("/oauth2callback", controllers.Oauth2Callback)
	router.GET("/auth/:authType", controllers.AuthGet)
	// subscriber
	router.GET("/subscribe", controllers.SubscribeGet)
	router.POST("/subscribe", controllers.Subscribe)
	router.GET("/active", controllers.ActiveSubscriber)
	router.GET("/unsubscribe", controllers.UnSubscribe)

	router.GET("/page/:id", controllers.PageGet)
	router.GET("/post/:id", controllers.PostGet)
	router.GET("/tag/:tag", controllers.TagGet)
	router.GET("/archives/:year/:month", v2.ArchiveGet)

	router.GET("/link/:id", controllers.LinkGet)
	// captcha
	router.GET("/captcha", controllers.CaptchaGet)

	visitor := router.Group("/visitor")
	visitor.Use(middleware.AuthRequired())
	{
		visitor.POST("/new_comment", controllers.CommentPost)
		visitor.POST("/comment/:id/delete", controllers.CommentDelete)
	}
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.SigninPost)
		auth.POST("/reg", controllers.SignupPost)
		auth.GET("/init", controllers.InitPage)

	}
	adminRouter := router.Group("/admin", middleware.JwtHandler())
	{
		adminRouter.GET("/userInfo", controllers.UserInfo)
		adminRouter.GET("/posts", controllers.PostIndex)
		adminRouter.POST("/posts", controllers.PostCreate)
		adminRouter.DELETE("/posts/:id", controllers.PostDelete)
		adminRouter.GET("/links", controllers.LinkIndex)
		adminRouter.POST("/link", controllers.LinkCreate)
		adminRouter.DELETE("/link/:id", controllers.LinkDelete)
		adminRouter.GET("/comment/unread", controllers.ListCommentUnRead)

		//user
		adminRouter.GET("/profile", controllers.ProfileGet)
		adminRouter.POST("/profile", controllers.ProfileUpdate)
	}
	//超级管理员
	authorized := router.Group("/super", middleware.JwtHandler(), middleware.AdminScopeRequired())
	{
		// index
		authorized.GET("/index", controllers.AdminIndex)

		// image upload
		authorized.POST("/upload", controllers.Upload)

		// page
		authorized.GET("/page", controllers.PageIndex)
		authorized.POST("/new_page", controllers.PageCreate)
		authorized.GET("/page/:id/edit", controllers.PageEdit)
		authorized.POST("/page/:id/edit", controllers.PageUpdate)
		authorized.POST("/page/:id/publish", controllers.PagePublish)
		authorized.POST("/page/:id/delete", controllers.PageDelete)

		// post
		authorized.GET("/post", controllers.PostIndex)
		authorized.POST("/new_post", controllers.PostCreate)
		authorized.GET("/post/:id/edit", controllers.PostEdit)
		authorized.POST("/post/:id/edit", controllers.PostUpdate)
		authorized.POST("/post/:id/publish", controllers.PostPublish)
		authorized.POST("/post/:id/delete", controllers.PostDelete)

		// tag
		authorized.POST("/new_tag", controllers.TagCreate)

		//
		authorized.GET("/user", controllers.UserIndex)
		authorized.POST("/user/:id/lock", controllers.UserLock)

		// profile

		authorized.POST("/profile/email/bind", controllers.BindEmail)
		authorized.POST("/profile/email/unbind", controllers.UnbindEmail)
		authorized.POST("/profile/github/unbind", controllers.UnbindGithub)

		// subscriber
		authorized.GET("/subscriber", controllers.SubscriberIndex)
		authorized.POST("/subscriber", controllers.SubscriberPost)

		// link
		authorized.GET("/link", controllers.LinkIndex)
		authorized.POST("/new_link", controllers.LinkCreate)
		authorized.POST("/link/:id/edit", controllers.LinkUpdate)
		authorized.POST("/link/:id/delete", controllers.LinkDelete)

		// comment
		authorized.POST("/comment/:id", controllers.CommentRead)
		authorized.POST("/read_all", controllers.CommentReadAll)

		// backup
		authorized.POST("/backup", controllers.BackupPost)
		authorized.POST("/restore", controllers.RestorePost)

		// mail
		authorized.POST("/new_mail", controllers.SendMail)
		authorized.POST("/new_batchmail", controllers.SendBatchMail)
	}
	homeRouter := router.Group("/home")
	{
		homeRouter.GET("/posts", controllers.PostIndex)
		homeRouter.GET("/post/:id", controllers.PostGet)
		homeRouter.GET("/links", controllers.LinkIndex)

		homeRouter.POST("/comment/:postId", controllers.CommentPost)
	}

}

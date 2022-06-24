package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wblog-server/helpers"
	"wblog-server/models"
)

func CommentPost(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)

	s := sessions.Default(c)

	verifyCode := c.PostForm("verifyCode")
	captchaId := s.Get(helpers.SESSION_CAPTCHA)
	_captchaId, _ := captchaId.(string)
	if !(_captchaId == verifyCode) {
		res["message"] = "error verifycode"
		return
	}

	postId := c.Param("postId")
	content := c.PostForm("content")
	nickname := c.PostForm("nickname")
	if len(content) == 0 {
		res["message"] = "content cannot be empty."
		return
	}

	//post, err := models.GetPostById(postId)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	pid, err := strconv.ParseUint(postId, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	comment := &models.Comment{
		PostID:   uint(pid),
		Content:  content,
		Nickname: nickname,
	}
	err = comment.Insert()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	//helpers.NotifyEmail("[wblog]您有一条新评论", fmt.Sprintf("<a href=\"%s/post/%d\" target=\"_blank\">%s</a>:%s", system.GetConfiguration().Domain, post.ID, post.Title, content))
	res["succeed"] = true
	helpers.JSON(c, http.StatusOK, "success", res)
}

func CommentDelete(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)

	s := sessions.Default(c)
	sessionUserID := s.Get(helpers.SESSION_KEY)
	userId, _ := sessionUserID.(uint)

	commentId := c.Param("id")

	comment := &models.Comment{
		UserID: uint(userId),
	}
	comment.ID = commentId
	err = comment.Delete()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}
func ListCommentUnRead(c *gin.Context) {
	comments := models.MustListUnreadComment()
	helpers.JSON(c, http.StatusOK, "获取成功", comments)
}

func CommentRead(c *gin.Context) {
	var (
		id  string
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)
	id = c.Param("id")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	comment := new(models.Comment)
	comment.ID = id
	err = comment.Update()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func CommentReadAll(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)
	err = models.SetAllCommentRead()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

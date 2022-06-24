package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"wblog-server/helpers"
	"wblog-server/models"
	"wblog-server/system"
)

func CommentPost(c *gin.Context) {
	var (
		err  error
		res  = gin.H{}
		post *models.Post
	)
	defer helpers.WriteJson(c, res)
	s := sessions.Default(c)
	sessionUserID := s.Get(helpers.SESSION_KEY)
	userId, _ := sessionUserID.(uint)

	verifyCode := c.PostForm("verifyCode")
	captchaId := s.Get(helpers.SESSION_CAPTCHA)
	s.Delete(helpers.SESSION_CAPTCHA)
	_captchaId, _ := captchaId.(string)
	if !(_captchaId == verifyCode) {
		res["message"] = "error verifycode"
		return
	}

	postId := c.PostForm("postId")
	content := c.PostForm("content")
	if len(content) == 0 {
		res["message"] = "content cannot be empty."
		return
	}

	post, err = models.GetPostById(postId)
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
		PostID:  uint(pid),
		Content: content,
		UserID:  userId,
	}
	err = comment.Insert()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	helpers.NotifyEmail("[wblog]您有一条新评论", fmt.Sprintf("<a href=\"%s/post/%d\" target=\"_blank\">%s</a>:%s", system.GetConfiguration().Domain, post.ID, post.Title, content))
	res["succeed"] = true
}

func CommentDelete(c *gin.Context) {
	var (
		err error
		res = gin.H{}
		cid string
	)
	defer helpers.WriteJson(c, res)

	s := sessions.Default(c)
	sessionUserID := s.Get(helpers.SESSION_KEY)
	userId, _ := sessionUserID.(uint)

	commentId := c.Param("id")
	cid = commentId
	if err != nil {
		res["message"] = err.Error()
		return
	}
	comment := &models.Comment{
		UserID: uint(userId),
	}
	comment.ID = cid
	err = comment.Delete()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func CommentRead(c *gin.Context) {
	var (
		id  string
		_id string
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)
	id = c.Param("id")
	_id = id
	if err != nil {
		res["message"] = err.Error()
		return
	}
	comment := new(models.Comment)
	comment.ID = _id
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

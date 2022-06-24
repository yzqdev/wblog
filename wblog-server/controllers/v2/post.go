package controllers

import (
	"github.com/gookit/color"
	"net/http"
	"strings"
	"wblog-server/helpers"

	"github.com/gin-gonic/gin"
	"wblog-server/models"
)

func PostGet(c *gin.Context) {
	id := c.Param("id")
	post, err := models.GetPostById(id)
	if err != nil || !post.IsPublished {
		helpers.Handle404(c)
		return
	}
	post.View++
	post.UpdateView()
	post.Tags, _ = models.ListTagByPostId(id)
	post.Comments, _ = models.ListCommentByPostID(id)
	helpers.JSON(c, http.StatusOK, "获取成功", post)
}

func PostNew(c *gin.Context) {
	c.HTML(http.StatusOK, "post/new.html", nil)
}

func PostCreate(c *gin.Context) {
	tags := c.PostForm("tags")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("is_published")
	published := "true" == isPublished
	color.Redln(isPublished)
	color.Redln("发布")
	post := &models.Post{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	err := post.Insert()
	if err != nil {
		helpers.JSON(c, http.StatusInternalServerError, "失败", gin.H{
			"post":    post,
			"message": err.Error(),
		})
		return
	}

	// add tag for post
	if len(tags) > 0 {
		tagArr := strings.Split(tags, ",")
		for _, tag := range tagArr {

			pt := &models.PostTag{
				PostId: post.ID,
				TagId:  tag,
			}
			pt.Insert()
		}
	}
	helpers.JSON(c, http.StatusOK, "成功", true)
	//c.Redirect(http.StatusMovedPermanently, "/admin/post")
}

func PostEdit(c *gin.Context) {
	id := c.Param("id")
	post, err := models.GetPostById(id)
	if err != nil {
		helpers.Handle404(c)
		return
	}
	post.Tags, _ = models.ListTagByPostId(id)
	c.HTML(http.StatusOK, "post/modify.html", gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	tags := c.PostForm("tags")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished

	post := &models.Post{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	post.ID = id
	err := post.Update()
	if err != nil {
		c.HTML(http.StatusOK, "post/modify.html", gin.H{
			"post":    post,
			"message": err.Error(),
		})
		return
	}
	// 删除tag
	models.DeletePostTagByPostId(post.ID)
	// 添加tag
	if len(tags) > 0 {
		tagArr := strings.Split(tags, ",")
		for _, tag := range tagArr {

			pt := &models.PostTag{
				PostId: post.ID,
				TagId:  tag,
			}
			pt.Insert()
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/post")
}

func PostPublish(c *gin.Context) {
	var (
		err  error
		res  = gin.H{}
		post *models.Post
	)
	defer helpers.WriteJson(c, res)
	id := c.Param("id")
	post, err = models.GetPostById(id)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post.IsPublished = !post.IsPublished
	err = post.Update()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func PostDelete(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.JSON(c, http.StatusOK, "成功", res)
	id := c.Param("id")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post := &models.Post{}
	post.ID = id
	err = post.Delete()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	models.DeletePostTagByPostId(id)
	res["succeed"] = true
}

func PostIndex(c *gin.Context) {
	posts, _ := models.ListAllPost("")
	user, _ := c.Get(helpers.CONTEXT_USER_KEY)
	helpers.JSON(c, http.StatusOK, "获取成功", gin.H{
		"posts":    posts,
		"Active":   "posts",
		"user":     user,
		"comments": models.MustListUnreadComment(),
	})

}

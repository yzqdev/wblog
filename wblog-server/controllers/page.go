package controllers

import (
	"net/http"
	"strconv"
	"wblog-server/helpers"

	"github.com/gin-gonic/gin"
	"wblog-server/models"
)

func PageGet(c *gin.Context) {
	id := c.Param("id")
	page, err := models.GetPageById(id)
	if err != nil || !page.IsPublished {
		helpers.Handle404(c)
		return
	}
	page.View++
	page.UpdateView()
	c.HTML(http.StatusOK, "page/display.html", gin.H{
		"page": page,
	})
}

func PageNew(c *gin.Context) {
	c.HTML(http.StatusOK, "page/new.html", nil)
}

func PageCreate(c *gin.Context) {
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished
	page := &models.Page{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	err := page.Insert()
	if err != nil {
		c.HTML(http.StatusOK, "page/new.html", gin.H{
			"message": err.Error(),
			"page":    page,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/page")
}

func PageEdit(c *gin.Context) {
	id := c.Param("id")
	page, err := models.GetPageById(id)
	if err != nil {
		helpers.Handle404(c)
	}
	c.HTML(http.StatusOK, "page/modify.html", gin.H{
		"page": page,
	})
}

func PageUpdate(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	page := &models.Page{Title: title, Body: body, IsPublished: published}
	page.ID = uint(pid)
	err = page.Update()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/page")
}

func PagePublish(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)
	id := c.Param("id")
	page, err := models.GetPageById(id)
	if err == nil {
		res["message"] = err.Error()
		return
	}
	page.IsPublished = !page.IsPublished
	err = page.Update()
	if err == nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func PageDelete(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)
	id := c.Param("id")
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	page := &models.Page{}
	page.ID = uint(pid)
	err = page.Delete()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func PageIndex(c *gin.Context) {
	pages, _ := models.ListAllPage()
	user, _ := c.Get(helpers.CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "admin/page.html", gin.H{
		"pages":    pages,
		"user":     user,
		"comments": models.MustListUnreadComment(),
	})
}

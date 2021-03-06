package controllers

import (
	"net/http"
	"strconv"
	"wblog-server/helpers"

	"github.com/gin-gonic/gin"
	"wblog-server/models"
)

func LinkIndex(c *gin.Context) {
	links, _ := models.ListLinks()

	helpers.JSON(c, http.StatusOK, "admin/link.html", gin.H{
		"links": links,
		//"user":     user,
		"comments": models.MustListUnreadComment(),
	})
}

func LinkCreate(c *gin.Context) {
	var (
		err   error
		res   = gin.H{}
		_sort int64
	)
	defer helpers.JSON(c, http.StatusOK, "success", res)
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")
	if len(name) == 0 || len(url) == 0 {
		res["message"] = "error parameter"
		return
	}
	_sort, err = strconv.ParseInt(sort, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	link := &models.Link{
		Name: name,
		Url:  url,
		Sort: int(_sort),
	}
	err = link.Insert()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func LinkUpdate(c *gin.Context) {
	var (
		_sort int64
		err   error
		res   = gin.H{}
	)
	defer helpers.JSON(c, http.StatusOK, "success", res)
	id := c.Param("id")
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")
	if len(id) == 0 || len(name) == 0 || len(url) == 0 {
		res["message"] = "error parameter"
		return
	}
	if err != nil {
		res["message"] = err.Error()
		return
	}
	_sort, err = strconv.ParseInt(sort, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	link := &models.Link{
		Name: name,
		Url:  url,
		Sort: int(_sort),
	}
	link.Id = id
	err = link.Update()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func LinkGet(c *gin.Context) {
	id := c.Param("id")
	_id, _ := strconv.ParseInt(id, 10, 64)
	link, err := models.GetLinkById(uint(_id))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	link.View++
	link.Update()
	c.Redirect(http.StatusFound, link.Url)
}

func LinkDelete(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.JSON(c, http.StatusOK, "success", true)
	id := c.Param("id")
	if err != nil {
		res["message"] = err.Error()
		return
	}

	link := new(models.Link)
	link.Id = id
	err = link.Delete()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

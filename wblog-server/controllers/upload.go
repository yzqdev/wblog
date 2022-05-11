package controllers

import (
	"mime/multipart"
	"wblog-server/helpers"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	var (
		err      error
		res      = gin.H{}
		url      string
		uploader Uploader
		file     multipart.File
		fh       *multipart.FileHeader
	)
	defer helpers.WriteJson(c, res)
	file, fh, err = c.Request.FormFile("file")
	if err != nil {
		res["message"] = err.Error()
		return
	}

	//uploader = QiniuUploader{}
	uploader = SmmsUploader{}

	url, err = uploader.upload(file, fh)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
	res["url"] = url
}

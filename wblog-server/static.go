package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed static views
var indexFs embed.FS

func AssetsStatic() http.FileSystem {
	files, _ := fs.Sub(indexFs, "static")
	return http.FS(files)
}
func AssetsViews() http.FileSystem {
	files, _ := fs.Sub(indexFs, "views")
	return http.FS(files)
}

func Static(router *gin.Engine) {
	//router.NoRoute(func(c *gin.Context) {
	//	c.Header("Content-Type", "text/html")
	//	c.Status(200)
	//	index, _ := Assets().Open("index.html")
	//	indexHtml, _ := ioutil.ReadAll(index)
	//	_, _ = c.Writer.WriteString(string(indexHtml))
	//
	//	c.Writer.Flush()
	//	c.Writer.WriteHeaderNow()
	//})
	router.StaticFS("/static", AssetsStatic())
	router.StaticFS("/views", AssetsViews())

}

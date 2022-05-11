package controllers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"wblog-server/helpers"
	"wblog-server/system"
)

func BackupPost(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer helpers.WriteJson(c, res)
	err = Backup()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func RestorePost(c *gin.Context) {
	var (
		fileName  string
		fileUrl   string
		err       error
		res       = gin.H{}
		resp      *http.Response
		bodyBytes []byte
	)
	defer helpers.WriteJson(c, res)
	fileName = c.PostForm("fileName")
	if fileName == "" {
		res["message"] = "fileName cannot be empty."
		return
	}
	fileUrl = system.GetConfiguration().QiniuFileServer + fileName
	resp, err = http.Get(fileUrl)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	bodyBytes, err = helpers.Decrypt(bodyBytes, system.GetConfiguration().BackupKey)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	err = ioutil.WriteFile(fileName, bodyBytes, os.ModePerm)
	if err == nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func Backup() (err error) {
	var (
		u           *url.URL
		exist       bool
		ret         PutRet
		bodyBytes   []byte
		encryptData []byte
	)
	u, err = url.Parse(system.GetConfiguration().DSN)
	if err != nil {
		system.BlogLog.Error("parse dsn error:%v", zap.Error(err))
		return
	}
	exist, _ = helpers.PathExists(u.Path)
	if !exist {
		err = errors.New("database file doesn't exists.")
		system.BlogLog.Error("database file doesn't exists.", zap.Error(err))
		return
	}
	system.BlogLog.Warn("start backup...")
	bodyBytes, err = ioutil.ReadFile(u.Path)
	if err != nil {
		system.BlogLog.Error("read error", zap.Error(err))
		return
	}
	encryptData, err = helpers.Encrypt(bodyBytes, system.GetConfiguration().BackupKey)
	if err != nil {
		system.BlogLog.Error("encrypt error", zap.Error(err))
		return
	}

	putPolicy := storage.PutPolicy{
		Scope: system.GetConfiguration().QiniuBucket,
	}
	mac := qbox.NewMac(system.GetConfiguration().QiniuAccessKey, system.GetConfiguration().QiniuSecretKey)
	token := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	uploader := storage.NewFormUploader(&cfg)
	putExtra := storage.PutExtra{}

	fileName := fmt.Sprintf("wblog_%s.db", helpers.GetCurrentTime().Format("20060102150405"))
	err = uploader.Put(context.Background(), &ret, token, fileName, bytes.NewReader(encryptData), int64(len(encryptData)), &putExtra)
	if err != nil {
		system.BlogLog.Error("backup error:%v", zap.Error(err))
		return
	}
	system.BlogLog.Warn("backup succeefully.")
	return err
}

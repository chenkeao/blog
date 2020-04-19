package controllers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"wblog/models"
)

func BannerIndex(ctx *gin.Context) {
	images, err := models.ListAllImages()
	user, _ := ctx.Get(CONTEXT_USER_KEY)
	if err != nil {
		log.Println(err.Error())
		return
	}
	ctx.HTML(http.StatusOK, "admin/banner.html", gin.H{
		"images":   images,
		"user":     user,
			"comments": models.MustListUnreadComment(),
	})
}

func BannerDelete(ctx *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(ctx, res)
	id := ctx.Param("id")
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	banner := &models.BannerImage{}
	banner.ID = uint(pid)
	err = models.DB.Where("id = ?",pid).Find(&banner).Error
	if err != nil {
		res["message"] = err.Error()
		return
	}
	err = banner.Delete()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	err = os.Remove(banner.Path)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func BannerNew(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	desc := ctx.PostForm("desc")
	url := ctx.PostForm("url")
	if err != nil {
		log.Println(err.Error())
		ctx.HTML(http.StatusOK, "errors/error.html", gin.H{
			"message": "文件上传失败",
		})
		return
	}
	//path := RootPath()
	path := "static//banner"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		ctx.HTML(http.StatusOK, "errors/error.html", gin.H{
			"message": "文件上传失败",
		})
		log.Panicln("创建文件夹失败", err.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	path = filepath.Join(path, fileName)
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.HTML(http.StatusOK, "errors/error.html", gin.H{
			"message": "文件上传失败",
		})
		log.Panicln("文件保存失败", err.Error())
	}
	if err = models.SaveImage(desc, url, path); err != nil {
		log.Println("保存数据库失败", err.Error())
	}
	ctx.Redirect(http.StatusMovedPermanently, "/admin/banner")
}

//func RootPath() string {
//	s, err := exec.LookPath(os.Args[0])
//	if err != nil {
//		log.Panicln("发生错误", err.Error())
//	}
//	i := strings.LastIndex(s, "\\")
//	path := s[0 : i+1]
//	return path
//}

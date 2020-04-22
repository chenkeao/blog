package controllers

import (
	"log"
	"net/http"
	"strconv"
	"wblog/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Game2048(ctx *gin.Context) {
	user, _ := ctx.Get(CONTEXT_USER_KEY)
	gamers, err := models.GetAllGamer()
	if err != nil {
		log.Println(err)
		return
	}

	ctx.HTML(http.StatusOK, "game/index.html", gin.H{
		"user":   user,
		"gamers": gamers,
	})
}

func GamerSave(ctx *gin.Context) {
	log.Println("开始保存")

	s := ctx.PostForm("score")
	score, err := strconv.Atoi(s)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("上传的成绩为", score)
	u, _ := ctx.Get(CONTEXT_USER_KEY)
	user, ok := u.(*models.User)
	if !ok {
		return
	}
	log.Println("当前用户为", user.Email)
	gamer := &models.GameBoard{}
	err = models.DB.Where("email = ?", user.Email).Find(gamer).Error
	if err == nil {
		if score > gamer.Score {
			log.Println("开始更新成绩")
			err = models.DB.Model(gamer).Update("score", score).Error
			if err != nil {
				log.Println(err)
			}
		}
		//表示查询成功但没有结果
	} else if err == gorm.ErrRecordNotFound {
		log.Println("数据库无此用户，新建成绩")
		gamer = &models.GameBoard{Score: score, Email: user.Email}
		_ = gamer.Create()
	}
	ctx.JSON(http.StatusOK, nil)
}

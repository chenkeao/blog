package controllers

import (
	"fmt"
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
		fmt.Println(err)
		return
	}

	ctx.HTML(http.StatusOK, "game/index.html", gin.H{
		"user":   user,
		"gamers": gamers,
	})
}

func GamerSave(ctx *gin.Context) {
	fmt.Println("开始保存")

	s := ctx.PostForm("score")
	score, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("上传的成绩为", score)
	u, _ := ctx.Get(CONTEXT_USER_KEY)
	user, ok := u.(*models.User)
	if !ok {
		return
	}
	fmt.Println("当前用户为", user.Email)
	err = models.DB.Where("email = ?", user.Email).Error
	if err == nil {
		fmt.Println("数据库有此用户")
		var gamer *models.GameBoard
		models.DB.Where("email = ?", user.Email).Find(gamer)
		if gamer.Score > score {
			fmt.Println("开始更新成绩")
			err = models.DB.Model(&models.GameBoard{
				Email: user.Email,
			}).Update("score", score).Error
			if err != nil {
				fmt.Println(err)
			}
		}
		//表示查询成功但没有结果
	} else if err == gorm.ErrRecordNotFound {
		fmt.Println("数据库无此用户，新建成绩")
		err = models.SaveGamer(user.Email, score)
	}
	ctx.Redirect(http.StatusMovedPermanently, "/game")
}

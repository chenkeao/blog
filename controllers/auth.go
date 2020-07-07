package controllers

import (
	"net/http"

	"github.com/chenkeao/blog/helpers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthGet(c *gin.Context) {
	authType := c.Param("authType")

	session := sessions.Default(c)
	uuid := helpers.UUID()
	session.Delete(SESSION_GITHUB_STATE)
	session.Set(SESSION_GITHUB_STATE, uuid)
	session.Save()

	authurl := "/signin"
	switch authType {
	//case "github":
	//	authurl = fmt.Sprintf(system.GetConfiguration().GithubAuthUrl, system.GetConfiguration().GithubClientId, uuid)
	//case "weibo":
	//case "qq":
	//case "wechat":
	//case "oschina":
	default:
	}
	c.Redirect(http.StatusFound, authurl)
}

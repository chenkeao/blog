package controllers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"wblog/helpers"
	"wblog/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type GithubUserInfo struct {
	AvatarPath        string      `json:"avatar_path"`
	Bio               interface{} `json:"bio"`
	Blog              string      `json:"blog"`
	Company           interface{} `json:"company"`
	CreatedAt         string      `json:"created_at"`
	Email             interface{} `json:"email"`
	EventsURL         string      `json:"events_url"`
	Followers         int         `json:"followers"`
	FollowersURL      string      `json:"followers_url"`
	Following         int         `json:"following"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	GravatarID        string      `json:"gravatar_id"`
	Hireable          interface{} `json:"hireable"`
	HTMLURL           string      `json:"html_url"`
	ID                int         `json:"id"`
	Location          interface{} `json:"location"`
	Login             string      `json:"login"`
	Name              interface{} `json:"name"`
	OrganizationsURL  string      `json:"organizations_url"`
	PublicGists       int         `json:"public_gists"`
	PublicRepos       int         `json:"public_repos"`
	ReceivedEventsURL string      `json:"received_events_url"`
	ReposURL          string      `json:"repos_url"`
	SiteAdmin         bool        `json:"site_admin"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	Type              string      `json:"type"`
	UpdatedAt         string      `json:"updated_at"`
	URL               string      `json:"url"`
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/signin.html", nil)
}

func SignupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/signup.html", nil)
}

func LogoutGet(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.Redirect(http.StatusSeeOther, "/signin")
}

func SignupPost(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	email := c.PostForm("email")
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	user := &models.User{
		Email:    email,
		NickName: nickname,
		Password: password,
		IsAdmin:  true,
	}
	if len(user.Email) == 0 || len(user.Password) == 0 {
		res["message"] = "邮箱和密码不能为空"
		return
	}
	user.Password = helpers.Md5(user.Email + user.Password)
	err = user.Insert()
	if err != nil {
		res["message"] = "邮箱已经注册"
		return
	}
	res["succeed"] = true
}

func SigninPost(c *gin.Context) {
	var (
		err  error
		user *models.User
	)
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		c.HTML(http.StatusOK, "auth/signin.html", gin.H{
			"message": "用户名和密码不能为空",
		})
		return
	}
	user, err = models.GetUserByUsername(email)
	if err != nil || user.Password != helpers.Md5(email+password) {
		c.HTML(http.StatusOK, "auth/signin.html", gin.H{
			"message": "用户名或密码无效",
		})
		return
	}
	if user.LockState {
		c.HTML(http.StatusOK, "auth/signin.html", gin.H{
			"message": "您的账户已经被锁定",
		})
		return
	}
	s := sessions.Default(c)
	s.Clear()
	s.Set(SESSION_KEY, user.ID)
	s.Save()
	//重定向到首页
	c.Redirect(http.StatusMovedPermanently, "/")
}

func ProfileGet(c *gin.Context) {
	sessionUser, exists := c.Get(CONTEXT_USER_KEY)
	if exists {
		c.HTML(http.StatusOK, "profile.html", gin.H{
			"user":     sessionUser,
			"comments": models.MustListUnreadComment(),
		})
	} else {
		c.HTML(http.StatusOK, "errors/error.html", gin.H{
			"message": "尚未登录",
		})
	}
}

func ProfileUpdate(c *gin.Context) {
	avatar, err := c.FormFile("avatar")
	nickName := c.PostForm("nickName")
	sessionUser, _ := c.Get(CONTEXT_USER_KEY)
	user, ok := sessionUser.(*models.User)
	if !ok {
		c.HTML(http.StatusOK, "/error.html", gin.H{
			"message": "个人信息更新失败",
		})
		return
	}
	if avatar != nil {
		path := "static/avatar"
		path = filepath.Join(path, user.Email)
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			c.HTML(http.StatusOK, "errors/error.html", gin.H{
				"message": err.Error(),
			})
			return
		}
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + avatar.Filename
		path = filepath.Join(path, fileName)
		err = c.SaveUploadedFile(avatar, path)
		if err != nil {
			c.HTML(http.StatusOK, "errors/error.html", gin.H{
				"message": err.Error(),
			})
			return
		}
		err = user.UpdateProfile(path, nickName)
		if err != nil {
			c.HTML(http.StatusOK, "errors/error.html", gin.H{
				"message": err.Error(),
			})
			return
		}
	} else {
		err = user.UpdateProfile("", nickName)
	}
	c.Redirect(http.StatusMovedPermanently, "/profile")
}

func BindEmail(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	email := c.PostForm("email")
	sessionUser, _ := c.Get(CONTEXT_USER_KEY)
	user, ok := sessionUser.(*models.User)
	if !ok {
		res["message"] = "server interval error"
		return
	}
	if len(user.Email) > 0 {
		res["message"] = "email have bound"
		return
	}
	_, err = models.GetUserByUsername(email)
	if err == nil {
		res["message"] = "email have be registered"
		return
	}
	err = user.UpdateEmail(email)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func UnbindEmail(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	sessionUser, _ := c.Get(CONTEXT_USER_KEY)
	user, ok := sessionUser.(*models.User)
	if !ok {
		res["message"] = "server interval error"
		return
	}
	if user.Email == "" {
		res["message"] = "email haven't bound"
		return
	}
	err = user.UpdateEmail("")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func UserIndex(c *gin.Context) {
	users, _ := models.ListUsers()
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "admin/user.html", gin.H{
		"users":    users,
		"user":     user,
		"comments": models.MustListUnreadComment(),
	})
}

func UserLock(c *gin.Context) {
	var (
		err  error
		_id  uint64
		res  = gin.H{}
		user *models.User
	)
	defer writeJSON(c, res)
	id := c.Param("id")
	_id, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	user, err = models.GetUser(uint(_id))
	if err != nil {
		res["message"] = err.Error()
		return
	}
	user.LockState = !user.LockState
	err = user.Lock()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

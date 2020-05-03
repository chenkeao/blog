package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/sessions"

	"path/filepath"

	"os"
	"strings"

	"github.com/cihub/seelog"
	_ "github.com/claudiu/gocron"

	"blog/controllers"
	"blog/helpers"
	"blog/models"
	"blog/system"

	"github.com/gin-gonic/gin"
)

func main() {
	configFilePath := flag.String("C", "conf/conf.yaml", "config file path")
	logConfigPath := flag.String("L", "conf/seelog.xml", "log config file path")
	flag.Parse()

	logger, err := seelog.LoggerFromConfigAsFile(*logConfigPath)
	if err != nil {
		log.Println("读取配置文件错误：", err.Error())
		seelog.Critical("err parsing seelog config file", err)
		return
	}
	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	if err := system.LoadConfiguration(*configFilePath); err != nil {
		log.Println("解析配置文件错误：", err.Error())
		seelog.Critical("err parsing config log file", err)
		return
	}

	db, err := models.InitDB()
	if err != nil {
		log.Println("数据库加载失败:", err.Error())
		seelog.Critical("err open databases", err)
		return
	}
	defer db.Close()

	//是否为调试模式
	if system.GetConfiguration().Mode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	setTemplate(router)
	setSessions(router)
	router.Use(SharedData())

	//Periodic tasks
	//gocron.Every(1).Day().Do(controllers.CreateXMLSitemap)
	//gocron.Every(7).Days().Do(controllers.Backup)
	//gocron.Start()

	router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	//router.StaticFS("/banner", http.Dir(controllers.RootPath()+"banner/"))

	router.NoRoute(controllers.Handle404)

	router.GET("/", controllers.IndexGet)
	router.GET("/index", controllers.IndexGet)
	router.GET("/rss", controllers.RssGet)
	router.GET("/about", controllers.AboutGet)

	if system.GetConfiguration().SignupEnabled {
		router.GET("/signup", controllers.SignupGet)
		router.POST("/signup", controllers.SignupPost)
	}

	router.GET("/signin", controllers.Login)
	router.POST("/signin", controllers.SigninPost)
	router.GET("/forgot_password", func(context *gin.Context) {
		context.HTML(http.StatusOK, "auth/forgot_password.html", nil)
	})
	router.POST("/forgot_password", controllers.ForgotPassword)
	router.GET("/logout", controllers.LogoutGet)

	router.GET("/captcha", controllers.CaptchaGet)

	visitor := router.Group("/visitor")
	visitor.Use(AuthRequired())
	{
		visitor.POST("/new_comment", controllers.CommentPost)
		visitor.POST("/comment/:id/delete", controllers.CommentDelete)
	}

	// subscriber
	router.GET("/subscribe", controllers.SubscribeGet)
	router.POST("/subscribe", controllers.Subscribe)
	router.GET("/active", controllers.ActiveSubscriber)
	router.GET("/unsubscribe", controllers.UnSubscribe)

	router.GET("/page/:id", controllers.PageGet)
	router.GET("/post/:id", controllers.PostGet)
	router.GET("/tag/:tag", controllers.TagGet)
	router.GET("/archives/:year/:month", controllers.ArchiveGet)

	router.GET("/link/:id", controllers.LinkGet)

	router.GET("/profile", controllers.ProfileGet)
	router.POST("/profile", controllers.ProfileUpdate)
	router.POST("/profile/email/bind", controllers.BindEmail)
	router.POST("/profile/email/unbind", controllers.UnbindEmail)

	authorized := router.Group("/admin")
	authorized.Use(AdminScopeRequired())
	{
		authorized.GET("/", controllers.AdminIndex)
		// index
		authorized.GET("/index", controllers.AdminIndex)

		// image upload
		authorized.POST("/upload", controllers.Upload)

		// page
		authorized.GET("/page", controllers.PageIndex)
		authorized.GET("/new_page", controllers.PageNew)
		authorized.POST("/new_page", controllers.PageCreate)
		authorized.GET("/page/:id/edit", controllers.PageEdit)
		authorized.POST("/page/:id/edit", controllers.PageUpdate)
		authorized.POST("/page/:id/publish", controllers.PagePublish)
		authorized.POST("/page/:id/delete", controllers.PageDelete)

		// post
		authorized.GET("/post", controllers.PostIndex)
		authorized.GET("/new_post", controllers.PostNew)
		authorized.POST("/new_post", controllers.PostCreate)
		authorized.GET("/post/:id/edit", controllers.PostEdit)
		authorized.POST("/post/:id/edit", controllers.PostUpdate)
		authorized.POST("/post/:id/publish", controllers.PostPublish)
		authorized.POST("/post/:id/delete", controllers.PostDelete)

		// tag
		authorized.POST("/new_tag", controllers.TagCreate)

		// user
		authorized.GET("/user", controllers.UserIndex)
		authorized.POST("/user/:id/lock", controllers.UserLock)

		//banner
		authorized.GET("banner", controllers.BannerIndex)
		authorized.POST("/new_banner", controllers.BannerNew)
		authorized.POST("/banner/:id/delete", controllers.BannerDelete)

		// subscriber
		authorized.GET("/subscriber", controllers.SubscriberIndex)
		authorized.POST("/subscriber", controllers.SubscriberPost)

		// link
		authorized.GET("/link", controllers.LinkIndex)
		authorized.POST("/new_link", controllers.LinkCreate)
		authorized.POST("/link/:id/edit", controllers.LinkUpdate)
		authorized.POST("/link/:id/delete", controllers.LinkDelete)

		// comment
		authorized.POST("/comment/:id", controllers.CommentRead)
		authorized.POST("/read_all", controllers.CommentReadAll)

		//备份，暂时不需要
		//authorized.POST("/backup", controllers.BackupPost)
		//authorized.POST("/restore", controllers.RestorePost)

		// mail
		authorized.POST("/new_mail", controllers.SendMail)
		authorized.POST("/new_batchmail", controllers.SendBatchMail)
	}

	game := router.Group("game")
	{
		game.GET("/", controllers.Game2048)
		game.POST("/save", AuthRequired(), controllers.GamerSave)
	}

	router.Run(system.GetConfiguration().Addr)
}

func setTemplate(engine *gin.Engine) {

	funcMap := template.FuncMap{
		"dateFormat": helpers.DateFormat,
		"substring":  helpers.Substring,
		"isOdd":      helpers.IsOdd,
		"isEven":     helpers.IsEven,
		"truncate":   helpers.Truncate,
		"add":        helpers.Add,
		"minus":      helpers.Minus,
		"listtag":    helpers.ListTag,
	}

	engine.SetFuncMap(funcMap)
	engine.LoadHTMLGlob(filepath.Join(getCurrentDirectory(), "./views/**/*"))
}

//setSessions initializes sessions & csrf middlewares
func setSessions(router *gin.Engine) {
	config := system.GetConfiguration()
	//https://github.com/gin-gonic/contrib/tree/master/sessions
	store := cookie.NewStore([]byte(config.SessionSecret))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"}) //Also set Secure: true if using SSL, you should though
	router.Use(sessions.Sessions("gin-session", store))
	//https://github.com/utrack/gin-csrf
	/*router.Use(csrf.Middleware(csrf.Options{
		Secret: config.SessionSecret,
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))*/
}

//+++++++++++++ middlewares +++++++++++++++++++++++

//SharedData fills in common data, such as user info, etc...
func SharedData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if uID := session.Get(controllers.SESSION_KEY); uID != nil {
			user, err := models.GetUser(uID)
			if err == nil {
				c.Set(controllers.CONTEXT_USER_KEY, user)
			}
		}
		if system.GetConfiguration().SignupEnabled {
			c.Set("SignupEnabled", true)
		}
		c.Next()
	}
}

//AdminScopeRequired 中间件
func AdminScopeRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get(controllers.CONTEXT_USER_KEY)
		if user != nil {
			if u, ok := user.(*models.User); ok && u.IsAdmin {
				c.Next()
				return
			}
		}
		seelog.Warnf("User not authorized to visit %s", c.Request.RequestURI)
		c.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": "权限不足，请使用管理员账号登录!",
		})
		c.Abort()
	}
}

//AuthRequired 中间件
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(controllers.CONTEXT_USER_KEY); user != nil {
			if _, ok := user.(*models.User); ok {
				c.Next()
				return
			}
		}
		seelog.Warnf("User not authorized to visit %s", c.Request.RequestURI)
		c.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden!",
		})
		c.Abort()
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		seelog.Critical(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//func getCurrentDirectory() string {
//	return ""
//}

package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"seckillsys/common"
	"seckillsys/fronted/controllers"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("//Users/linjunhan/project/seckillSys/fronted/views/**/*")
	router.Static("/Users/linjunhan/project/seckillSys/fronted/public", "/public")

	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Fatal("数据库连接失败！")
	}
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	uc := controllers.NewUserController("user", db)
	router.GET("/user/login", func(c *gin.Context) {
		uc.GetLogin(c)
	})
	router.GET("/user/register", func(c *gin.Context) {
		uc.GetRegister(c)
	})
	router.POST("/user/register", func(c *gin.Context) {
		uc.PostRegister(c)
	})

	router.Run(":34281")
}

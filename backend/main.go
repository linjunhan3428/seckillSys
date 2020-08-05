package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go/log"
	_ "net/http"
	"seckillsys/backend/web/controllers"
	"seckillsys/common"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("/Users/linjunhan/project/seckillSys/backend/web/views/**/*")
	router.Static("/assets", "/Users/linjunhan/project/seckillSys/backend/web/assets")

	//连接数据库
	db,err :=common.NewMysqlConn()
	if err !=nil {
		log.Error(err)
	}
	pc := controllers.NewProductController(db)

	router.GET("/product/all", func(c *gin.Context) {
		pc.GetAll(c)
	})

	router.Run(":3428")
}
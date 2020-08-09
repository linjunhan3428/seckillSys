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
	router.Static("/Users/linjunhan/project/seckillSys/backend/web/assets", "/assets")

	//连接数据库
	db, err := common.NewMysqlConn()
	defer db.Close()
	if err != nil {
		log.Error(err)
	}
	pc := controllers.NewProductController(db)
	oc := controllers.NewOrderController(db)

	productRouter := router.Group("/product")
	{
		productRouter.GET("/all", func(c *gin.Context) { pc.GetAll(c) })
		productRouter.GET("/delete", func(c *gin.Context) { pc.GetDelete(c) })
		productRouter.GET("/manager", func(c *gin.Context) { pc.GetManager(c) })
		productRouter.POST("/update", func(c *gin.Context) { pc.PostUpdate(c) })
		productRouter.GET("/add", func(c *gin.Context) { pc.GetAdd(c) })
		productRouter.POST("/add", func(c *gin.Context) { pc.PostAdd(c) })
	}

	orderRouter := router.Group("/order")
	{
		orderRouter.GET("/all", func(c *gin.Context) { oc.Get(c) })
	}

	router.Run(":3428")
}

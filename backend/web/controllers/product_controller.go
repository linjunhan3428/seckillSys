package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"seckillsys/repositories"
	"seckillsys/service"
)

type ProductController struct {
	ProductService service.IProductService
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{ProductService: service.NewProductService(repositories.NewProductManager("product", db))}
}

func (p *ProductController) GetAll(c *gin.Context) {

	productArray ,_ := p.ProductService.GetAllProduct()
	c.HTML(http.StatusOK, "product/view.html", gin.H{
		"productArray" : productArray,
	})
}

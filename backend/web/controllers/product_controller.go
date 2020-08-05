package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"seckillsys/datamodels"
	"seckillsys/repositories"
	"seckillsys/service"
	"strconv"
)

type ProductController struct {
	ProductService service.IProductService
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{ProductService: service.NewProductService(repositories.NewProductManager("product", db))}
}

func (p *ProductController) GetAll(c *gin.Context) {

	productArray, err := p.ProductService.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + err.Error() + "</b>"})
		return
	}
	c.HTML(http.StatusOK, "product/view.html", gin.H{
		"productArray": productArray,
	})
}

func (p *ProductController) GetDelete(c *gin.Context) {

	id := c.Query("id")
	idint, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + err.Error() + "</b>"})
		return
	}
	isSuccess := p.ProductService.DeleteProductByID(idint)
	if isSuccess == false {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "delete product failed!" + "</b>"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/product/all")
}

func (p *ProductController) GetManager(c *gin.Context) {

	idString := c.Query("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "商品id解析失败:" + err.Error() + "</b>"})
		return
	}

	product, err := p.ProductService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "获取商品信息失败:" + err.Error() + "</b>"})
		return
	}

	c.HTML(http.StatusOK, "product/manager.html", gin.H{
		"product": product,
	})
}

// 修改商品
func (p *ProductController) PostUpdate(c *gin.Context) {
	product := &datamodels.Product{}
	IDString := c.PostForm("ID")
	ProductNumString := c.PostForm("ProductNum")
	ProductImageString := c.PostForm("ProductImage")
	ProductUrlString := c.PostForm("ProductUrl")
	ProductNameString := c.PostForm("ProductName")

	ID, err := strconv.ParseInt(IDString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "商品id解析失败:" + err.Error() + "</b>"})
		return
	}
	ProductNum, err := strconv.ParseInt(ProductNumString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "商品数量解析失败:" + err.Error() + "</b>"})
		return
	}
	product.ID = ID
	product.ProductNum = ProductNum
	product.ProductImage = ProductImageString
	product.ProductUrl = ProductUrlString
	product.ProductName = ProductNameString

	err = p.ProductService.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "更新商品失败:" + err.Error() + "</b>"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/product/all")
}

func (p *ProductController) GetAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "product/add.html", nil)
}

func (p *ProductController) PostAdd(c *gin.Context) {
	product := &datamodels.Product{}
	ProductNumString := c.PostForm("ProductNum")
	ProductImageString := c.PostForm("ProductImage")
	ProductUrlString := c.PostForm("ProductUrl")
	ProductNameString := c.PostForm("ProductName")

	ProductNum, err := strconv.ParseInt(ProductNumString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "商品数量解析失败:" + err.Error() + "</b>"})
		return
	}
	product.ProductNum = ProductNum
	product.ProductImage = ProductImageString
	product.ProductUrl = ProductUrlString
	product.ProductName = ProductNameString

	_, err = p.ProductService.InsertProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + "新增商品失败:" + err.Error() + "</b>"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/product/all")
}

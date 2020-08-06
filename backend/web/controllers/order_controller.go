package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"seckillsys/service"
)

type OrderController struct {
	OrderService service.IOrderService
}

func NewOrderController(db *sql.DB) *OrderController {
	return &OrderController{OrderService: service.NewOrderService("`order`", db)}
}

func (o *OrderController) Get(c *gin.Context) {

	orderArray, err := o.OrderService.GetAllOrderInfo()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + err.Error() + "</b>"})
		return
	}

	c.HTML(http.StatusOK, "order/view.html", gin.H{"order": orderArray})
}

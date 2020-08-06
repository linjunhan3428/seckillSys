package repositories

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"seckillsys/common"
	"seckillsys/datamodels"
	"testing"
)

var db, _ = common.NewMysqlConn()
var or = NewOrderMangerRepository("`order`", db)

func TestSelectAll(t *testing.T) {
	orderList, err := or.SelectAll()
	assert.NoError(t, err)

	for _, v := range orderList {
		fmt.Println(v.ID, v.UserId, v.ProductId, v.OrderStatus)
	}
}

func TestInsertOrder(t *testing.T) {

	productId, err := or.Insert(&datamodels.Order{
		UserId:      12334,
		ProductId:   129,
		OrderStatus: datamodels.OrderWait,
	})
	assert.NoError(t, err)
	fmt.Println(productId)
}

func TestUpdateOrder(t *testing.T) {

	err := or.Update(&datamodels.Order{ID: 12345679, UserId: 12334, ProductId: 125, OrderStatus: 3})
	assert.NoError(t, err)
}

func TestSelectByKey(t *testing.T) {
	order, err := or.SelectByKey(12345679)
	assert.NoError(t, err)
	fmt.Println(order.ID, order.UserId, order.ProductId, order.OrderStatus)
}

func TestSelectAllWithInfo(t *testing.T) {
	info, err := or.SelectAllWithInfo()
	assert.NoError(t, err)

	for k, v := range info {
		fmt.Println(k, ":", v)
	}
}

func TestDelete(t *testing.T) {
	b := or.Delete(12345678)
	assert.Equal(t, b, true)
}

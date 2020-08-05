package repositories

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"seckillsys/common"
	"seckillsys/datamodels"
	"testing"
)

func TestInsert(t *testing.T) {

	db, err := common.NewMysqlConn()
	defer db.Close()
	assert.NoError(t, err)

	manager := NewProductManager("product", db)

	//num, err := manager.Insert(&datamodels.Product{122, "mac", 100, "xxxxzzz", "httpd"})
	//assert.NoError(t, err)
	//assert.Equal(t, int64(126), num)

	products, err := manager.SelectAll()
	assert.NoError(t, err)
	for _, v := range products {
		fmt.Println(v.ID, v.ProductImage, v.ProductName, v.ProductNum, v.ProductUrl)
	}

	err = manager.Update(&datamodels.Product{123, "mac", 100, "aaajjjj", "httpd"})
	assert.NoError(t, err)

	product, err := manager.SelectByKey(123)
	assert.NoError(t, err)
	fmt.Println(product.ProductName, product.ProductUrl, product.ProductNum, product.ProductImage)

	//b := manager.Delete(126)
	//assert.Equal(t, true, b)
}

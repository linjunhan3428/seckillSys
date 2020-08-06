package repositories

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"seckillsys/datamodels"
	"testing"
)

var ur = NewUserRepository("user", db)

func TestInsertUser(t *testing.T) {
	userId, err := ur.Insert(&datamodels.User{ID: 3, NickName: "dasdas", UserName: "dasdasd", HashPassword: string([]byte{})})
	assert.NoError(t, err)
	fmt.Println(userId)
}

func TestSelect(t *testing.T) {
	user, err := ur.Select("dasdasd")
	assert.NoError(t, err)
	fmt.Println(user.ID, user.NickName, user.UserName, user.HashPassword)
}

func TestSelectByID(t *testing.T) {
	user, err := ur.(*UserManagerRepository).SelectByID(4)
	assert.NoError(t, err)
	fmt.Println(user.ID, user.NickName, user.UserName, user.HashPassword)
}

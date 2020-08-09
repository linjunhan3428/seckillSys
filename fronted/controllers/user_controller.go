package controllers

import (
	"database/sql"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"seckillsys/datamodels"
	"seckillsys/repositories"
	"seckillsys/service"
)

type UserController struct {
	Service service.IUserService
	Session *sessions.Session
}

func NewUserController(table string, db *sql.DB) *UserController {
	return &UserController{
		Service: service.NewService(repositories.NewUserRepository(table, db)),
	}
}

func (u *UserController) GetRegister(c *gin.Context)  {
	c.HTML(http.StatusOK, "user/register.html", nil)
}

func (u *UserController) GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login.html", nil)
}

func (u *UserController) PostRegister(c *gin.Context) {
	var (
		nickName = c.PostForm("nickName")
		userName = c.PostForm("userName")
		password = c.PostForm("password")
	)
	//ozzo-validation
	user := &datamodels.User{
		UserName:     userName,
		NickName:     nickName,
		HashPassword: password,
	}

	_, err := u.Service.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"html": "<b>" + err.Error() + "</b>"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/user/login")
	return
}

func (u *UserController) PostLogin(c *gin.Context) {
	//1.获取用户提交的表单信息
	var (
		userName = c.PostForm("userName")
		password = c.PostForm("password")
	)
	//2、验证账号密码正确
	_, isOk := u.Service.IsPwdSuccess(userName, password)
	if !isOk {
		c.Redirect(http.StatusMovedPermanently, "/user/login")
		return
	}

	//3、写入用户ID到cookie中

	c.HTML(http.StatusOK, "/product", nil)

}



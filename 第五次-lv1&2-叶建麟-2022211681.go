package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var users map[string]string = make(map[string]string)

type userinfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func Regestar() gin.HandlerFunc {

	return func(context *gin.Context) {
		fmt.Println("请注册")
	}

}
func Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("请输入用户名和密码")
	}

}

func regestar(c *gin.Context) {
	var u userinfo
	c.ShouldBind(&u)
	u.Password = c.PostForm("username")
	u.Password = c.PostForm("password")
	users[u.Username] = u.Password
	fmt.Println("注册成功\n用户名：" + u.Username + "密码：" + u.Password)
	c.JSON(200, "注册成功")
}

func log(c *gin.Context) {
	var l userinfo
	c.ShouldBind(&l)
	l.Username = c.PostForm("username")
	l.Password = c.PostForm("password")
	if users[l.Username] == l.Password {
		fmt.Println("登陆成功")
		c.JSON(200, "登陆成功")
		_, err := c.Cookie("usercookie")
		if err != nil {
			c.SetCookie("usercookie", "123456",
				3600, ",/", "localhost",
				false, true)
		}
	} else {
		fmt.Println("用户名或密码错误")
	}
}

func main() {
	r := gin.Default()
	usr := r.Group("/user")
	{
		usr.POST("/regestar", Regestar(), regestar)
		usr.POST("/login", Login(), log)
	}
	r.Run()
}

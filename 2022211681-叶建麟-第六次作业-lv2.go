package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

var users map[string]string = make(map[string]string)

type userinfo struct {
	Username     string `json:"username" form:"username"`
	Password     string `json:"password" form:"password"`
	SafeQuestion string `json:"safeQuestion" form:"safeQuestion"`
	SafeAnswer   string `json:"safeAnswer" form:"safeAnswer"`
}

func Regestar() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("请注册")
	}
}
func Login() gin.HandlerFunc {
	return func(context *gin.Context) {
	}

}

func regestar(c *gin.Context) {
	var u userinfo
	c.ShouldBind(&u)
	users[u.Username] = u.Password
	usersSql(u)
	fmt.Println("注册成功\n用户名：" + u.Username + "密码：" + u.Password)
	c.JSON(200, "注册成功")
}

func Log(c *gin.Context) {
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
		var u userinfo
		fmt.Println("用户名或密码错误")
		db, err := sql.Open("mysql", "root:277187@tcp(127.0.0.1:3306)/users")
		err = db.Ping()
		if err != nil {
			return
		}
		row := db.QueryRow("select * from regestar where username=?", l.Username)
		err = row.Scan(&u.Username, &u.Password, &u.SafeQuestion, &u.SafeAnswer)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("请回答密保问题：" + u.SafeQuestion)
		fmt.Println(l.SafeAnswer)
		if strings.EqualFold(l.SafeAnswer, u.SafeAnswer) {
			fmt.Println("登陆成功,您的密码是：" + l.Password + "下次不要忘记了哟")
			c.JSON(200, "登陆成功")
		} else {
			fmt.Println("登录失败")
		}
	}
}

func usersSql(u userinfo) {
	db, err := sql.Open("mysql", "root:277187@tcp(127.0.0.1:3306)/users")
	err = db.Ping()
	if err != nil {
		return
	}
	result, err := db.Exec("insert into regestar (username,password,safequestion,safeanswer) values (?,?,?,?)", u.Username, u.Password, u.SafeQuestion, u.SafeAnswer)
	if err != nil {
		log.Println(err)
		return
	}
	result.LastInsertId()
}

func main() {

	r := gin.Default()
	usr := r.Group("/user")
	{
		usr.POST("/regestar", Regestar(), regestar)
		usr.POST("/login", Login(), Log)
	}
	r.Run()
}

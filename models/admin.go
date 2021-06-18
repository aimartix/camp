package models

import (
	"camp_b/tool"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
}

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var ad Admin
	db.Exec("insert into admin_log (username,ip) values ((select realname from admin where username = ?),?)",username,c.ClientIP())
	db.Table("admin").Select("username,password").Scan(&ad)
	if username == ad.Username && tool.GetMD5Encode(password) == ad.Password {
		c.JSON(200,map[string]interface{}{
			"info" : "登录成功",
			"status" : 1,
		})
	}else {
		c.JSON(200,map[string]interface{}{
			"info" : "账号或密码错误",
			"status" : 0,
		})
	}

}

func ChangeP(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if db.Exec("update admin set password = ? where username = ?",tool.GetMD5Encode(password),username).RowsAffected == 0 {
		c.JSON(200,map[string]interface{}{
			"info" : "修改失败",
			"status" : 0,
		})
	}else {
		c.JSON(200,map[string]interface{}{
			"info" : "修改成功",
			"status" : 1,
		})
	}
}
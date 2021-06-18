package models

import (
	"camp_b/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InList(c *gin.Context)  {
	page, _ := strconv.Atoi(c.PostForm("page"))
	pagesize, _ := strconv.Atoi(c.PostForm("pagesize"))
	title := c.PostForm("title")
	age := c.PostForm("age")
	var count Count
	var inlist []InformationList
	if title==""&&age==""{
		if db.Raw("select id,title,age,address,status from information").Scan(&inlist).Limit(pagesize).Offset((page-1)*pagesize).RowsAffected==0{
			c.JSON(200,map[string]interface{}{
				"info" : "数据为空",
				"status" : 0,
			})
		}else {
			db.Raw("select count(1) as Total from information").Scan(&count)
			c.JSON(200,map[string]interface{}{
				"info" : "成功",
				"status" : 1,
				"data" : inlist,
				"total":count.Total,
			})
		}
	}else if title!=""&&age==""{
		if db.Raw("select id,title,age,address,status from information where title like?","%"+title+"%").Scan(&inlist).Limit(pagesize).Offset((page-1)*pagesize).RowsAffected==0{
			c.JSON(200,map[string]interface{}{
				"info" : "数据为空",
				"status" : 0,
			})
		}else {
			db.Raw("select count(1) as Total from information where title like?","%"+title+"%").Scan(&count)
			c.JSON(200,map[string]interface{}{
				"info" : "成功",
				"status" : 1,
				"data" : inlist,
				"total":count.Total,
			})
		}
	}else if title==""&&age!=""{
		if db.Raw("select id,title,age,address,status from information where age like?","%"+age+"%").Scan(&inlist).Limit(pagesize).Offset((page-1)*pagesize).RowsAffected==0{
			c.JSON(200,map[string]interface{}{
				"info" : "数据为空",
				"status" : 0,
			})
		}else {
			db.Raw("select count(1) as Total from information where age like?","%"+age+"%").Scan(&count)
			c.JSON(200,map[string]interface{}{
				"info" : "成功",
				"status" : 1,
				"data" : inlist,
				"total":count.Total,
			})
		}
	}else if title!=""&&age!=""{
		if db.Raw("select id,title,age,address,status from information where title like? and age like?","%"+title+"%","%"+age+"%").Scan(&inlist).Limit(pagesize).Offset((page-1)*pagesize).RowsAffected==0{
			c.JSON(200,map[string]interface{}{
				"info" : "数据为空",
				"status" : 0,
			})
		}else {
			db.Raw("select count(1) as Total from information where title like? and age like?","%"+title+"%","%"+age+"%").Scan(&count)
			c.JSON(200,map[string]interface{}{
				"info" : "成功",
				"status" : 1,
				"data" : inlist,
				"total":count.Total,
			})
		}
	}

}

func InDetail(c *gin.Context)  {
	id :=c.PostForm("id")
	var indetail InformationDetail
	db.Raw("select * from information where id=?",id).Scan(&indetail)
	c.JSON(200,map[string]interface{}{
		"info" : "数据为空",
		"status" : 1,
		"data" : indetail,
	})
}

func AddIn(c *gin.Context)  {
	title :=c.PostForm("title")
	age := c.PostForm("age")
	r_status := c.PostForm("r_status")
	address := c.PostForm("address")
	period := c.PostForm("period")
	price := c.PostForm("price")
	cover := c.PostForm("cover")
	content := c.PostForm("content")
	status := c.PostForm("status")

	if db.Exec("insert into information (title,age,r_status,address,period,price,cover,content,status) values (?,?,?,?,?,?,?,?,?)",title,age,r_status,address,period,price,tool.Base64(title,cover),content,status).RowsAffected==0{
		c.JSON(200,map[string]interface{}{
			"info" : "添加失败",
			"status" : 0,
		})
	}else {
		c.JSON(200,map[string]interface{}{
			"info" : "添加成功",
			"status" : 1,
		})
	}
}

func EditIn(c *gin.Context)  {
	title :=c.PostForm("title")
	age := c.PostForm("age")
	r_status := c.PostForm("r_status")
	address := c.PostForm("address")
	period := c.PostForm("period")
	price := c.PostForm("price")
	cover := c.PostForm("cover")
	content := c.PostForm("content")
	status := c.PostForm("status")
	if cover !=""{
		if db.Exec("update information set title=?,age=?,r_status=?,address=?,period=?,price=?,cover=?,content=?,status=?",title,age,r_status,address,period,price,tool.Base64(title,cover),content,status).RowsAffected==0{
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
	}else {
		if db.Exec("update information set title=?,age=?,r_status=?,address=?,period=?,price=?,content=?,status=?",title,age,r_status,address,period,price,content,status).RowsAffected==0{
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
}

func DeleteIn(c *gin.Context)  {
	id := c.PostForm("id")
	if db.Exec("delete from information where id=?",id).RowsAffected==0{
		c.JSON(200,map[string]interface{}{
			"info" : "删除失败",
			"status" : 0,
		})
	}else {
		c.JSON(200,map[string]interface{}{
			"info" : "删除成功",
			"status" : 1,
		})
	}
}

func OperateIn(c *gin.Context)  {
	id :=c.PostForm("id")
	status := c.PostForm("status")
	if db.Exec("update information set status=? where id=?",status,id).RowsAffected==0{
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


type Count struct {
	Total   int
}
type InformationList struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Age       int     `json:"age"`
	Address   string     `json:"address"`
	Status    int        `json:"status"`
}




type InformationDetail struct {
	Id        string  `json:"id"`
	Title     string  `json:"title"`
	Age       int     `json:"age"`
	RStatus   string     `json:"r_status"`
	Address   string     `json:"address"`
	Period    string     `json:"period"`
	Price     float64    `json:"price"`
	Cover     string     `json:"cover"`
	Content   string     `json:"content"`
	Status    int        `json:"status"`
}

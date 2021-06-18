package main

import (
	"camp_b/middleware"
	"camp_b/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Println("谷志诚后台")
	r := gin.Default()
	r.Use(middleware.Cors())
	routers.Admin(r)//登录
	routers.Information(r)
	if err := r.Run(":100"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

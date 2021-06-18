package routers

import (
	"camp_b/models"
	"github.com/gin-gonic/gin"
)

func Information(e *gin.Engine)  {
	e.POST("/camp_b/addin",models.AddIn)
	e.POST("/camp_b/deletein",models.DeleteIn)
	e.POST("/camp_b/editin",models.EditIn)
	e.POST("/camp_b/operatein",models.OperateIn)
	e.POST("/camp_b/inlist",models.InList)
	e.POST("/camp_b/indetail",models.InDetail)
}

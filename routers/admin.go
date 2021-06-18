package routers

import (
	"camp_b/models"
	"github.com/gin-gonic/gin"
)

func Admin(e *gin.Engine)  {
	e.POST("/camp_b/login",models.Login)
	e.POST("/camp_b/changep",models.ChangeP)
}

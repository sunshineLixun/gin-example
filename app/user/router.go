package user

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/user", QueryHandler)
	e.POST("/user", CreateHandler)
	e.DELETE("/user", DeleteHandler)
	e.PUT("/user", UpdateHandler)
}
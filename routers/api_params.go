package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupApiRouter() {
	r := gin.Default()
	r.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	err := r.Run()
	if err != nil {
		return
	}
}

package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() {
	r := gin.Default()
	r.GET("/user", func(context *gin.Context) {
		name := context.DefaultQuery("name", "defaultValue")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	err := r.Run()
	if err != nil {
		return
	}
}

package main

import (
	"fmt"
	"gin-example/app/user"
	"gin-example/routers"
)

//var db = make(map[string]string)
//
//func setupRouter() *gin.Engine {
//	// Disable Console Color
//	// gin.DisableConsoleColor()
//	r := gin.Default()
//
//	// Ping test
//	r.GET("/ping", func(c *gin.Context) {
//		c.String(http.StatusOK, "pong")
//	})
//
//	// Get user value
//	r.GET("/user/:name", func(c *gin.Context) {
//		user := c.Params.ByName("name")
//		value, ok := db[user]
//		if ok {
//			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
//		} else {
//			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
//		}
//	})
//
//	// Authorized group (uses gin.BasicAuth() middleware)
//	// Same than:
//	// authorized := r.Group("/")
//	// authorized.Use(gin.BasicAuth(gin.Credentials{
//	//	  "foo":  "bar",
//	//	  "manu": "123",
//	//}))
//	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
//		"foo":  "bar", // user:foo password:bar
//		"manu": "123", // user:manu password:123
//	}))
//
//	/* example curl for /admin with basicAuth header
//	   Zm9vOmJhcg== is base64("foo:bar")
//
//		curl -X POST \
//	  	http://localhost:8080/admin \
//	  	-H 'authorization: Basic Zm9vOmJhcg==' \
//	  	-H 'content-type: application/json' \
//	  	-d '{"value":"bar"}'
//	*/
//	authorized.POST("admin", func(c *gin.Context) {
//		user := c.MustGet(gin.AuthUserKey).(string)
//
//		// Parse JSON
//		var json struct {
//			Value string `json:"value" binding:"required"`
//		}
//
//		if c.Bind(&json) == nil {
//			db[user] = json.Value
//			c.JSON(http.StatusOK, gin.H{"status": "ok"})
//		}
//	})
//
//	return r
//}

func main() {
	//r := setupRouter()
	//// Listen and Server in 0.0.0.0:8080
	//err := r.Run(":8080")
	//if err != nil {
	//	return
	//}
	// routers.SetupRouter()
	// routers.SetupApiRouter()

		// 加载多个APP的路由配置
		routers.Include(user.Routers)
		// 初始化路由
		r := routers.Init()
		if err := r.Run(); err != nil {
			fmt.Println("startup service failed, err:%v\n", err)
		}
}

package user

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)


type Part struct {
	ID    int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Name  string // json序列化是默认使用字段名作为key
	Count int
}

func getJSONData() []*Part {
	parts := make([]*Part, 0, 200)
	for i := 0; i < 10; i++ {
		part := &Part{
			Name:  "it",
			Count: 10,
			ID:    i,
		}
		parts = append(parts, part)
	}
	return parts
}

func QueryHandler(context *gin.Context) {
	data, err := json.Marshal(getJSONData())
	if err != nil {
		context.JSON(400, gin.H{
			"list": "error",
		})
	}
	context.JSON(200, gin.H{
		"list": data,
		"message": "GET",
	})
}

func CreateHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "create",
	})
}

func DeleteHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "delete",
	})
}


func UpdateHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "update",
	})
}
package helpers

import "github.com/gin-gonic/gin"

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

func JSON(c *gin.Context, status int, data interface{}) {
	resp := Response{
		Status: status,
		Result: data,
	}
	c.JSON(status, resp)
}

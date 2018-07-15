package handler

import (
	"github.com/gin-gonic/gin"
	"go-restapp-demo/pkg/errno"
	"net/http"
)

type Response struct {
	Code        int            `json:"code"`
	Message     string         `json:"message"`
	Data        interface{}    `json:"data"`
}

func SendRespnose(c *gin.Context, err error, data interface{}) {
	// 解析错误，没有错误code为0
	code,message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		code,
		message,
		data,
	})
}
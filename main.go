package main

import (
     "github.com/gin-gonic/gin"
     "go-restapp-demo/router"
	"log"
	"net/http"
)

func main() {

     g := gin.New()

     //自定义中间件
     middlewares := []gin.HandlerFunc{}

     // 注册路由
     router.Load(
     	 g,
		 middlewares...,
     )

     log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
     log.Printf(http.ListenAndServe(":8080", g).Error())
}




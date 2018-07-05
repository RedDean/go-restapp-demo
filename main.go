package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-restapp-demo/router"
	"log"
	"net/http"
	"time"
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

	//服务检查协程
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}

func pingServer() error {

	for i := 0; i < 3; i++ {
		resp, err := http.Get("http://127.0.0.1:8080/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Println("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("server cannot provide service")
}

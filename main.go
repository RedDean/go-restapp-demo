package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go-restapp-demo/config"
	"go-restapp-demo/router"
	"log"
	"net/http"
	"time"

	"go-restapp-demo/model"
)

var (
	cfg = pflag.StringP("config", "c", "", "app config file path")
)

func main() {

	pflag.Parse()
	// 加载配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	/*
		for {
			fmt.Println(viper.GetString("runmode"))
			time.Sleep(4*time.Second)
		}
	*/

	// 设置gin模式: gin有三种模式debug, release, test
	gin.SetMode(viper.GetString("runmode"))

	// 初始化数据库连接
	model.DB.Init()

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

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

func pingServer() error {

	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Println("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("server cannot provide service")
}

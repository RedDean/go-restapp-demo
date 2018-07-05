package router

import (
	"github.com/gin-gonic/gin"
	"go-restapp-demo/router/middleware"
	"net/http"
)

/**
 * description: 注册路由及中间件
 */

func Load (g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	g.Use(gin.Recovery())
    g.Use(middleware.Nocache)
    g.Use(middleware.Options)
    g.Use(middleware.Secure)
    //自定义中间件
    g.Use(mw...)

    // 404处理
    g.NoRoute(func(context *gin.Context) {
		context.String(http.StatusNotFound, "The incorrect API route.")
	})

    // 注册路由

    // sd分组：检测服务状态
    svcd := g.Group("/sd")
    {
		svcd.GET("/health", sd.HealthCheck)  // 健康检查
		svcd.GET("/disk", sd.DiskCheck)    // 硬盘情框
		svcd.GET("/cpu", sd.CPUCheck)      // CPU检查
		svcd.GET("/ram", sd.RAMCheck)      // 内存检查
	}

    return g
}

package main

import "github.com/gin-gonic/gin"

func main() {
	//使用默认中间件（logger 和 recovery 中间件）创建 gin 路由
	router := gin.Default()

	//router.GET("/someGet", getting)
	//router.POST("/somePost", posting)
	//router.PUT("/somePut", putting)
	//router.DELETE("/someDelete", deleting)
	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// 默认在 8080 端口启动服务，除非定义了一个 PORT 的环境变量。.
	router.Run()
	//  router.Run(":3000") hardcode 端口号
}

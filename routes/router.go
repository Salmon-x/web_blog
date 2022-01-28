package routes

import (
	"blog/middleware"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

//func createMyRender() multitemplate.Renderer {
//	p := multitemplate.NewRenderer()
//	p.AddFromFiles("admin", "web/admin/dist/index.html")
//	p.AddFromFiles("front", "web/front/dist/index.html")
//	return p
//}

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	// 初始化总路由
	DefRouter := r.Group("")
	{
		AdminRouterInit(DefRouter)
		V1RouterInit(DefRouter)
	}
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin","static/admin")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	_ = r.Run(utils.HttpPort)
}
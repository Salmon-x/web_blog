package routes

import (
	v1 "blog/api/v1"
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

	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin","static/admin")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	Auth := r.Group("api/admin")
	Auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		Auth.PUT("user/:id/", v1.EditUser)
		Auth.DELETE("user/:id/", v1.DeleteUser)
		// Category模块的路由接口
		Auth.POST("category/", v1.AddCategory)
		Auth.PUT("category/:id/", v1.EditCategory)
		Auth.DELETE("category/:id/", v1.DeleteCategory)
		// Article模块的路由接口
		Auth.POST("article/", v1.AddArticle)
		Auth.PUT("article/:id/", v1.EditArticle)
		Auth.DELETE("article/:id/", v1.DeleteArticle)
		Auth.POST("file/", v1.UploadFile)
		Auth.DELETE("file/", v1.DelFile)
		Auth.PUT("editpass/:id/", v1.AdminEditPass)
	}

	router := r.Group("api/v1")
	{
		router.GET("user/", v1.GetUsers)
		router.GET("/user/:id/", v1.GetUserInfo)
		router.GET("file/", v1.GetFile)
		router.POST("user/", v1.AddUser)

		router.POST("login/", v1.Login)

		router.GET("category/", v1.GetCategorys)
		router.GET("category/:id/", v1.GetCateInfo)

		router.GET("article/", v1.GetArticles)
		router.GET("article/:id/", v1.GetArticleInfo)
		router.GET("catelist/:id/", v1.GetCateArt)
	}


	_ = r.Run(utils.HttpPort)

}
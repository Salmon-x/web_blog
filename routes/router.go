package routes

import (
	v1 "blog/api/v1"
	"blog/middleware"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	Auth := r.Group("")
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
	}

	router := r.Group("api/v1")
	{
		router.GET("user/", v1.GetUsers)
		Auth.POST("user/", v1.AddUser)
		router.POST("login/", v1.Login)
		router.GET("category/", v1.GetCategorys)
		router.GET("category/:id/", v1.GetCateInfo)
		router.GET("article/", v1.GetArticles)
		router.GET("article/:id/", v1.GetArticleInfo)
		router.GET("catelist/:id/", v1.GetCateArt)
	}


	_ = r.Run(utils.HttpPort)

}
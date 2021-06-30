package routes

import (
	v1 "blog/api/v1"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		// User模块的路由接口
		router.POST("user/", v1.AddUser)
		router.GET("user/", v1.GetUsers)
		router.PUT("user/:id/", v1.EditUser)
		router.DELETE("user/:id/", v1.DeleteUser)

		// Category模块的路由接口
		router.POST("category/", v1.AddCategory)
		router.GET("category/", v1.GetCategorys)
		router.GET("category/:id/", v1.GetCateInfo)
		router.PUT("category/:id/", v1.EditCategory)
		router.DELETE("category/:id/", v1.DeleteCategory)

		// Article模块的路由接口
		router.POST("article/", v1.AddArticle)
		router.GET("article/", v1.GetArticles)
		router.GET("article/:id/", v1.GetArticleInfo)
		router.GET("catelist/:id/", v1.GetCateArt)
		router.PUT("article/:id/", v1.EditArticle)
		router.DELETE("article/:id/", v1.DeleteArticle)
	}

	_ = r.Run(utils.HttpPort)

}
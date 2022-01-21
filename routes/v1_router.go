package routes

import (
	"blog/api"
	"github.com/gin-gonic/gin"
)



func V1RouterInit(r *gin.RouterGroup)  {
	router := r.Group("api/v1")
	var ApiV1Router = api.ApiGroupApp.V1Group
	{
		router.GET("user/", ApiV1Router.GetUsers)
		router.GET("/user/:id/", ApiV1Router.GetUserInfo)
		router.GET("file/", ApiV1Router.GetFile)
		router.POST("user/", ApiV1Router.AddUser)
		router.POST("login/", ApiV1Router.Login)

		router.GET("category/", ApiV1Router.GetCategorys)
		router.GET("wks/", ApiV1Router.GetWks)
		router.GET("category/:id/", ApiV1Router.GetCateInfo)
		router.GET("wks/:id/", ApiV1Router.GetWksInfo)


		router.GET("article/", ApiV1Router.GetArticles)
		router.GET("article/:id/", ApiV1Router.GetArticleInfo)
		router.GET("frontarticle/:id/", ApiV1Router.FrontArticleInfo)
		router.GET("catelist/:id/", ApiV1Router.GetCateArt)

		router.GET("/captcha/", ApiV1Router.GetCaptCha)
		router.POST("/captcha/", ApiV1Router.VerifyCaptCha)
	}
}
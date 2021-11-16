package routes

import (
	"blog/api"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.RouterGroup) {

	Auth := r.Group("api/admin")
	var ApiAdmiRouter = api.ApiGroupApp.V1Group
	Auth.Use(middleware.JwtToken())
	{
		// User模块的路由接口
		Auth.PUT("user/:id/", ApiAdmiRouter.EditUser)
		Auth.DELETE("user/:id/", ApiAdmiRouter.DeleteUser)
		// Category模块的路由接口
		Auth.POST("category/", ApiAdmiRouter.AddCategory)
		Auth.PUT("category/:id/", ApiAdmiRouter.EditCategory)
		Auth.DELETE("category/:id/", ApiAdmiRouter.DeleteCategory)
		// Article模块的路由接口
		Auth.POST("article/", ApiAdmiRouter.AddArticle)
		Auth.PUT("article/:id/", ApiAdmiRouter.EditArticle)
		Auth.DELETE("article/:id/", ApiAdmiRouter.DeleteArticle)
		Auth.POST("file/", ApiAdmiRouter.UploadFile)
		Auth.DELETE("file/", ApiAdmiRouter.DelFile)
		Auth.PUT("editpass/:id/", ApiAdmiRouter.AdminEditPass)
		// Category模块的路由接口
		Auth.POST("wks/", ApiAdmiRouter.AddWks)
		Auth.PUT("wks/:id/", ApiAdmiRouter.EditWks)
		Auth.DELETE("wks/:id/", ApiAdmiRouter.DeleteWks)
	}
}

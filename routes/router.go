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
		router.POST("user/add/", v1.AddUser)
		router.GET("users/", v1.GetUsers)
		router.PUT("users/:id/", v1.EditUser)
		router.DELETE("users/:id/", v1.DeleteUser)

		// Category模块的路由接口

		// Article模块的路由接口
	}

	r.Run(utils.HttpPort)

}
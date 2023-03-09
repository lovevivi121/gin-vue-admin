package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UsersRouter struct {
}

// InitUsersRouter 初始化 Users 路由信息
func (s *UsersRouter) InitUsersRouter(Router *gin.RouterGroup) {
	usersRouter := Router.Group("users").Use(middleware.OperationRecord())
	usersRouterWithoutRecord := Router.Group("users")
	var usersApi = v1.ApiGroupApp.UserApiGroup.UsersApi
	{
		usersRouter.POST("createUsers", usersApi.CreateUsers)   // 新建Users
		usersRouter.DELETE("deleteUsers", usersApi.DeleteUsers) // 删除Users
		usersRouter.DELETE("deleteUsersByIds", usersApi.DeleteUsersByIds) // 批量删除Users
		usersRouter.PUT("updateUsers", usersApi.UpdateUsers)    // 更新Users
	}
	{
		usersRouterWithoutRecord.GET("findUsers", usersApi.FindUsers)        // 根据ID获取Users
		usersRouterWithoutRecord.GET("getUsersList", usersApi.GetUsersList)  // 获取Users列表
	}
}

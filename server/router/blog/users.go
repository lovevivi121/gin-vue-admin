package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// InitUserRouter 初始化 User 路由信息
func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("blog/user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := Router.Group("blog/user")
	var userApi = v1.ApiGroupApp.BlogApiGroup.UserApi
	{
		userRouter.POST("createUser", userApi.CreateUser)   // 新建User
		userRouter.DELETE("deleteUser", userApi.DeleteUser) // 删除User
		userRouter.DELETE("deleteUserByIds", userApi.DeleteUserByIds) // 批量删除User
		userRouter.PUT("updateUser", userApi.UpdateUser)    // 更新User
	}
	{
		userRouterWithoutRecord.GET("findUser", userApi.FindUser)        // 根据ID获取User
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList)  // 获取User列表
	}
}

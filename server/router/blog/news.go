package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type NewsRouter struct {
}

// InitNewsRouter 初始化 News 路由信息
func (s *NewsRouter) InitNewsRouter(Router *gin.RouterGroup) {
	newsRouter := Router.Group("news").Use(middleware.OperationRecord())
	newsRouterWithoutRecord := Router.Group("news")
	var newsApi = v1.ApiGroupApp.BlogApiGroup.NewsApi
	{
		newsRouter.POST("createNews", newsApi.CreateNews)   // 新建News
		newsRouter.DELETE("deleteNews", newsApi.DeleteNews) // 删除News
		newsRouter.DELETE("deleteNewsByIds", newsApi.DeleteNewsByIds) // 批量删除News
		newsRouter.PUT("updateNews", newsApi.UpdateNews)    // 更新News
	}
	{
		newsRouterWithoutRecord.GET("findNews", newsApi.FindNews)        // 根据ID获取News
		newsRouterWithoutRecord.GET("getNewsList", newsApi.GetNewsList)  // 获取News列表
	}
}

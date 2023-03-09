package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/blog"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type NewsApi struct {
}

var newsService = service.ServiceGroupApp.BlogServiceGroup.NewsService


// CreateNews 创建News
// @Tags News
// @Summary 创建News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.News true "创建News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/createNews [post]
func (newsApi *NewsApi) CreateNews(c *gin.Context) {
	var news blog.News
	err := c.ShouldBindJSON(&news)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := newsService.CreateNews(news); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteNews 删除News
// @Tags News
// @Summary 删除News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.News true "删除News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteNews [delete]
func (newsApi *NewsApi) DeleteNews(c *gin.Context) {
	var news blog.News
	err := c.ShouldBindJSON(&news)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := newsService.DeleteNews(news); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteNewsByIds 批量删除News
// @Tags News
// @Summary 批量删除News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /news/deleteNewsByIds [delete]
func (newsApi *NewsApi) DeleteNewsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := newsService.DeleteNewsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateNews 更新News
// @Tags News
// @Summary 更新News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body blog.News true "更新News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /news/updateNews [put]
func (newsApi *NewsApi) UpdateNews(c *gin.Context) {
	var news blog.News
	err := c.ShouldBindJSON(&news)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := newsService.UpdateNews(news); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindNews 用id查询News
// @Tags News
// @Summary 用id查询News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query blog.News true "用id查询News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /news/findNews [get]
func (newsApi *NewsApi) FindNews(c *gin.Context) {
	var news blog.News
	err := c.ShouldBindQuery(&news)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if renews, err := newsService.GetNews(news.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"renews": renews}, c)
	}
}

// GetNewsList 分页获取News列表
// @Tags News
// @Summary 分页获取News列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query blogReq.NewsSearch true "分页获取News列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/getNewsList [get]
func (newsApi *NewsApi) GetNewsList(c *gin.Context) {
	var pageInfo blogReq.NewsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := newsService.GetNewsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

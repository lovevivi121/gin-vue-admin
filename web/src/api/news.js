import service from '@/utils/request'

// @Tags News
// @Summary 创建News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.News true "创建News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/createNews [post]
export const createNews = (data) => {
  return service({
    url: '/news/createNews',
    method: 'post',
    data
  })
}

// @Tags News
// @Summary 删除News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.News true "删除News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteNews [delete]
export const deleteNews = (data) => {
  return service({
    url: '/news/deleteNews',
    method: 'delete',
    data
  })
}

// @Tags News
// @Summary 删除News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteNews [delete]
export const deleteNewsByIds = (data) => {
  return service({
    url: '/news/deleteNewsByIds',
    method: 'delete',
    data
  })
}

// @Tags News
// @Summary 更新News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.News true "更新News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /news/updateNews [put]
export const updateNews = (data) => {
  return service({
    url: '/news/updateNews',
    method: 'put',
    data
  })
}

// @Tags News
// @Summary 用id查询News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.News true "用id查询News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /news/findNews [get]
export const findNews = (params) => {
  return service({
    url: '/news/findNews',
    method: 'get',
    params
  })
}

// @Tags News
// @Summary 分页获取News列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取News列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/getNewsList [get]
export const getNewsList = (params) => {
  return service({
    url: '/news/getNewsList',
    method: 'get',
    params
  })
}

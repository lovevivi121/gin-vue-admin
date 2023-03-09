package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/user"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type UsersApi struct {
}

var usersService = service.ServiceGroupApp.UserServiceGroup.UsersService


// CreateUsers 创建Users
// @Tags Users
// @Summary 创建Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Users true "创建Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /users/createUsers [post]
func (usersApi *UsersApi) CreateUsers(c *gin.Context) {
	var users user.Users
	err := c.ShouldBindJSON(&users)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    users.CreatedBy = utils.GetUserID(c)
	if err := usersService.CreateUsers(users); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUsers 删除Users
// @Tags Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Users true "删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /users/deleteUsers [delete]
func (usersApi *UsersApi) DeleteUsers(c *gin.Context) {
	var users user.Users
	err := c.ShouldBindJSON(&users)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    users.DeletedBy = utils.GetUserID(c)
	if err := usersService.DeleteUsers(users); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUsersByIds 批量删除Users
// @Tags Users
// @Summary 批量删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /users/deleteUsersByIds [delete]
func (usersApi *UsersApi) DeleteUsersByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := usersService.DeleteUsersByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUsers 更新Users
// @Tags Users
// @Summary 更新Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.Users true "更新Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /users/updateUsers [put]
func (usersApi *UsersApi) UpdateUsers(c *gin.Context) {
	var users user.Users
	err := c.ShouldBindJSON(&users)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    users.UpdatedBy = utils.GetUserID(c)
	if err := usersService.UpdateUsers(users); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUsers 用id查询Users
// @Tags Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /users/findUsers [get]
func (usersApi *UsersApi) FindUsers(c *gin.Context) {
	var users user.Users
	err := c.ShouldBindQuery(&users)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reusers, err := usersService.GetUsers(users.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reusers": reusers}, c)
	}
}

// GetUsersList 分页获取Users列表
// @Tags Users
// @Summary 分页获取Users列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.UsersSearch true "分页获取Users列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /users/getUsersList [get]
func (usersApi *UsersApi) GetUsersList(c *gin.Context) {
	var pageInfo userReq.UsersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := usersService.GetUsersInfoList(pageInfo); err != nil {
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

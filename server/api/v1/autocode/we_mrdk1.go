package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type WeMrdk1Api struct {
}

var weMrdk1Service = service.ServiceGroupApp.AutoCodeServiceGroup.WeMrdk1Service


// CreateWeMrdk1 创建WeMrdk1
// @Tags WeMrdk1
// @Summary 创建WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.WeMrdk1 true "创建WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk1/createWeMrdk1 [post]
func (weMrdk1Api *WeMrdk1Api) CreateWeMrdk1(c *gin.Context) {
	var weMrdk1 autocode.WeMrdk1
	_ = c.ShouldBindJSON(&weMrdk1)
	if err := weMrdk1Service.CreateWeMrdk1(weMrdk1); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWeMrdk1 删除WeMrdk1
// @Tags WeMrdk1
// @Summary 删除WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.WeMrdk1 true "删除WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weMrdk1/deleteWeMrdk1 [delete]
func (weMrdk1Api *WeMrdk1Api) DeleteWeMrdk1(c *gin.Context) {
	var weMrdk1 autocode.WeMrdk1
	_ = c.ShouldBindJSON(&weMrdk1)
	if err := weMrdk1Service.DeleteWeMrdk1(weMrdk1); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWeMrdk1ByIds 批量删除WeMrdk1
// @Tags WeMrdk1
// @Summary 批量删除WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /weMrdk1/deleteWeMrdk1ByIds [delete]
func (weMrdk1Api *WeMrdk1Api) DeleteWeMrdk1ByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := weMrdk1Service.DeleteWeMrdk1ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWeMrdk1 更新WeMrdk1
// @Tags WeMrdk1
// @Summary 更新WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.WeMrdk1 true "更新WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weMrdk1/updateWeMrdk1 [put]
func (weMrdk1Api *WeMrdk1Api) UpdateWeMrdk1(c *gin.Context) {
	var weMrdk1 autocode.WeMrdk1
	_ = c.ShouldBindJSON(&weMrdk1)
	if err := weMrdk1Service.UpdateWeMrdk1(weMrdk1); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWeMrdk1 用id查询WeMrdk1
// @Tags WeMrdk1
// @Summary 用id查询WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.WeMrdk1 true "用id查询WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weMrdk1/findWeMrdk1 [get]
func (weMrdk1Api *WeMrdk1Api) FindWeMrdk1(c *gin.Context) {
	var weMrdk1 autocode.WeMrdk1
	_ = c.ShouldBindQuery(&weMrdk1)
	fmt.Println("log_id: ",weMrdk1.LogId)
	fmt.Println("xuehao",weMrdk1.Xh)
	if err, reweMrdk1 := weMrdk1Service.GetWeMrdk1(weMrdk1.LogId); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reweMrdk1": reweMrdk1}, c)
	}
}

// GetWeMrdk1List 分页获取WeMrdk1列表
// @Tags WeMrdk1
// @Summary 分页获取WeMrdk1列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.WeMrdk1Search true "分页获取WeMrdk1列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk1/getWeMrdk1List [get]
func (weMrdk1Api *WeMrdk1Api) GetWeMrdk1List(c *gin.Context) {
	var pageInfo autocodeReq.WeMrdk1Search
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := weMrdk1Service.GetWeMrdk1InfoList(pageInfo); err != nil {
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

package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type WeMrdk2Api struct {
}

var weMrdk2Service = service.ServiceGroupApp.AutoCodeServiceGroup.WeMrdk2Service


// CreateWeMrdk2 创建WeMrdk2
// @Tags WeMrdk2
// @Summary 创建WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.WeMrdk2 true "创建WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk2/createWeMrdk2 [post]
func (weMrdk2Api *WeMrdk2Api) CreateWeMrdk2(c *gin.Context) {
	var weMrdk2 autocode.WeMrdk2
	_ = c.ShouldBindJSON(&weMrdk2)
	if err := weMrdk2Service.CreateWeMrdk2(weMrdk2); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWeMrdk2 删除WeMrdk2
// @Tags WeMrdk2
// @Summary 删除WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.WeMrdk2 true "删除WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weMrdk2/deleteWeMrdk2 [delete]
func (weMrdk2Api *WeMrdk2Api) DeleteWeMrdk2(c *gin.Context) {
	var weMrdk2 autocode.WeMrdk2
	_ = c.ShouldBindJSON(&weMrdk2)
	if err := weMrdk2Service.DeleteWeMrdk2(weMrdk2); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWeMrdk2ByIds 批量删除WeMrdk2
// @Tags WeMrdk2
// @Summary 批量删除WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /weMrdk2/deleteWeMrdk2ByIds [delete]
func (weMrdk2Api *WeMrdk2Api) DeleteWeMrdk2ByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := weMrdk2Service.DeleteWeMrdk2ByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWeMrdk2 更新WeMrdk2
// @Tags WeMrdk2
// @Summary 更新WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.WeMrdk2 true "更新WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weMrdk2/updateWeMrdk2 [put]
func (weMrdk2Api *WeMrdk2Api) UpdateWeMrdk2(c *gin.Context) {
	var weMrdk2 autocode.WeMrdk2
	_ = c.ShouldBindJSON(&weMrdk2)
	if err := weMrdk2Service.UpdateWeMrdk2(weMrdk2); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWeMrdk2 用id查询WeMrdk2
// @Tags WeMrdk2
// @Summary 用id查询WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.WeMrdk2 true "用id查询WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weMrdk2/findWeMrdk2 [get]
func (weMrdk2Api *WeMrdk2Api) FindWeMrdk2(c *gin.Context) {
	var weMrdk2 autocode.WeMrdk2
	_ = c.ShouldBindQuery(&weMrdk2)
	if err, reweMrdk2 := weMrdk2Service.GetWeMrdk2(weMrdk2.LogId); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reweMrdk2": reweMrdk2}, c)
	}
}

// GetWeMrdk2List 分页获取WeMrdk2列表
// @Tags WeMrdk2
// @Summary 分页获取WeMrdk2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.WeMrdk2Search true "分页获取WeMrdk2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk2/getWeMrdk2List [get]
func (weMrdk2Api *WeMrdk2Api) GetWeMrdk2List(c *gin.Context) {
	var pageInfo autocodeReq.WeMrdk2Search
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := weMrdk2Service.GetWeMrdk2InfoList(pageInfo); err != nil {
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

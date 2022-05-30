package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WeMrdk2Router struct {
}

// InitWeMrdk2Router 初始化 WeMrdk2 路由信息
func (s *WeMrdk2Router) InitWeMrdk2Router(Router *gin.RouterGroup) {
	weMrdk2Router := Router.Group("weMrdk2").Use(middleware.OperationRecord())
	weMrdk2RouterWithoutRecord := Router.Group("weMrdk2")
	var weMrdk2Api = v1.ApiGroupApp.AutoCodeApiGroup.WeMrdk2Api
	{
		weMrdk2Router.POST("createWeMrdk2", weMrdk2Api.CreateWeMrdk2)   // 新建WeMrdk2
		weMrdk2Router.DELETE("deleteWeMrdk2", weMrdk2Api.DeleteWeMrdk2) // 删除WeMrdk2
		weMrdk2Router.DELETE("deleteWeMrdk2ByIds", weMrdk2Api.DeleteWeMrdk2ByIds) // 批量删除WeMrdk2
		weMrdk2Router.PUT("updateWeMrdk2", weMrdk2Api.UpdateWeMrdk2)    // 更新WeMrdk2
	}
	{
		weMrdk2RouterWithoutRecord.GET("findWeMrdk2", weMrdk2Api.FindWeMrdk2)        // 根据ID获取WeMrdk2
		weMrdk2RouterWithoutRecord.GET("getWeMrdk2List", weMrdk2Api.GetWeMrdk2List)  // 获取WeMrdk2列表
	}
}

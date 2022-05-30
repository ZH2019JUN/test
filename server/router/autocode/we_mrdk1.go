package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WeMrdk1Router struct {
}

// InitWeMrdk1Router 初始化 WeMrdk1 路由信息
func (s *WeMrdk1Router) InitWeMrdk1Router(Router *gin.RouterGroup) {
	weMrdk1Router := Router.Group("weMrdk1").Use(middleware.OperationRecord())
	weMrdk1RouterWithoutRecord := Router.Group("weMrdk1")
	var weMrdk1Api = v1.ApiGroupApp.AutoCodeApiGroup.WeMrdk1Api
	{
		weMrdk1Router.POST("createWeMrdk1", weMrdk1Api.CreateWeMrdk1)   // 新建WeMrdk1
		weMrdk1Router.DELETE("deleteWeMrdk1", weMrdk1Api.DeleteWeMrdk1) // 删除WeMrdk1
		weMrdk1Router.DELETE("deleteWeMrdk1ByIds", weMrdk1Api.DeleteWeMrdk1ByIds) // 批量删除WeMrdk1
		weMrdk1Router.PUT("updateWeMrdk1", weMrdk1Api.UpdateWeMrdk1)    // 更新WeMrdk1
	}
	{
		weMrdk1RouterWithoutRecord.GET("findWeMrdk1", weMrdk1Api.FindWeMrdk1)        // 根据ID获取WeMrdk1
		weMrdk1RouterWithoutRecord.GET("getWeMrdk1List", weMrdk1Api.GetWeMrdk1List)  // 获取WeMrdk1列表
	}
}

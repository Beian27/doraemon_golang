package tool

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ToolSourceRouter struct {
}

// InitToolSourceRouter 初始化 ToolSource 路由信息
func (s *ToolSourceRouter) InitToolSourceRouter(Router *gin.RouterGroup) {
	toolSourceRouter := Router.Group("toolSource").Use(middleware.OperationRecord())
	toolSourceRouterWithoutRecord := Router.Group("toolSource")
	var toolSourceApi = v1.ApiGroupApp.ToolApiGroup.ToolSourceApi
	{
		toolSourceRouter.POST("createToolSource", toolSourceApi.CreateToolSource)             // 新建ToolSource
		toolSourceRouter.DELETE("deleteToolSource", toolSourceApi.DeleteToolSource)           // 删除ToolSource
		toolSourceRouter.DELETE("deleteToolSourceByIds", toolSourceApi.DeleteToolSourceByIds) // 批量删除ToolSource
		toolSourceRouter.PUT("updateToolSource", toolSourceApi.UpdateToolSource)              // 更新ToolSource
	}
	{
		toolSourceRouterWithoutRecord.GET("findToolSource", toolSourceApi.FindToolSource)       // 根据ID获取ToolSource
		toolSourceRouterWithoutRecord.GET("getToolSourceList", toolSourceApi.GetToolSourceList) // 获取ToolSource列表
	}
}

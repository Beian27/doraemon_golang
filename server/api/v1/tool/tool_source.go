package tool

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tool"
	toolReq "github.com/flipped-aurora/gin-vue-admin/server/model/tool/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ToolSourceApi struct {
}

var toolSourceService = service.ServiceGroupApp.ToolServiceGroup.ToolSourceService

// CreateToolSource 创建ToolSource
// @Tags ToolSource
// @Summary 创建ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ToolSource true "创建ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolSource/createToolSource [post]
func (toolSourceApi *ToolSourceApi) CreateToolSource(c *gin.Context) {
	var toolSource tool.ToolSource
	_ = c.ShouldBindJSON(&toolSource)
	if err := toolSourceService.CreateToolSource(toolSource); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteToolSource 删除ToolSource
// @Tags ToolSource
// @Summary 删除ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ToolSource true "删除ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /toolSource/deleteToolSource [delete]
func (toolSourceApi *ToolSourceApi) DeleteToolSource(c *gin.Context) {
	var toolSource tool.ToolSource
	_ = c.ShouldBindJSON(&toolSource)
	if err := toolSourceService.DeleteToolSource(toolSource); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteToolSourceByIds 批量删除ToolSource
// @Tags ToolSource
// @Summary 批量删除ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /toolSource/deleteToolSourceByIds [delete]
func (toolSourceApi *ToolSourceApi) DeleteToolSourceByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := toolSourceService.DeleteToolSourceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateToolSource 更新ToolSource
// @Tags ToolSource
// @Summary 更新ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ToolSource true "更新ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /toolSource/updateToolSource [put]
func (toolSourceApi *ToolSourceApi) UpdateToolSource(c *gin.Context) {
	var toolSource tool.ToolSource
	_ = c.ShouldBindJSON(&toolSource)
	if err := toolSourceService.UpdateToolSource(toolSource); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindToolSource 用id查询ToolSource
// @Tags ToolSource
// @Summary 用id查询ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ToolSource true "用id查询ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /toolSource/findToolSource [get]
func (toolSourceApi *ToolSourceApi) FindToolSource(c *gin.Context) {
	var toolSource tool.ToolSource
	_ = c.ShouldBindQuery(&toolSource)
	if err, retoolSource := toolSourceService.GetToolSource(toolSource.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retoolSource": retoolSource}, c)
	}
}

// GetToolSourceList 分页获取ToolSource列表
// @Tags ToolSource
// @Summary 分页获取ToolSource列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ToolSourceSearch true "分页获取ToolSource列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolSource/getToolSourceList [get]
func (toolSourceApi *ToolSourceApi) GetToolSourceList(c *gin.Context) {
	var pageInfo toolReq.ToolSourceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := toolSourceService.GetToolSourceInfoList(pageInfo); err != nil {
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

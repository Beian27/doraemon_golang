import service from '@/utils/request'

// @Tags ToolSource
// @Summary 创建ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ToolSource true "创建ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolSource/createToolSource [post]
export const createToolSource = (data) => {
  return service({
    url: '/toolSource/createToolSource',
    method: 'post',
    data
  })
}

// @Tags ToolSource
// @Summary 删除ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ToolSource true "删除ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /toolSource/deleteToolSource [delete]
export const deleteToolSource = (data) => {
  return service({
    url: '/toolSource/deleteToolSource',
    method: 'delete',
    data
  })
}

// @Tags ToolSource
// @Summary 删除ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /toolSource/deleteToolSource [delete]
export const deleteToolSourceByIds = (data) => {
  return service({
    url: '/toolSource/deleteToolSourceByIds',
    method: 'delete',
    data
  })
}

// @Tags ToolSource
// @Summary 更新ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ToolSource true "更新ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /toolSource/updateToolSource [put]
export const updateToolSource = (data) => {
  return service({
    url: '/toolSource/updateToolSource',
    method: 'put',
    data
  })
}

// @Tags ToolSource
// @Summary 用id查询ToolSource
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ToolSource true "用id查询ToolSource"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /toolSource/findToolSource [get]
export const findToolSource = (params) => {
  return service({
    url: '/toolSource/findToolSource',
    method: 'get',
    params
  })
}

// @Tags ToolSource
// @Summary 分页获取ToolSource列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ToolSource列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /toolSource/getToolSourceList [get]
export const getToolSourceList = (params) => {
  return service({
    url: '/toolSource/getToolSourceList',
    method: 'get',
    params
  })
}

import service from '@/utils/request'

// @Tags WeMrdk2
// @Summary 创建WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeMrdk2 true "创建WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk2/createWeMrdk2 [post]
export const createWeMrdk2 = (data) => {
  return service({
    url: '/weMrdk2/createWeMrdk2',
    method: 'post',
    data
  })
}

// @Tags WeMrdk2
// @Summary 删除WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeMrdk2 true "删除WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weMrdk2/deleteWeMrdk2 [delete]
export const deleteWeMrdk2 = (data) => {
  return service({
    url: '/weMrdk2/deleteWeMrdk2',
    method: 'delete',
    data
  })
}

// @Tags WeMrdk2
// @Summary 删除WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weMrdk2/deleteWeMrdk2 [delete]
export const deleteWeMrdk2ByIds = (data) => {
  return service({
    url: '/weMrdk2/deleteWeMrdk2ByIds',
    method: 'delete',
    data
  })
}

// @Tags WeMrdk2
// @Summary 更新WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeMrdk2 true "更新WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weMrdk2/updateWeMrdk2 [put]
export const updateWeMrdk2 = (data) => {
  return service({
    url: '/weMrdk2/updateWeMrdk2',
    method: 'put',
    data
  })
}

// @Tags WeMrdk2
// @Summary 用id查询WeMrdk2
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WeMrdk2 true "用id查询WeMrdk2"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weMrdk2/findWeMrdk2 [get]
export const findWeMrdk2 = (params) => {
  return service({
    url: '/weMrdk2/findWeMrdk2',
    method: 'get',
    params
  })
}

// @Tags WeMrdk2
// @Summary 分页获取WeMrdk2列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WeMrdk2列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk2/getWeMrdk2List [get]
export const getWeMrdk2List = (params) => {
  return service({
    url: '/weMrdk2/getWeMrdk2List',
    method: 'get',
    params
  })
}

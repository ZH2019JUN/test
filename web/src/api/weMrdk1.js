import service from '@/utils/request'

// @Tags WeMrdk1
// @Summary 创建WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeMrdk1 true "创建WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk1/createWeMrdk1 [post]
export const createWeMrdk1 = (data) => {
  return service({
    url: '/weMrdk1/createWeMrdk1',
    method: 'post',
    data
  })
}

// @Tags WeMrdk1
// @Summary 删除WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeMrdk1 true "删除WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weMrdk1/deleteWeMrdk1 [delete]
export const deleteWeMrdk1 = (data) => {
  return service({
    url: '/weMrdk1/deleteWeMrdk1',
    method: 'delete',
    data
  })
}

// @Tags WeMrdk1
// @Summary 删除WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /weMrdk1/deleteWeMrdk1 [delete]
export const deleteWeMrdk1ByIds = (data) => {
  return service({
    url: '/weMrdk1/deleteWeMrdk1ByIds',
    method: 'delete',
    data
  })
}

// @Tags WeMrdk1
// @Summary 更新WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WeMrdk1 true "更新WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /weMrdk1/updateWeMrdk1 [put]
export const updateWeMrdk1 = (data) => {
  return service({
    url: '/weMrdk1/updateWeMrdk1',
    method: 'put',
    data
  })
}

// @Tags WeMrdk1
// @Summary 用id查询WeMrdk1
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WeMrdk1 true "用id查询WeMrdk1"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /weMrdk1/findWeMrdk1 [get]
export const findWeMrdk1 = (params) => {
  return service({
    url: '/weMrdk1/findWeMrdk1',
    method: 'get',
    params
  })
}

// @Tags WeMrdk1
// @Summary 分页获取WeMrdk1列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WeMrdk1列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /weMrdk1/getWeMrdk1List [get]
export const getWeMrdk1List = (params) => {
  return service({
    url: '/weMrdk1/getWeMrdk1List',
    method: 'get',
    params
  })
}

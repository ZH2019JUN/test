package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type WeMrdk2Service struct {
}

// CreateWeMrdk2 创建WeMrdk2记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk2Service *WeMrdk2Service) CreateWeMrdk2(weMrdk2 autocode.WeMrdk2) (err error) {
	err = global.GVA_DB.Create(&weMrdk2).Error
	return err
}

// DeleteWeMrdk2 删除WeMrdk2记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk2Service *WeMrdk2Service)DeleteWeMrdk2(weMrdk2 autocode.WeMrdk2) (err error) {
	err = global.GVA_DB.Delete(&weMrdk2).Error
	return err
}

// DeleteWeMrdk2ByIds 批量删除WeMrdk2记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk2Service *WeMrdk2Service)DeleteWeMrdk2ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.WeMrdk2{},"log_id in ?",ids.Ids).Error
	return err
}

// UpdateWeMrdk2 更新WeMrdk2记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk2Service *WeMrdk2Service)UpdateWeMrdk2(weMrdk2 autocode.WeMrdk2) (err error) {
	err = global.GVA_DB.Save(&weMrdk2).Error
	return err
}

// GetWeMrdk2 根据id获取WeMrdk2记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk2Service *WeMrdk2Service)GetWeMrdk2(id uint) (err error, weMrdk2 autocode.WeMrdk2) {
	err = global.GVA_DB.Where("log_id = ?", id).First(&weMrdk2).Error
	return
}

// GetWeMrdk2InfoList 分页获取WeMrdk2记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk2Service *WeMrdk2Service)GetWeMrdk2InfoList(info autoCodeReq.WeMrdk2Search) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.WeMrdk2{})
    var weMrdk2s []autocode.WeMrdk2
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Xh != "" {
        db = db.Where("xh = ?",info.Xh)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&weMrdk2s).Error
	return err, weMrdk2s, total
}

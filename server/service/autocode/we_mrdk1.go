package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type WeMrdk1Service struct {
}

// CreateWeMrdk1 创建WeMrdk1记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk1Service *WeMrdk1Service) CreateWeMrdk1(weMrdk1 autocode.WeMrdk1) (err error) {
	err = global.GVA_DB.Create(&weMrdk1).Error
	return err
}

// DeleteWeMrdk1 删除WeMrdk1记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk1Service *WeMrdk1Service)DeleteWeMrdk1(weMrdk1 autocode.WeMrdk1) (err error) {
	err = global.GVA_DB.Delete(&weMrdk1).Error
	return err
}

// DeleteWeMrdk1ByIds 批量删除WeMrdk1记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk1Service *WeMrdk1Service)DeleteWeMrdk1ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.WeMrdk1{},"log_id in ?",ids.Ids).Error
	return err
}

// UpdateWeMrdk1 更新WeMrdk1记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk1Service *WeMrdk1Service)UpdateWeMrdk1(weMrdk1 autocode.WeMrdk1) (err error) {
	err = global.GVA_DB.Save(&weMrdk1).Error
	return err
}

// GetWeMrdk1 根据id获取WeMrdk1记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk1Service *WeMrdk1Service)GetWeMrdk1(id uint) (err error, weMrdk1 autocode.WeMrdk1) {
	err = global.GVA_DB.Where("log_id = ?", id).First(&weMrdk1).Error
	return
}

// GetWeMrdk1InfoList 分页获取WeMrdk1记录
// Author [piexlmax](https://github.com/piexlmax)
func (weMrdk1Service *WeMrdk1Service)GetWeMrdk1InfoList(info autoCodeReq.WeMrdk1Search) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.WeMrdk1{})
    var weMrdk1s []autocode.WeMrdk1
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Xh != "" {
        db = db.Where("xh = ?",info.Xh)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&weMrdk1s).Error
	return err, weMrdk1s, total
}

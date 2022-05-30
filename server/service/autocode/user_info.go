package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserInfoService struct {
}

// CreateUserInfo 创建UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService) CreateUserInfo(userInfo autocode.UserInfo) (err error) {
	err = global.GVA_DB.Create(&userInfo).Error
	return err
}

// DeleteUserInfo 删除UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService)DeleteUserInfo(userInfo autocode.UserInfo) (err error) {
	err = global.GVA_DB.Delete(&userInfo).Error
	return err
}

// DeleteUserInfoByIds 批量删除UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService)DeleteUserInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.UserInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateUserInfo 更新UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService)UpdateUserInfo(userInfo autocode.UserInfo) (err error) {
	err = global.GVA_DB.Save(&userInfo).Error
	return err
}

// GetUserInfo 根据id获取UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService)GetUserInfo(xh string) (err error, userInfo autocode.UserInfo) {
	err = global.GVA_DB.Where("xh = ?", xh).First(&userInfo).Error
	return
}

// GetUserInfoInfoList 分页获取UserInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (userInfoService *UserInfoService)GetUserInfoInfoList(info autoCodeReq.UserInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.UserInfo{})
    var userInfos []autocode.UserInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    //增加学院条件
    db = db.Where("xy = ?",info.Xy)
    if info.Xh != "" {
        db = db.Where("xh = ?",info.Xh)
    }
    if info.Xm != "" {
        db = db.Where("xm LIKE ?","%"+ info.Xm+"%")
    }
    if info.Bj != "" {
        db = db.Where("bj = ?",info.Bj)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&userInfos).Error
	return err, userInfos, total
}

package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	bmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/bmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BmtUserService struct {
}

// CreateBmtUser 创建BmtUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUserService *BmtUserService) CreateBmtUser(bmtUser bmt.BmtUser) (err error) {
	err = global.GVA_DB.Create(&bmtUser).Error
	return err
}

// DeleteBmtUser 删除BmtUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUserService *BmtUserService) DeleteBmtUser(bmtUser bmt.BmtUser) (err error) {
	err = global.GVA_DB.Delete(&bmtUser).Error
	return err
}

// DeleteBmtUserByIds 批量删除BmtUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUserService *BmtUserService) DeleteBmtUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]bmt.BmtUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBmtUser 更新BmtUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUserService *BmtUserService) UpdateBmtUser(bmtUser bmt.BmtUser) (err error) {
	err = global.GVA_DB.Save(&bmtUser).Error
	return err
}

// GetBmtUser 根据id获取BmtUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUserService *BmtUserService) GetBmtUser(id uint) (err error, bmtUser bmt.BmtUser) {
	err = global.GVA_DB.Where("id = ?", id).First(&bmtUser).Error
	return
}

// GetBmtUserInfoList 分页获取BmtUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUserService *BmtUserService) GetBmtUserInfoList(info bmtReq.BmtUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bmt.BmtUser{})
	var bmtUsers []bmt.BmtUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.BmtId != nil {
		db = db.Where("bmt_id = ?", info.BmtId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&bmtUsers).Error
	return err, bmtUsers, total
}

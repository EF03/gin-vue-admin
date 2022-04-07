package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	bmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/bmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BmtActivityService struct {
}

// CreateBmtActivity 创建BmtActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtActivityService *BmtActivityService) CreateBmtActivity(bmtActivity bmt.BmtActivity) (err error) {
	err = global.GVA_DB.Create(&bmtActivity).Error
	return err
}

// DeleteBmtActivity 删除BmtActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtActivityService *BmtActivityService) DeleteBmtActivity(bmtActivity bmt.BmtActivity) (err error) {
	err = global.GVA_DB.Delete(&bmtActivity).Error
	return err
}

// DeleteBmtActivityByIds 批量删除BmtActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtActivityService *BmtActivityService) DeleteBmtActivityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]bmt.BmtActivity{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBmtActivity 更新BmtActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtActivityService *BmtActivityService) UpdateBmtActivity(bmtActivity bmt.BmtActivity) (err error) {
	err = global.GVA_DB.Save(&bmtActivity).Error
	return err
}

// GetBmtActivity 根据id获取BmtActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtActivityService *BmtActivityService) GetBmtActivity(id uint) (err error, bmtActivity bmt.BmtActivity) {
	err = global.GVA_DB.Where("id = ?", id).First(&bmtActivity).Error
	return
}

// GetBmtActivityInfoList 分页获取BmtActivity记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtActivityService *BmtActivityService) GetBmtActivityInfoList(info bmtReq.BmtActivitySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bmt.BmtActivity{})
	var bmtActivitys []bmt.BmtActivity
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&bmtActivitys).Error
	return err, bmtActivitys, total
}

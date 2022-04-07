package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	bmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/bmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BmtPeriodService struct {
}

// CreateBmtPeriod 创建BmtPeriod记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtPeriodService *BmtPeriodService) CreateBmtPeriod(bmtPeriod bmt.BmtPeriod) (err error) {
	err = global.GVA_DB.Create(&bmtPeriod).Error
	return err
}

// DeleteBmtPeriod 删除BmtPeriod记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtPeriodService *BmtPeriodService) DeleteBmtPeriod(bmtPeriod bmt.BmtPeriod) (err error) {
	err = global.GVA_DB.Delete(&bmtPeriod).Error
	return err
}

// DeleteBmtPeriodByIds 批量删除BmtPeriod记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtPeriodService *BmtPeriodService) DeleteBmtPeriodByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]bmt.BmtPeriod{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBmtPeriod 更新BmtPeriod记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtPeriodService *BmtPeriodService) UpdateBmtPeriod(bmtPeriod bmt.BmtPeriod) (err error) {
	err = global.GVA_DB.Save(&bmtPeriod).Error
	return err
}

// GetBmtPeriod 根据id获取BmtPeriod记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtPeriodService *BmtPeriodService) GetBmtPeriod(id uint) (err error, bmtPeriod bmt.BmtPeriod) {
	err = global.GVA_DB.Where("id = ?", id).First(&bmtPeriod).Error
	return
}

// GetBmtPeriodInfoList 分页获取BmtPeriod记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtPeriodService *BmtPeriodService) GetBmtPeriodInfoList(info bmtReq.BmtPeriodSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bmt.BmtPeriod{})
	var bmtPeriods []bmt.BmtPeriod
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Winner != nil {
		db = db.Where("winner = ?", info.Winner)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&bmtPeriods).Error
	return err, bmtPeriods, total
}

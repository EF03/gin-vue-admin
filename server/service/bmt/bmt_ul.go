package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	bmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/bmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BmtUlService struct {
}

// CreateBmtUl 创建BmtUl记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUlService *BmtUlService) CreateBmtUl(bmtUl bmt.BmtUl) (err error) {
	err = global.GVA_DB.Create(&bmtUl).Error
	return err
}

// DeleteBmtUl 删除BmtUl记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUlService *BmtUlService) DeleteBmtUl(bmtUl bmt.BmtUl) (err error) {
	err = global.GVA_DB.Delete(&bmtUl).Error
	return err
}

// DeleteBmtUlByIds 批量删除BmtUl记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUlService *BmtUlService) DeleteBmtUlByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]bmt.BmtUl{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateBmtUl 更新BmtUl记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUlService *BmtUlService) UpdateBmtUl(bmtUl bmt.BmtUl) (err error) {
	err = global.GVA_DB.Save(&bmtUl).Error
	return err
}

// GetBmtUl 根据id获取BmtUl记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUlService *BmtUlService) GetBmtUl(id uint) (err error, bmtUl bmt.BmtUl) {
	err = global.GVA_DB.Where("id = ?", id).First(&bmtUl).Error
	return
}

// GetBmtUlInfoList 分页获取BmtUl记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtUlService *BmtUlService) GetBmtUlInfoList(info bmtReq.BmtUlSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&bmt.BmtUl{})
	var bmtUls []bmt.BmtUl
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+info.Remark+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&bmtUls).Error
	return err, bmtUls, total
}

package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	bmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/bmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"math/rand"

	"time"
)

type BmtPeriodService struct {
}

// CreateBmtPeriod 创建BmtPeriod记录
// Author [piexlmax](https://github.com/piexlmax)
func (bmtPeriodService *BmtPeriodService) CreateBmtPeriod(bmtPeriod bmt.BmtPeriod) (err error) {
	err = global.GVA_DB.Create(&bmtPeriod).Error
	return err
}

// DrawBmtPeriod 抽籤BmtPeriod记录
func (bmtPeriodService *BmtPeriodService) DrawBmtPeriod(bmtPeriod bmt.BmtPeriod) (err error) {
	//// 创建db
	//db := global.GVA_DB.Model(&bmt.BmtUser{})
	//var bmtUsers []bmt.BmtUser
	//// 如果有条件搜索 下方会自动创建搜索语句
	//if bmtPeriod.BmtId != nil {
	//	db = db.Where("bmt_id = ?", bmtPeriod.BmtId)
	//}
	//err = db.Find(&bmtUsers).Error
	var total int64
	//total := 0
	limit := 1000
	offset := 0
	// 创建db
	db := global.GVA_DB.Model(&bmt.BmtUser{})
	var bmtUsers []bmt.BmtUser
	var bmtUsersResult = make([]bmt.BmtUser, 0)

	//var bmtUsersResult = []bmt.BmtUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if bmtPeriod.BmtId != nil {
		db = db.Where("bmt_id = ?", bmtPeriod.BmtId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&bmtUsers).Error

	bmtUserMap := bmt.SplitSlice(bmtUsers)

	for i := 0; i < len(bmtUserMap); i++ {
		temp := bmtUserMap[i]
		rand.Seed(time.Now().UnixNano())
		for k := len(temp) - 1; k > 0; k-- { // Fisher–Yates shuffle
			j := rand.Intn(k + 1)
			temp[k], temp[j] = temp[j], temp[k]
		}
		for j := 0; j < len(temp); j++ {
			if len(bmtUsersResult) == 4 {
				break
			}
			tempInt := *(temp[j].Times) + 1
			temp[j].Times = &tempInt
			//fmt.Printf("Uid = %d,times = %d\n", *(temp[j].Uid), *(temp[j].Times))
			bmtUsersResult = append(bmtUsersResult, temp[j])
		}
	}
	err = global.GVA_DB.Save(&bmtUsersResult).Error

	result := bmt.BmtPeriod{
		BmtId:       bmtPeriod.BmtId,
		RedUserId1:  bmtUsersResult[0].Uid,
		RedUserId2:  bmtUsersResult[1].Uid,
		BlueUserId1: bmtUsersResult[2].Uid,
		BlueUserId2: bmtUsersResult[3].Uid,
	}
	err = global.GVA_DB.Create(&result).Error
	//bmtUsersResult
	//for _, tt := range bmtUsersResult {
	//	fmt.Printf("uid = %d,times = %d\n", *(tt.Uid), *(tt.Times))
	//}
	//fmt.Println("bmtUsersResult =", bmtUsersResult)
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
	if info.BmtId != nil {
		db = db.Where("bmt_id = ?", info.BmtId)
	}
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

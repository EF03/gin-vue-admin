package bmt

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	bmtReq "github.com/flipped-aurora/gin-vue-admin/server/model/bmt/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/mitchellh/hashstructure/v2"
	"math/rand"
	"sort"
	"strconv"
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

//// DrawBmtPeriod 抽籤BmtPeriod记录
//func (bmtPeriodService *BmtPeriodService) DrawBmtPeriod(bmtPeriod bmt.BmtPeriod) (err error) {
//	var total int64
//	limit := 1000
//	offset := 0
//	// 创建db
//	db := global.GVA_DB.Model(&bmt.BmtUser{})
//	var bmtUsers []bmt.BmtUser
//	//var bmtUsersResult = []bmt.BmtUser
//	// 如果有条件搜索 下方会自动创建搜索语句
//	if bmtPeriod.BmtId != nil {
//		db = db.Where("bmt_id = ?", bmtPeriod.BmtId)
//	}
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//	err = db.Limit(limit).Offset(offset).Find(&bmtUsers).Error
//
//	bmtUserMultiSlice := bmt.SplitSlice(bmtUsers)
//
//	var bmtUsersResult = make([]bmt.BmtUser, 0)
//	//var bmtUsersResult []bmt.BmtUser
//	// 做出三次並算出雜湊值
//	//for k := 0; k < 3; k++ {
//	for i := 0; i < len(bmtUserMultiSlice); i++ {
//		temp := bmtUserMultiSlice[i]
//		// 洗牌
//		shuffle(temp)
//		for j := 0; j < len(temp); j++ {
//			if len(bmtUsersResult) == 4 {
//				break
//			}
//			tempInt := *(temp[j].Times) + 1
//			temp[j].Times = &tempInt
//			//fmt.Printf("Uid = %d,times = %d\n", *(temp[j].Uid), *(temp[j].Times))
//			bmtUsersResult = append(bmtUsersResult, temp[j])
//		}
//
//	}
//	// 排序後算出雜湊值
//	sort.Sort(bmt.ByUid(bmtUsersResult))
//	v1 := []bmt.BmtUser{
//		bmtUsersResult[0],
//		bmtUsersResult[1],
//	}
//	v2 := []bmt.BmtUser{
//		bmtUsersResult[2],
//		bmtUsersResult[3],
//	}
//	hash1, _ := hashstructure.Hash(v1, hashstructure.FormatV2, nil)
//	hash2, _ := hashstructure.Hash(v2, hashstructure.FormatV2, nil)
//	//}
//	// 最後結果洗牌
//	shuffle(bmtUsersResult)
//	err = global.GVA_DB.Save(&bmtUsersResult).Error
//	hash1Str := strconv.FormatUint(hash1, 10)
//	hash2Str := strconv.FormatUint(hash2, 10)
//
//	result := bmt.BmtPeriod{
//		BmtId:       bmtPeriod.BmtId,
//		RedUserId1:  bmtUsersResult[0].Uid,
//		RedUserId2:  bmtUsersResult[1].Uid,
//		BlueUserId1: bmtUsersResult[2].Uid,
//		BlueUserId2: bmtUsersResult[3].Uid,
//		Hash1:       &hash1Str,
//		Hash2:       &hash2Str,
//	}
//	err = global.GVA_DB.Create(&result).Error
//	return err
//}

// DrawBmtPeriod 抽籤BmtPeriod记录
func (bmtPeriodService *BmtPeriodService) DrawBmtPeriod(bmtPeriod bmt.BmtPeriod) (err error) {
	db := global.GVA_DB.Model(&bmt.BmtUser{})
	var bmtUsers []bmt.BmtUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if bmtPeriod.BmtId != nil {
		db = db.Where("bmt_id = ?", bmtPeriod.BmtId)
	}
	err = db.Find(&bmtUsers).Error
	if err != nil {
		return
	}
	bmtUserMultiSlice := bmt.SplitSlice(bmtUsers)

	// 创建db
	db2 := global.GVA_DB.Model(&bmt.BmtPeriod{})
	var bmtPeriods []bmt.BmtPeriod
	// 如果有条件搜索 下方会自动创建搜索语句
	if bmtPeriod.BmtId != nil {
		db2 = db2.Where("bmt_id = ?", bmtPeriod.BmtId)
	}
	err = db2.Find(&bmtPeriods).Error

	var bmtUsersResult []bmt.BmtUser
	var max = int(^uint(0) >> 1)

	var bmtUsersSlice = make([]bmt.BmtPeriod, 0)
	var result bmt.BmtPeriod
	// 做出三次並算出雜湊值
	for k := 0; k < 5; k++ {
		var bmtUsersTemp = make([]bmt.BmtUser, 0)
		for i := 0; i < len(bmtUserMultiSlice); i++ {
			temp := bmtUserMultiSlice[i]
			// 洗牌
			shuffle(temp)
			for j := 0; j < len(temp); j++ {
				if len(bmtUsersTemp) == 4 {
					break
				}
				bmtUsersTemp = append(bmtUsersTemp, temp[j])
			}
		}
		// 排序後算出雜湊值
		sort.Sort(bmt.ByUid(bmtUsersTemp))
		v1 := []bmt.BmtUser{
			bmtUsersTemp[0],
			bmtUsersTemp[1],
		}
		v2 := []bmt.BmtUser{
			bmtUsersTemp[2],
			bmtUsersTemp[3],
		}

		hash1, _ := hashstructure.Hash(v1, hashstructure.FormatV2, nil)
		hash2, _ := hashstructure.Hash(v2, hashstructure.FormatV2, nil)
		hash1Str := strconv.FormatUint(hash1, 10)
		hash2Str := strconv.FormatUint(hash2, 10)

		// 最後結果洗牌
		shuffle(bmtUsersTemp)

		bmtPeriod := bmt.BmtPeriod{
			BmtId:       bmtPeriod.BmtId,
			RedUserId1:  bmtUsersTemp[0].Uid,
			RedUserId2:  bmtUsersTemp[1].Uid,
			BlueUserId1: bmtUsersTemp[2].Uid,
			BlueUserId2: bmtUsersTemp[3].Uid,
			Hash1:       &hash1Str,
			Hash2:       &hash2Str,
		}
		bmtUsersSlice = append(bmtUsersSlice, bmtPeriod)

		var score = 0
		for _, period := range bmtPeriods {
			if *period.Hash1 == hash1Str {
				score++
			}
			if *period.Hash2 == hash2Str {
				score++
			}
		}
		if max > score {
			fmt.Printf("max = %d,score = %d\n", max, score)
			max = score
			bmtUsersResult = bmtUsersTemp
			result = bmtPeriod
		}
	}

	//// 最後結果洗牌
	for j := range bmtUsersResult {
		tempInt := *(bmtUsersResult[j].Times) + 1
		bmtUsersResult[j].Times = &tempInt
	}

	err = global.GVA_DB.Save(&bmtUsersResult).Error
	if err != nil {
		return
	}
	err = global.GVA_DB.Create(&result).Error
	return err
}

func shuffle(temp []bmt.BmtUser) {
	rand.Seed(time.Now().UnixNano())
	for k := len(temp) - 1; k > 0; k-- { // Fisher–Yates shuffle
		j := rand.Intn(k + 1)
		temp[k], temp[j] = temp[j], temp[k]
	}
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

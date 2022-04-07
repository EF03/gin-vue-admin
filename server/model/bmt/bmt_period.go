// 自动生成模板BmtPeriod
package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BmtPeriod 结构体
// 如果含有time.Time 请自行import time包
type BmtPeriod struct {
	global.GVA_MODEL
	BmtId       *int `json:"bmtId" form:"bmtId" gorm:"column:bmt_id;comment:羽球活动主键;size:19;"`
	RedUserId1  *int `json:"redUserId1" form:"redUserId1" gorm:"column:red_user_id1;comment:红对队员1;size:10;"`
	RedUserId2  *int `json:"redUserId2" form:"redUserId2" gorm:"column:red_user_id2;comment:红对队员2;size:10;"`
	BlueUserId1 *int `json:"blueUserId1" form:"blueUserId1" gorm:"column:blue_user_id1;comment:蓝对队员1;size:10;"`
	BlueUserId2 *int `json:"blueUserId2" form:"blueUserId2" gorm:"column:blue_user_id2;comment:蓝对队员2;size:10;"`
	RedScore    *int `json:"redScore" form:"redScore" gorm:"column:red_score;comment:红对分数;size:10;"`
	BlueScore   *int `json:"blueScore" form:"blueScore" gorm:"column:blue_score;comment:藍隊分數;size:10;"`
	Winner      *int `json:"winner" form:"winner" gorm:"column:winner;comment:胜负：1-红对,2-蓝队;size:2;"`
}

// TableName BmtPeriod 表名
func (BmtPeriod) TableName() string {
	return "bmt_period"
}

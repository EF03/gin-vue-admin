// 自动生成模板BmtUser
package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BmtUser 结构体
// 如果含有time.Time 请自行import time包
type BmtUser struct {
	global.GVA_MODEL
	BmtId *int `json:"bmtId" form:"bmtId" gorm:"column:bmt_id;comment:羽球活动id;size:19;"`
	Uid   *int `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:10;"`
	Times *int `json:"times" form:"times" gorm:"column:times;comment:参予小节次数;size:10;"`
}

// TableName BmtUser 表名
func (BmtUser) TableName() string {
	return "bmt_user"
}

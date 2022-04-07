// 自动生成模板BmtActivity
package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// BmtActivity 结构体
// 如果含有time.Time 请自行import time包
type BmtActivity struct {
	global.GVA_MODEL
	Name      string     `json:"name" form:"name" gorm:"column:name;comment:活动主题;size:50;"`
	StartTime *time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:活动开始时间;"`
	EndTime   *time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:活动结束时间;"`
	Content   string     `json:"content" form:"content" gorm:"column:content;comment:活动内容;size:250;"`
	Cost      *int       `json:"cost" form:"cost" gorm:"column:cost;comment:活动花费;size:10;"`
	PersonNum *int       `json:"personNum" form:"personNum" gorm:"column:person_num;comment:活动人数;"`
}

// TableName BmtActivity 表名
func (BmtActivity) TableName() string {
	return "bmt_activity"
}

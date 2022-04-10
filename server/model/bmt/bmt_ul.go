// 自动生成模板BmtUl
package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BmtUl 结构体
// 如果含有time.Time 请自行import time包
type BmtUl struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"column:name;comment:名称;size:50;"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:250;"`
}

// TableName BmtUl 表名
func (BmtUl) TableName() string {
	return "bmt_ul"
}

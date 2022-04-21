// 自动生成模板BmtUser
package bmt

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"sort"
)

// BmtUser 结构体
// 如果含有time.Time 请自行import time包
type BmtUser struct {
	global.GVA_MODEL `hash:"ignore"`
	BmtId            *int `json:"bmtId" form:"bmtId" gorm:"column:bmt_id;comment:羽球活动id;size:19;" hash:"ignore"`
	Uid              *int `json:"uid" form:"uid" gorm:"column:uid;comment:參加者;size:10;"`
	Times            *int `json:"times" form:"times" gorm:"column:times;comment:参予小节次数;size:10;" hash:"ignore"`
}

// TableName BmtUser 表名
func (BmtUser) TableName() string {
	return "bmt_user"
}

type ByTimes []BmtUser

func (a ByTimes) Len() int           { return len(a) }
func (a ByTimes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTimes) Less(i, j int) bool { return *(a[i].Times) < *(a[j].Times) }

type ByUid []BmtUser

func (a ByUid) Len() int           { return len(a) }
func (a ByUid) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUid) Less(i, j int) bool { return *(a[i].Uid) < *(a[j].Uid) }

//KeyFunc func(item interface{}) interface{}
//
//func (s Stream) Group(fn KeyFunc) Stream {
//	groups := make(map[interface{}][]interface{})
//	for item := range s.source {
//		key := fn(item)
//		groups[key] = append(groups[key], item)
//	}
//	source := make(chan interface{})
//	go func() {
//		for _, group := range groups {
//			source <- group
//		}
//		close(source)
//	}()
//	return Range(source)
//}

// 切片分組
func SplitSlice(list []BmtUser) [][]BmtUser {

	//ss := []string{"golang", "google", "php", "python", "java", "c++"}
	//fx.From(func(source chan<- interface{}) {
	//	for _, s := range ss {
	//		source <- s
	//	}
	//}).Group(func(item interface{}) interface{} {
	//	if strings.HasPrefix(item.(string), "g") {
	//		return "g"
	//	} else if strings.HasPrefix(item.(string), "p") {
	//		return "p"
	//	}
	//	return ""
	//}).ForEach(func(item interface{}) {
	//	fmt.Println(item)
	//})

	//========
	sort.Sort(ByTimes(list))
	result := make([][]BmtUser, 0)
	i := 0
	var j int
	for {
		if i >= len(list) {
			break
		}
		for j = i + 1; j < len(list) && *list[i].Times == *list[j].Times; j++ {
		}
		result = append(result, list[i:j])
		i = j
	}
	return result
}

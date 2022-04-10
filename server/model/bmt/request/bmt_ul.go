package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/bmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BmtUlSearch struct {
	bmt.BmtUl
	request.PageInfo
}

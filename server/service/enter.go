package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/bmt"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	BmtServiceGroup     bmt.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

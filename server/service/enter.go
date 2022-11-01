package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/service/EmailMarketing"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	EmailmarketingServiceGroup EmailMarketing.ServiceGroup
	ContactsServiceGroup       Contacts.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

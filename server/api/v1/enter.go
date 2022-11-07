package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Automation"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/EmailMarketing"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	EmailmarketingApiGroup EmailMarketing.ApiGroup
	ContactsApiGroup       Contacts.ApiGroup
	SocialApiGroup         Social.ApiGroup
	AutomationApiGroup     Automation.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

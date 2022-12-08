package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Automation"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/router/EmailMarketing"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Scenario"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Webhook"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System         system.RouterGroup
	Example        example.RouterGroup
	Emailmarketing EmailMarketing.RouterGroup
	Contacts       Contacts.RouterGroup
	Social         Social.RouterGroup
	Automation     Automation.RouterGroup
	Webhook        Webhook.RouterGroup
	Scenario       Scenario.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Automation"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/router/EmailMarketing"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Social"
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
}

var RouterGroupApp = new(RouterGroup)

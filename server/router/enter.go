package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/router/EmailMarketing"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System         system.RouterGroup
	Example        example.RouterGroup
	Emailmarketing EmailMarketing.RouterGroup
	Contacts       Contacts.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

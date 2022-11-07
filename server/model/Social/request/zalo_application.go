package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ZaloApplicationSearch struct{
    Social.ZaloApplication
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

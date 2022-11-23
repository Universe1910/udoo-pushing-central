package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TriggerSearch struct{
    Automation.Trigger
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

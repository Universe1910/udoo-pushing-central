package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Scenario"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ScenarioSearch struct{
    Scenario.Scenario
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

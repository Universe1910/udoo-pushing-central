package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ContactTagsSearch struct{
    Contacts.ContactTags
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

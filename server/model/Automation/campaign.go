// 自动生成模板Campaign
package Automation

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
)

// Campaign 结构体
type Campaign struct {
	global.GVA_MODEL
	Name              string     `json:"name" form:"name" gorm:"column:name;comment:;"`
	StartAt           *time.Time `json:"startAt" form:"startAt" gorm:"column:start_at;comment:;"`
	EndAt             *time.Time `json:"endAt" form:"endAt" gorm:"column:end_at;comment:;"`
	Description       string     `json:"description" form:"description" gorm:"column:description;comment:;"`
	ParentID          *int       `json:"parentID" form:"parentID" gorm:"column:parent_id;comment:;"`
	CreatedBy         *int       `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:;"`
	Status            *bool      `json:"status" form:"status" gorm:"column:status;comment;"`
	TriggerId         *int       `json:"triggerId" form:"triggerId" gorm:"column:trigger_id;comment:;"`
	SequenceId        *int       `json:"sequenceId" form:"sequenceId" gorm:"column:sequence_id;comment:;"`
	ZaloApplicationID *int       `json:"zaloApplicationID" form:"zaloApplication" gorm:"column:zalo_application_id;comment:;"`

	// Fields   []*DocumentFields   `json:"fields" gorm:"many2many:document_field_references;foreignKey:ID;joinForeignKey:DocumentId;References:ID;joinReferences:FieldId;"`
	// ContactId
	Contacts []*Contacts.Contact `json:"Contacts" gorm:"many2many:contact_campaign;foreignKey:ID;joinForeignKey:CampaignId;References:ID;joinReferences:ContactId;"`

	TriggerObject   Trigger                `json:"triggerObject" form:"triggerObject" gorm:"foreignKey:TriggerId"`
	ZaloApplication Social.ZaloApplication `json:"zaloApplication" form:"zaloApplication" gorm:"foreignKey:ZaloApplicationID"`
	// CreatedObject sysModel.SysUser `json:"createdObject" form:"createdObject" gorm:"foreignKey:CreatedBy"`
}

// TableName Campaign 表名
func (Campaign) TableName() string {
	return "campaign"
}

package Contacts

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ContactRegister struct {
	global.GVA_MODEL
	Firstname  string `json:"firstname" form:"firstname" gorm:"column:firstname;comment:;"`
	Lastname   string `json:"lastname" form:"lastname" gorm:"column:lastname;comment:;"`
	Email      string `json:"email" form:"email" gorm:"column:email;comment:;"`
	Phone      string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Fbid       string `json:"fbid" form:"fbid" gorm:"column:fbid;comment:;"`
	Zid        string `json:"zid" form:"zid" gorm:"column:zid;comment:;"`
	Address    string `json:"address" form:"address" gorm:"column:address;comment:;"`
	City       string `json:"city" form:"city" gorm:"column:city;comment:;"`
	State      string `json:"state" form:"state" gorm:"column:state;comment:;"`
	Zipcode    string `json:"zipcode" form:"zipcode" gorm:"column:zipcode;comment:;"`
	Country    string `json:"country" form:"country" gorm:"column:country;comment:;"`
	CampaignID *uint  `json:"campaignId" form:"campaignId"`
}

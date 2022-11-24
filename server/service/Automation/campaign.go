package Automation

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Action"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	ServiceSocial "github.com/flipped-aurora/gin-vue-admin/server/service/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/service/zalo"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CampaignService struct {
}

// CreateCampaign Create Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) CreateCampaign(campaign Automation.Campaign) (err error) {
	err = global.GVA_DB.Create(&campaign).Error
	return err
}

// DeleteCampaign Delete Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) DeleteCampaign(campaign Automation.Campaign) (err error) {
	err = global.GVA_DB.Delete(&campaign).Error
	return err
}

// DeleteCampaignByIds 批量Delete Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) DeleteCampaignByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]Automation.Campaign{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateCampaign 更新Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) UpdateCampaign(campaign Automation.Campaign) (err error) {
	// db.Model(&job).Association('Skills').Replace(&job.Skills
	// delete(campaign, "Contacts")

	err = global.GVA_DB.Save(&campaign).Error
	return err
}

// GetCampaign 根据id获取Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) GetCampaign(id uint) (campaign Automation.Campaign, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("Contacts").Preload("TriggerObject").Preload("ZaloApplication").First(&campaign).Error
	return
}

// GetCampaignInfoList 分页获取Campaign记录
// Author [piexlmax](https://github.com/piexlmax)
func (campaignService *CampaignService) GetCampaignInfoList(info AutomationReq.CampaignSearch) (list []Automation.Campaign, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// Create db
	db := global.GVA_DB.Model(&Automation.Campaign{})
	var campaigns []Automation.Campaign
	// 如果有条件搜索 下方会自动Create 搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Contacts").Preload("TriggerObject").Preload("ZaloApplication").Find(&campaigns).Error
	return campaigns, total, err
}

func (CampaignService *CampaignService) DebugCampaign(id uint) (err error) {
	var campaign Automation.Campaign
	err = global.GVA_DB.Where("id = ?", id).Preload("Contacts").Preload("TriggerObject").Preload("ZaloApplication").First(&campaign).Error
	if err != nil {
		return err
	}
	var triggerFlow []map[string]interface{}
	err = json.Unmarshal([]byte(campaign.TriggerObject.Data), &triggerFlow)
	if err != nil {
		fmt.Println(err)
	}
	zaloApplication := campaign.ZaloApplication

	var zaloApplicationService ServiceSocial.ZaloApplicationService
	_, err = zaloApplicationService.FetchAccessToken(&zaloApplication)
	if err != nil {
		return err
	}

	var actionLog []map[string]interface{}
	for _, v := range triggerFlow {
		nexts := v["nexts"].([]interface{})
		for _, next := range nexts {
			trigger := next.(map[string]interface{})
			node := trigger["node"].(map[string]interface{})
			actionNode := Action.ActionNode{
				Label:      node["label"].(string),
				ActionName: node["actionName"].(string),
				ID:         node["id"].(string),
				Type:       node["type"].(string),
				Data:       node["data"],
			}
			if actionNode.ActionName == "action-send-zns" {
				znsTemplate := actionNode.Data.(map[string]interface{})
				for _, contact := range campaign.Contacts {
					res, err := CampaignService.ActionSendZNS(znsTemplate, zaloApplication, campaign, *contact)
					actionLog = append(actionLog, map[string]interface{}{
						"response": res,
						"err":      err,
					})
				}
			}
		}
	}
	fmt.Print(actionLog)
	return
}

func (campaignService *CampaignService) ActionSendZNS(znsTemplate map[string]interface{}, zaloApplication Social.ZaloApplication, campaign Automation.Campaign, contact Contacts.Contact) (res map[string]interface{}, err error) {
	dataReplaceList := znsTemplate["data"].([]interface{})
	templateData := make(map[string]interface{})
	for _, drl := range dataReplaceList {
		mia := drl.(map[string]interface{})
		key := mia["replaceKey"].(string)
		value := campaignService.ReplaceValue(mia["replaceVal"].(string), campaign, contact)
		templateData[key] = value
	}
	phone := contact.Phone
	zaloPhone := phone[1:]
	zaloPhone = "84" + zaloPhone
	templateId := znsTemplate["id"].(float64)
	body := map[string]interface{}{
		"mode":          "development",
		"phone":         zaloPhone,
		"template_id":   templateId,
		"template_data": templateData,
	}
	var zaloNotificationServiceApi zalo.ZaloNotificationServiceAPI
	zaloNotificationServiceApi.InitializeAPI()
	res, err = zaloNotificationServiceApi.SendZaloNotification(zaloApplication.AccessToken, body)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (campaignService *CampaignService) ReplaceValue(value string, campaign Automation.Campaign, contact Contacts.Contact) string {
	variables := []string{
		"__CID__", "__LASTNAME__", "__FIRSTNAME__", "__PHONE__", "__EMAIL__",
		"__CAMPAIGN_NAME__", "__CAMPAIGN_START__", "__CAMPAIGN_END__",
		"__ZALO_PHONE__", "__FB_ID__", "__ADDRESS__", "__CITY__", "__STATE__", "__ZIPCODE__", "__COUNTRY__",
		"__DATE__", "__TIME__",
	}

	currentTime := time.Now()

	mappingValue := map[string]string{
		"__CID__":            fmt.Sprint(contact.ID),
		"__LASTNAME__":       fmt.Sprint(contact.Lastname),
		"__FIRSTNAME__":      fmt.Sprint(contact.Firstname),
		"__PHONE__":          fmt.Sprint(contact.Phone),
		"__EMAIL__":          fmt.Sprint(contact.Email),
		"__ZALO_PHONE__":     fmt.Sprint(contact.Zid),
		"__FB_ID__":          fmt.Sprint(contact.Fbid),
		"__ADDRESS__":        fmt.Sprint(contact.Address),
		"__STATE__":          fmt.Sprint(contact.State),
		"__ZIPCODE__":        fmt.Sprint(contact.Zipcode),
		"__COUNTRY__":        fmt.Sprint(contact.Country),
		"__CAMPAIGN_NAME__":  fmt.Sprint(campaign.Name),
		"__CAMPAIGN_START__": fmt.Sprint(campaign.StartAt),
		"__CAMPAIGN_END__":   fmt.Sprint(campaign.EndAt),
		"__DATE__":           currentTime.Format("02-01-2006"),
		"__TIME__":           currentTime.Format("3:4:5 PM"),
	}

	exists, _ := utils.InArray(value, variables)

	if exists {
		return mappingValue[value]
	}
	tempVals := strings.Split(value, " ")
	if len(tempVals) > 1 {
		var result string
		for _, t := range tempVals {
			result = result + " " + campaignService.ReplaceValue(t, campaign, contact)
		}
		return strings.Trim(result, "")
	}
	return value
}

func (campaignService *CampaignService) RunTrigger(campaignId uint, contact Contacts.Contact) (err error) {
	var campaign Automation.Campaign
	err = global.GVA_DB.Where("id = ?", campaignId).Preload("TriggerObject").Preload("ZaloApplication").First(&campaign).Error
	if err != nil {
		return err
	}
	var triggerFlow []map[string]interface{}
	err = json.Unmarshal([]byte(campaign.TriggerObject.Data), &triggerFlow)
	if err != nil {
		fmt.Println(err)
	}
	zaloApplication := campaign.ZaloApplication

	var zaloApplicationService ServiceSocial.ZaloApplicationService
	_, err = zaloApplicationService.FetchAccessToken(&zaloApplication)
	if err != nil {
		return err
	}

	var actionLog []map[string]interface{}
	for _, v := range triggerFlow {
		nexts := v["nexts"].([]interface{})
		for _, next := range nexts {
			trigger := next.(map[string]interface{})
			node := trigger["node"].(map[string]interface{})
			actionNode := Action.ActionNode{
				Label:      node["label"].(string),
				ActionName: node["actionName"].(string),
				ID:         node["id"].(string),
				Type:       node["type"].(string),
				Data:       node["data"],
			}
			if actionNode.ActionName == "action-send-zns" {
				znsTemplate := actionNode.Data.(map[string]interface{})
				res, err := campaignService.ActionSendZNS(znsTemplate, zaloApplication, campaign, contact)
				actionLog = append(actionLog, map[string]interface{}{
					"response": res,
					"err":      err,
				})
			}
		}
	}

	return
}

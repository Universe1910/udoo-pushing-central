package Automation

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Action"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Automation"
	AutomationReq "github.com/flipped-aurora/gin-vue-admin/server/model/Automation/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Contacts"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	EmailUtils "github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/service/EmailMarketing"
	ServiceSocial "github.com/flipped-aurora/gin-vue-admin/server/service/Social"
	"github.com/flipped-aurora/gin-vue-admin/server/service/zalo"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gopkg.in/robfig/cron.v2"
	"gorm.io/gorm"
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
	err = global.GVA_DB.Where("id = ?", id).Preload("Contacts").Preload("TriggerObject").Preload("ZaloApplication").Preload("Sequences").Preload("Logs", func(db *gorm.DB) *gorm.DB {
		return db.Order("campaign_log.ID DESC").Limit(30)
	}).First(&campaign).Error
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
	err = db.Limit(limit).Offset(offset).Find(&campaigns).Error
	return campaigns, total, err
}

func (CampaignService *CampaignService) DebugCampaign(id uint) (res map[string]interface{}, err error) {
	var emailTemplateService EmailMarketing.EmailTemplateService
	var campaign Automation.Campaign
	err = global.GVA_DB.Where("id = ?", id).Preload("Contacts").Preload("TriggerObject").Preload("ZaloApplication").First(&campaign).Error
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: id,
			Action:     "Find Campaign",
			Message:    err.Error(),
			Type:       "Error",
			Source:     "Trigger",
		}
		CampaignService.CampaignLog(log)
		return nil, err
	}
	if !campaign.Status {
		log := Automation.CampaignLog{
			CampaignID: id,
			Action:     "Find Campaign",
			Message:    "campaing is not active",
			Type:       "Verbose",
			Source:     "Trigger",
		}
		CampaignService.CampaignLog(log)
		return nil, errors.New("campaing is not active")
	}
	var triggerFlow []map[string]interface{}
	err = json.Unmarshal([]byte(campaign.TriggerObject.Data), &triggerFlow)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: id,
			Action:     "Extract Trigger Flow",
			Message:    err.Error(),
			Type:       "Error",
		}
		CampaignService.CampaignLog(log)
	}
	zaloApplication := campaign.ZaloApplication
	var zaloApplicationService ServiceSocial.ZaloApplicationService
	_, err = zaloApplicationService.FetchAccessToken(&zaloApplication)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: id,
			Action:     "Fetch Zalo AccessToken",
			Message:    err.Error(),
			Type:       "Error",
			Source:     "Trigger",
		}
		CampaignService.CampaignLog(log)
		return nil, err
	}

	var lastResponse map[string]interface{}
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
			if actionNode.ActionName == "utils-wait" {
				data := actionNode.Data.(map[string]interface{})
				timeVal, err := strconv.Atoi(data["time"].(string))
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: id,
						Action:     "utils-wait",
						Message:    err.Error(),
						Type:       "Verbose",
					}
					CampaignService.CampaignLog(log)
					continue
				}
				time.Sleep(time.Duration(timeVal) * time.Second)
			}
			if actionNode.ActionName == "action-send-zns" {
				znsTemplate := actionNode.Data.(map[string]interface{})
				for _, contact := range campaign.Contacts {
					res, err := CampaignService.ActionSendZNS(znsTemplate, zaloApplication, campaign, *contact)
					if err != nil || (res["message"].(string) != "" && res["message"].(string) != "Success") {
						log := Automation.CampaignLog{
							CampaignID: id,
							Action:     "action-send-zns",
							Message:    res["message"].(string),
							Type:       "Verbose",
							ContactID:  contact.ID,
							Source:     "Trigger",
						}
						CampaignService.CampaignLog(log)
					} else {
						log := Automation.CampaignLog{
							CampaignID: id,
							Action:     "action-send-zns",
							Message:    res["message"].(string),
							ContactID:  contact.ID,
							Type:       "Success",
							Source:     "Trigger",
						}
						CampaignService.CampaignLog(log)
					}
					lastResponse := res
					fmt.Print(lastResponse)
				}
			}
			if actionNode.ActionName == "action-send-email" {
				var data = actionNode.Data.(map[string]interface{})
				emailId := uint(data["id"].(float64))
				emailTemplate, err := emailTemplateService.GetEmailTemplate(emailId)
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: id,
						Action:     "Get Email Template",
						Message:    err.Error(),
						Type:       "Error",
						Source:     "Trigger",
					}
					CampaignService.CampaignLog(log)
				}

				var emailContent map[string]interface{}
				err = json.Unmarshal([]byte(emailTemplate.Content), &emailContent)
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: id,
						Action:     "Extract email content",
						Message:    err.Error(),
						Type:       "Error",
						Source:     "Trigger",
					}
					CampaignService.CampaignLog(log)
				}
				contentHTML := emailContent["html"].(string)
				for _, contact := range campaign.Contacts {
					err := CampaignService.ActionSendEmail(contentHTML, emailTemplate.Subject, *contact, campaign)
					if err != nil {
						log := Automation.CampaignLog{
							CampaignID: id,
							Action:     "Action Send Email",
							Message:    err.Error(),
							Type:       "Error",
							Source:     "Trigger",
							ContactID:  contact.ID,
						}
						CampaignService.CampaignLog(log)

					} else {
						log := Automation.CampaignLog{
							CampaignID: id,
							Action:     "Action Send Email",
							Message:    "Success",
							Type:       "Success",
							Source:     "Trigger",
							ContactID:  contact.ID,
						}
						CampaignService.CampaignLog(log)
					}

				}
			}

		}
	}
	return lastResponse, err
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
	// "mode":          "development",
	body := map[string]interface{}{
		"phone":         zaloPhone,
		"template_id":   templateId,
		"template_data": templateData,
	}

	if zaloApplication.DevelopmentMode {
		body["mode"] = "development"
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
		"__DATE__", "__TIME__", "__CREATED_AT__",
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
		"__DATE__":           currentTime.Format("02-01-2006 15:04:05"),
		"__TIME__":           currentTime.Format("15:04:05"),
		"__CREATED_AT__":     fmt.Sprint(contact.CreatedAt),
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

func (campaignService *CampaignService) ActionSendEmail(body string, emailSubject string, contact Contacts.Contact, campaign Automation.Campaign) (err error) {
	err = EmailUtils.Email(contact.Email, emailSubject, body)
	if err != nil {
		return
	}
	return nil
}

//Triger run when add contact
func (campaignService *CampaignService) RunTrigger(campaignId uint, contact Contacts.Contact) (err error) {
	var campaign Automation.Campaign
	var emailTemplateService EmailMarketing.EmailTemplateService
	err = global.GVA_DB.Where("id = ?", campaignId).Preload("TriggerObject").Preload("ZaloApplication").First(&campaign).Error
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: campaignId,
			Action:     "Find Campaign",
			Message:    err.Error(),
			Type:       "Error",
			Source:     "Trigger",
		}
		campaignService.CampaignLog(log)
		return err
	}
	if !campaign.Status {
		log := Automation.CampaignLog{
			CampaignID: campaignId,
			Action:     "Find Campaign",
			Message:    "campaing is not active",
			Type:       "Verbose",
			Source:     "Trigger",
		}
		campaignService.CampaignLog(log)
		return errors.New("campaing is not active")
	}
	var triggerFlow []map[string]interface{}
	err = json.Unmarshal([]byte(campaign.TriggerObject.Data), &triggerFlow)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: campaignId,
			Action:     "Extract Trigger Flow",
			Message:    err.Error(),
			Type:       "Error",
		}
		campaignService.CampaignLog(log)
	}
	zaloApplication := campaign.ZaloApplication

	var zaloApplicationService ServiceSocial.ZaloApplicationService
	_, err = zaloApplicationService.FetchAccessToken(&zaloApplication)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: campaignId,
			Action:     "Fetch Zalo AccessToken",
			Message:    err.Error(),
			Type:       "Error",
			Source:     "Trigger",
		}
		campaignService.CampaignLog(log)
		return err
	}

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
			if actionNode.ActionName == "utils-wait" {
				data := actionNode.Data.(map[string]interface{})
				timeVal, err := strconv.Atoi(data["time"].(string))
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: campaignId,
						Action:     "utils-wait",
						Message:    err.Error(),
						Type:       "Verbose",
					}
					campaignService.CampaignLog(log)
					continue
				}
				time.Sleep(time.Duration(timeVal) * time.Second)
			}
			if actionNode.ActionName == "action-send-zns" {
				znsTemplate := actionNode.Data.(map[string]interface{})

				res, err := campaignService.ActionSendZNS(znsTemplate, zaloApplication, campaign, contact)
				if err != nil || (res["message"].(string) != "" && res["message"].(string) != "Success") {
					log := Automation.CampaignLog{
						CampaignID: campaignId,
						Action:     "action-send-zns",
						Message:    res["message"].(string),
						Type:       "Verbose",
						ContactID:  contact.ID,
						Source:     "Trigger",
					}
					campaignService.CampaignLog(log)
				} else {
					log := Automation.CampaignLog{
						CampaignID: campaignId,
						Action:     "action-send-zns",
						Message:    res["message"].(string),
						ContactID:  contact.ID,
						Type:       "Success",
						Source:     "Trigger",
					}
					campaignService.CampaignLog(log)
				}

			}
			if actionNode.ActionName == "action-send-email" {
				var data = actionNode.Data.(map[string]interface{})
				emailId := uint(data["id"].(float64))
				emailTemplate, err := emailTemplateService.GetEmailTemplate(emailId)
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: campaignId,
						Action:     "Get Email Template",
						Message:    err.Error(),
						Type:       "Error",
						Source:     "Trigger",
					}
					campaignService.CampaignLog(log)
				}
				var emailContent map[string]interface{}
				err = json.Unmarshal([]byte(emailTemplate.Content), &emailContent)
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: campaignId,
						Action:     "Extract email content",
						Message:    err.Error(),
						Type:       "Error",
						Source:     "Trigger",
					}
					campaignService.CampaignLog(log)
				}
				contentHTML := emailContent["html"].(string)
				err = campaignService.ActionSendEmail(contentHTML, emailTemplate.Subject, contact, campaign)
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: campaignId,
						Action:     "Action Send Email",
						Message:    err.Error(),
						Type:       "Error",
						Source:     "Trigger",
						ContactID:  contact.ID,
					}
					campaignService.CampaignLog(log)

				} else {
					log := Automation.CampaignLog{
						CampaignID: campaignId,
						Action:     "Action Send Email",
						Message:    "Success",
						Type:       "Success",
						Source:     "Trigger",
						ContactID:  contact.ID,
					}
					campaignService.CampaignLog(log)
				}

			}
		}
	}

	return err
}

func (campaignService *CampaignService) CampaignLog(log Automation.CampaignLog) {
	var logService CampaignLogService
	err := logService.CreateCampaignLog(log)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (CampaignService *CampaignService) DebugSequence(id uint) (res map[string]interface{}, err error) {
	var emailTemplateService EmailMarketing.EmailTemplateService
	var campaign Automation.Campaign
	err = global.GVA_DB.Where("id = ?", id).Preload("Contacts").Preload("Sequences").Preload("ZaloApplication").First(&campaign).Error
	if err != nil {
		return nil, err
	}
	if !campaign.Status {
		return nil, errors.New("campaing is not active")
	}
	sequences := campaign.Sequences
	sort.Slice(sequences, func(i, j int) bool {
		return sequences[i].Order < sequences[j].Order
	})

	//debug sequence
	for _, se := range sequences {
		if se.ActionName == "action-send-zns" {
			var znsTemplate map[string]interface{}
			err := json.Unmarshal([]byte(se.Data), &znsTemplate)
			if err != nil {
				log := Automation.CampaignLog{
					CampaignID: id,
					Action:     "Extract sequence data",
					Message:    fmt.Sprint("action data can not be unmarshaled for ", se, se.ActionName, ". Order ", se.Order),
					Type:       "Error",
					Source:     "Sequence",
				}
				CampaignService.CampaignLog(log)
				return nil, errors.New("action data can not be unmarshaled")
			}
			for _, contact := range campaign.Contacts {
				res, err := CampaignService.ActionSendZNS(znsTemplate, campaign.ZaloApplication, campaign, *contact)
				if err != nil || (res["message"].(string) != "" && res["message"].(string) != "Success") {
					log := Automation.CampaignLog{
						CampaignID: id,
						Action:     "action-send-zns",
						Message:    res["message"].(string),
						Type:       "Verbose",
						ContactID:  contact.ID,
						Source:     "Sequence",
					}
					CampaignService.CampaignLog(log)
				} else {
					log := Automation.CampaignLog{
						CampaignID: id,
						Action:     "action-send-zns",
						Message:    res["message"].(string),
						ContactID:  contact.ID,
						Type:       "Success",
						Source:     "Sequence",
					}
					CampaignService.CampaignLog(log)
				}
				lastResponse := res
				fmt.Print(lastResponse)
			}
		}
		if se.ActionName == "action-send-email" {
			var data map[string]interface{}
			err := json.Unmarshal([]byte(se.Data), &data)
			if err != nil {
				log := Automation.CampaignLog{
					CampaignID: id,
					Action:     "Extract sequence data",
					Message:    fmt.Sprint("action data can not be unmarshaled for ", err.Error()),
					Type:       "Error",
					Source:     "Sequence",
				}
				CampaignService.CampaignLog(log)
				return nil, err
			}
			emailId := uint(data["ID"].(float64))
			emailTemplate, err := emailTemplateService.GetEmailTemplate(emailId)
			if err != nil {
				log := Automation.CampaignLog{
					CampaignID: id,
					Action:     "Get Email Template",
					Message:    err.Error(),
					Type:       "Error",
					Source:     "Trigger",
				}
				CampaignService.CampaignLog(log)
			}
			var emailContent map[string]interface{}
			err = json.Unmarshal([]byte(emailTemplate.Content), &emailContent)
			if err != nil {
				log := Automation.CampaignLog{
					CampaignID: id,
					Action:     "Extract email content",
					Message:    err.Error(),
					Type:       "Error",
					Source:     "Trigger",
				}
				CampaignService.CampaignLog(log)
			}
			contentHTML := emailContent["html"].(string)
			for _, contact := range campaign.Contacts {
				err = CampaignService.ActionSendEmail(contentHTML, emailTemplate.Subject, *contact, campaign)
				if err != nil {
					log := Automation.CampaignLog{
						CampaignID: id,
						Action:     "Action Send Email",
						Message:    err.Error(),
						Type:       "Error",
						Source:     "Trigger",
						ContactID:  contact.ID,
					}
					CampaignService.CampaignLog(log)

				} else {
					log := Automation.CampaignLog{
						CampaignID: id,
						Action:     "Action Send Email",
						Message:    "Success",
						Type:       "Success",
						Source:     "Trigger",
						ContactID:  contact.ID,
					}
					CampaignService.CampaignLog(log)
				}

			}

		}
	}

	return nil, err
}

func (campaignService *CampaignService) RunSequence(id uint) (res map[string]interface{}, err error) {

	var campaign Automation.Campaign
	err = global.GVA_DB.Where("id = ?", id).Preload("Contacts").Preload("Sequences").Preload("ZaloApplication").First(&campaign).Error
	if err != nil {
		return nil, err
	}
	if !campaign.Status {
		return nil, errors.New("campaing is not active")
	}

	sequences := campaign.Sequences
	sort.Slice(sequences, func(i, j int) bool {
		return sequences[i].Order < sequences[j].Order
	})

	cronTask := cron.New()

	//debug sequence
	for _, se := range sequences {
		if se.ActionName == "action-send-zns" {
			// fmt.Print(se.Date.Format("04 15 02 01 *"))
			cronTask.AddFunc(se.Date.Format("04 15 02 01 *"), func() {
				campaignService.CreateSequenceSendZNS(se.ID)
			})
			// cronTask.AddFunc("0 5 * * * *", func() {
			// 	campaignService.CreateSequenceSendZNS(se.ID)
			// })
			// "0 30 * * * *"
		}
	}
	cronTask.Start()

	return nil, err
}

func (campaignService *CampaignService) CreateSequenceSendZNS(sequencId uint) {
	fmt.Println("CreateSequenceSendZNS")
	var sequenceService SequenceService
	sequence, err := sequenceService.GetSequence(sequencId)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Function: CreateSequenceSendZNS",
			Message:    "Can not get sequence ",
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}
	fmt.Println("get sequence")
	campaign, err := campaignService.GetCampaign(sequence.CampaignId)
	fmt.Println("get campaign")
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Function: CreateSequenceSendZNS",
			Message:    "Can not get campaign ",
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}

	if !campaign.Status {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Function: CreateSequenceSendZNS",
			Message:    "campaing is not active ",
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}
	fmt.Println("starting action")
	var znsTemplate map[string]interface{}
	err = json.Unmarshal([]byte(sequence.Data), &znsTemplate)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Extract sequence data",
			Message:    fmt.Sprint("action data can not be unmarshaled for ", sequence, sequence.ActionName, ". Order ", sequence.Order),
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}
	for _, contact := range campaign.Contacts {
		res, err := campaignService.ActionSendZNS(znsTemplate, campaign.ZaloApplication, campaign, *contact)
		if err != nil || (res["message"].(string) != "" && res["message"].(string) != "Success") {
			log := Automation.CampaignLog{
				CampaignID: sequence.CampaignId,
				Action:     "action-send-zns",
				Message:    res["message"].(string),
				Type:       "Verbose",
				ContactID:  contact.ID,
				Source:     "Sequence",
			}
			campaignService.CampaignLog(log)
		} else {
			log := Automation.CampaignLog{
				CampaignID: sequence.CampaignId,
				Action:     "action-send-zns",
				Message:    res["message"].(string),
				ContactID:  contact.ID,
				Type:       "Success",
				Source:     "Sequence",
			}
			campaignService.CampaignLog(log)
		}

	}
}

func (campaignService *CampaignService) CreateSequenceSendEmail(sequencId uint) {
	var emailTemplateService EmailMarketing.EmailTemplateService
	var sequenceService SequenceService
	sequence, err := sequenceService.GetSequence(sequencId)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Function: CreateSequenceSendEmail",
			Message:    "Can not get sequence ",
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}
	campaign, err := campaignService.GetCampaign(sequence.CampaignId)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Function: CreateSequenceSendEmail",
			Message:    "Can not get campaign ",
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}

	if !campaign.Status {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Function: CreateSequenceSend Email",
			Message:    "campaing is not active ",
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal([]byte(sequence.Data), &data)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Extract sequence data",
			Message:    fmt.Sprint("action data can not be unmarshaled for ", err.Error()),
			Type:       "Error",
			Source:     "Sequence",
		}
		campaignService.CampaignLog(log)
		return
	}
	emailId := uint(data["ID"].(float64))
	emailTemplate, err := emailTemplateService.GetEmailTemplate(emailId)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Get Email Template",
			Message:    err.Error(),
			Type:       "Error",
			Source:     "Trigger",
		}
		campaignService.CampaignLog(log)
	}
	var emailContent map[string]interface{}
	err = json.Unmarshal([]byte(emailTemplate.Content), &emailContent)
	if err != nil {
		log := Automation.CampaignLog{
			CampaignID: sequence.CampaignId,
			Action:     "Extract email content",
			Message:    err.Error(),
			Type:       "Error",
			Source:     "Trigger",
		}
		campaignService.CampaignLog(log)
	}
	contentHTML := emailContent["html"].(string)
	for _, contact := range campaign.Contacts {
		err := campaignService.ActionSendEmail(contentHTML, emailTemplate.Subject, *contact, campaign)
		if err != nil {
			log := Automation.CampaignLog{
				CampaignID: sequence.CampaignId,
				Action:     "Action Send Email",
				Message:    err.Error(),
				Type:       "Error",
				Source:     "Trigger",
				ContactID:  contact.ID,
			}
			campaignService.CampaignLog(log)

		} else {
			log := Automation.CampaignLog{
				CampaignID: sequence.CampaignId,
				Action:     "Action Send Email",
				Message:    "Success",
				Type:       "Success",
				Source:     "Trigger",
				ContactID:  contact.ID,
			}
			campaignService.CampaignLog(log)
		}

	}
}

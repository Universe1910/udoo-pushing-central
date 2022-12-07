package FacebookMessgener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	fbm "github.com/flipped-aurora/gin-vue-admin/server/model/FacebookMessgener"
)

type FacebookMessgenerService struct {
}
// https://graph.facebook.com/LATEST-API-VERSION/PAGE-ID/messages

const (
	FACEBOOK_API   = "https://graph.facebook.com/v15.0/me/messages?access_token=%s"
	IMAGE          = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
	VISIT_SHOW_URL = "http://labouardy.com"
)

func (fms *FacebookMessgenerService) ProcessMessage(event fbm.Messaging) {
	client := &http.Client{}
	var response fbm.Response
	// response = fbm.Response{
	// 	Recipient: fbm.User{
	// 		ID: event.Sender.ID,
	// 	},
	// 	Message: fbm.Message{
	// 		Attachment: &fbm.Attachment{
	// 			Type: "template",
	// 			Payload: fbm.Payload{
	// 				TemplateType: "generic",
	// 				// Elements:     BuildCarousel(shows[:10]),
	// 			},
	// 		},
	// 	},
	// }
	response = fbm.Response{
		Recipient: fbm.User{
			ID: event.Sender.ID,
		},
		Message: fbm.Message{
			Text: "Reply value",
		},
	}
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)
	pageToken := "EAAF51YITUEYBAC1UQU8UU7uvwKbQvlMQCuTTsl14cK60zfVk6LfAvjf8CxiW1wnv6ivdFtaMZAIY81WgRms87QygPYSbNZCahJk9wH1lCZCp2nmZC7FOvQfiAFjjtpBZBkqYtMZA9a6WAqLU0L4lXGmZBsoaRFTmWwZASCGGZAFr2OAdVpsaycatnhUO1fQRcZCQEkZBnyh1ChuFAZDZD"
	url := fmt.Sprintf(FACEBOOK_API, pageToken)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}

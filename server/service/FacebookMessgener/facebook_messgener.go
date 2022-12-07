package FacebookMessgener

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	pageToken := "EAAF51YITUEYBAOVaJUmp1CFt5pNcNrIUOc2Ji6anm8lj5AQZCYAYhSkEZB7TDcrbdktf1pkHwnOTBUX4pZCPgfghcqzsmGbckE6OK8uZBvXupVJuYWB9tn3YafQvMUXfr4y9KrXlxd8xC7DXmD69PeQSA09vLi0v7ZC0vH8amDAJSNXXOapmQFLIJkAr5BMI2VtS8Rtur8wZDZD"
	url := fmt.Sprintf(FACEBOOK_API, pageToken)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("after new request")
		fmt.Println(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("after client. do")
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
}

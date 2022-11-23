package zalo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ExecuteRequestService struct{}

func ExecuteRequest(apiBuilder map[string]interface{}) (map[string]interface{}, error) {
	method := apiBuilder["method"].(string)
	endpoint := apiBuilder["endpoint"].(string)
	header := apiBuilder["header"].(http.Header)
	client := &http.Client{}
	var req *http.Request
	var err error
	switch method {
	case "GET":
		{
			req, err = http.NewRequest(http.MethodGet, endpoint, nil)
			if err != nil {
				return nil, err
			}
			if header != nil {
				req.Header = header
			}
			break
		}

	case "POST":
		{
			body := apiBuilder["body"]
			jsonReq, err := json.Marshal(body)
			if body == nil {
				jsonReq = nil
			}
			req, err = http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonReq))
			if err != nil {
				return nil, err
			}
			if header != nil {
				req.Header = header
			}
			break
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	// fmt.Println(bodyString)
	result := make(map[string]interface{})
	json.Unmarshal([]byte(bodyString), &result)
	return result, nil
}

package hwpush

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/**
 * init
 */
func NewClient(clientID string, clientSecret string) *HuaweiPushClient {

	vers := &Vers{
		Ver:   "1",
		AppID: clientID,
	}
	nspCtx, _ := json.Marshal(vers)
	return &HuaweiPushClient{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		NspCtx:       string(nspCtx),
	}
}

/**
 * message init
 */
func NewDataMessage() *DataMessage {
	return &DataMessage{
		Message: Message{
			Android: Android{
				Notification: Notification{
					Title: "",
					Body:  "",
					ClickAction: ClickAction{
						Type:   1,
						Intent: "intent://com.bjhongren.redshop/hrsj?#Intent;scheme=pushscheme;launchFlags=0x4000000;i.type=0;end",
					},
				},
			},
			Token: []string{},
		},
	}
}

/**
 * form post
 */
func FormPost(url string, data url.Values) ([]byte, error) {
	u := ioutil.NopCloser(strings.NewReader(data.Encode()))
	r, err := http.Post(url, "application/x-www-form-urlencoded", u)
	if err != nil {

		return []byte(""), err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {

		return []byte(""), err
	}
	return b, err
}

/**
 * form post
 */
func FormPostPush(url string, data *DataMessage, accessToken string) ([]byte, error) {
	client := &http.Client{}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return []byte(""), err
	}
	fmt.Println(string(jsonData))
	body := bytes.NewBuffer(jsonData)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "Content-Type: application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {

		return []byte(""), err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return b, err
}

/**
 * get token
 */
func (this HuaweiPushClient) GetToken() string {
	reqUrl := TOKEN_URL
	param := make(url.Values)
	param["grant_type"] = []string{GRANTTYPE}
	param["client_id"] = []string{this.ClientId}
	param["client_secret"] = []string{this.ClientSecret}
	res, err := FormPost(reqUrl, param)

	if nil != err {
		return ""
	}
	var tokenRes = &TokenResStruct{}
	err = json.Unmarshal(res, tokenRes)
	if err != nil {
		return ""
	}
	return tokenRes.Access_token
}

/**
 * push msg
 */
func (this HuaweiPushClient) PushMsg(deviceToken, payload string) string {

	accessToken := this.GetToken()
	reqUrl := PUSH_URL + this.ClientId + "/messages:send"
	fmt.Println(reqUrl)

	// var originParam = map[string]string{
	// 	"nsp_svc":           NSP_SVC,
	// 	"nsp_ts":            strconv.Itoa(int(time.Now().Unix())),
	// 	"device_token_list": "[\"" + deviceToken + "\"]",
	// 	"payload":           payload,
	// 	"expire_time":       time.Now().Format("2006-01-02T15:04"),
	// }

	param := NewDataMessage()
	param.Message.Token = []string{deviceToken}
	param.Message.Android.Notification.Title = "华为推送测试"
	param.Message.Android.Notification.Body = "holle 红仁"

	// push
	res, _ := FormPostPush(reqUrl, param, accessToken)

	return string(res)
}

package vivopush

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * message init
 */
func NewVivoPushApp(appid string, appkey string, appSecret string, timestamp string, sign string) *VivoPushApp {
	return &VivoPushApp{
		AppId:     appid,
		AppKey:    appkey,
		AppSecret: appSecret,
		Timestamp: timestamp,
		Sign:      sign,
	}
}

func NewVivoPushMessage() *VivoPushMessage {
	return &VivoPushMessage{
		Alias:       "",
		NotifyType:  2,
		Title:       "",
		Content:     "",
		SkipType:    3,
		SkipContent: "",
		ClientCustomMap: ClientCustomMap{
			Status: 1,
		},
		RequestId:      "",
		Classification: 1,
	}
}

/**
 * form post
 */
func FormPostToken(url string, data *VivoPushApp) ([]byte, error) {
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

	req.Header.Set("Content-Type", "application/json;charset='utf-8'")
	req.Header.Set("Accept", "application/json")

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
 * form post
 */
func FormPostPush(url string, data *VivoPushMessage, accessToken string) ([]byte, error) {
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authToken", accessToken)

	resp, err := client.Do(req)
	if err != nil {

		return []byte(""), err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	fmt.Println(b)
	if err != nil {
		return []byte(""), err
	}

	return b, err
}

/**
 * get token
 */
func (this *VivoPushApp) GetToken() string {
	reqUrl := TOKEN_URL

	param := &VivoPushApp{
		AppId:     this.AppId,
		AppKey:    this.AppKey,
		Timestamp: this.Timestamp,
		Sign:      this.Sign,
	}
	res, err := FormPostToken(reqUrl, param)

	if nil != err {
		return ""
	}
	var tokenRes = &TokenResStruct{}
	err = json.Unmarshal(res, tokenRes)
	if err != nil {
		return ""
	}
	fmt.Printf("%#v", tokenRes)
	if tokenRes.Result == 0 {
		return tokenRes.AuthToken
	} else {
		return ""
	}

}

/**
 * push msg
 */
func (this *VivoPushApp) PushMsg(deviceToken, payload string) string {
	accessToken := this.GetToken()
	reqUrl := PUSH_URL
	fmt.Println(reqUrl)
	fmt.Println(accessToken)
	param := NewVivoPushMessage()
	param.Alias = "0b89b6ea-767e-8155-4832-1609746479528"
	param.Title = "vivo推送测试"
	param.Content = "hello 红仁"
	param.SkipContent = "vivo123456"
	param.RequestId = "324234234234"

	// push
	res, _ := FormPostPush(reqUrl, param, accessToken)

	return string(res)
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

package vivopush

// url
const (
	TOKEN_URL = "https://api-push.vivo.com.cn/message/auth"
	PUSH_URL  = "https://api-push.vivo.com.cn/message/send"
)

// config
const (
	GRANTTYPE = "client_credentials"
	NSP_SVC   = "openpush.message.api.send"
)

/**
 **************************************** 结构体
 */

type VivoPushApp struct {
	AppId     string `json:"appId"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
}

type VivoPushMessage struct {
	Alias           string `json:"alias"`
	NotifyType      int64  `json:"notifyType"`
	Title           string `json:"title"`
	Content         string `json:"content"`
	SkipType        int64  `json:"skipType"`
	SkipContent     string `json:"skipContent"`
	ClientCustomMap ClientCustomMap
	RequestId       string `json:"requestId"`
	Classification  int64  `json:"classification"`
}

type ClientCustomMap struct {
	Status int64 `json:"status"`
}

type TokenResStruct struct {
	Result    int64  `json:"result"`
	AuthToken string `json:"authToken"`
}

// func (this *Message) Json() string {
// 	bytes, err := json.Marshal(this)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(bytes)
// }

package hwpush

import (
	"encoding/json"
)

// url
const (
	TOKEN_URL = "https://oauth-login.cloud.huawei.com/oauth2/v2/token"
	PUSH_URL  = "https://push-api.cloud.huawei.com/v1/"
)

// config
const (
	GRANTTYPE = "client_credentials"
	NSP_SVC   = "openpush.message.api.send"
)

/**
 **************************************** 结构体
 */

type HuaweiPushClient struct {
	ClientId     string
	ClientSecret string
	NspCtx       string
}

type Vers struct {
	Ver   string `json:"ver"`
	AppID string `json:"appId"`
}

type TokenResStruct struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Token_type   string `json:"token_type"`
}

/**
 **************************************** 消息体
 */

type DataMessage struct {
	Message Message `json:"message"`
}

type Message struct {
	Android Android  `json:"android"`
	Token   []string `json:"token"`
}

type Android struct {
	Notification Notification `json:"notification"`
}
type Notification struct {
	Title       string      `json:"title"`
	Body        string      `json:"body"`
	ClickAction ClickAction `json:"click_action"`
}
type ClickAction struct {
	Type   int    `json:"type"`
	Intent string `json:"intent"`
}

// /**
//  **************************************** 封装
//  */

// func (this *Message) SetContent(content string) *Message {
// 	this.Hps.Msg.Body.Content = content
// 	return this
// }

// func (this *Message) SetTitle(title string) *Message {
// 	this.Hps.Msg.Body.Title = title
// 	return this
// }

// func (this *Message) SetIntent(intent string) *Message {
// 	this.Hps.Msg.Action.Param.Intent = intent
// 	return this
// }

// func (this *Message) SetAppPkgName(appPkgName string) *Message {
// 	this.Hps.Msg.Action.Param.AppPkgName = appPkgName
// 	return this
// }

// func (this *Message) SetExtAction(Action string) *Message {
// 	this.Hps.Ext.Action = Action
// 	return this
// }
// func (this *Message) SetExtFunc(Func string) *Message {
// 	this.Hps.Ext.Func = Func
// 	return this
// }
// func (this *Message) SetExtCollect(collect string) *Message {
// 	this.Hps.Ext.Collect = collect
// 	return this
// }
// func (this *Message) SetExtTitle(title string) *Message {
// 	this.Hps.Ext.Title = title
// 	return this
// }
// func (this *Message) SetExtContent(content string) *Message {
// 	this.Hps.Ext.Collect = content
// 	return this
// }

// func (this *Message) SetExtUrl(url string) *Message {
// 	this.Hps.Ext.Url = url
// 	return this
// }

func (this *Message) Json() string {
	bytes, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

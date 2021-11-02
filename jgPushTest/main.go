package main

import (
	"fmt"
	"gotest/01/jgPushTest/jpush"
)

func main() {
	client := jpush.NewJpushClient("90e7435db341ea78f3d3cab8", "22160bc2b9a53c46770dbfb0")
	cidList, err := client.PushCid(1, "push")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cidList)

	payload := &jpush.Payload{
		Platform: jpush.NewPlatform().All(),
		Audience: jpush.NewAudience().All().SetAlias("zghrds242823"),
		Notification: &jpush.Notification{
			// Alert: "test1",
			Android: &jpush.AndroidNotification{
				Alert: "红仁世界测试内容111",
				Title: "红仁商城测试222",
			},
		},
		Options: &jpush.Options{
			TimeLive:       60,
			ApnsProduction: false,
		},
		Cid: cidList[0],
	}
	msgID, err := client.Push(payload)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msgID)

	// msgId, err = client.PushValidate(payload)
}

package main

import (
	"fmt"
	"gotest/01/vivoPushTest/vivopush"
	"strconv"
	"time"
)

func main() {
	appId := "103925069"
	appKey := "7f85ea1006d512864728b3a9d57903dc"
	appSecret := "b3a11cc7-ab23-4ae3-978b-2d276b285be0"
	timestamp := strconv.Itoa(int(time.Now().Unix() * 1000))
	sign := vivopush.Md5(appId + appKey + timestamp + appSecret)
	vivoApp := vivopush.NewVivoPushApp(appId, appKey, appSecret, timestamp, sign)

	fmt.Println(vivoApp)
	vivoRes := vivoApp.PushMsg("IQAAAAC", "hello 红仁")

	fmt.Println(vivoRes)
	// fmt.Printf("%#v", vivoApp)
}

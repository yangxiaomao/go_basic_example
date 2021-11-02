package main

import (
	"fmt"
	"gotest/01/huaweiPush/hwpush"
)

func main() {
	hwClient := hwpush.NewClient("102629769", "822f69e3cbe3e72399e95cd5f3e118b6216cadbec98ca9b4c8ed7378711bba8c")
	hwToken := hwClient.GetToken()
	hwRes := hwClient.PushMsg("IQAAAACy0Yv0AACr5-FcLyETniumUFuwbOcgKUvAklHtl3t3WkCX76A3kbieGUcZPxsKOTGPyoA8WBypq3DxVlKmZu3G-y9IS0ZCinoKiojBoTXpjg", "hello 红仁")
	fmt.Println(hwRes)
	fmt.Printf("%#v", hwClient)
	fmt.Printf("%#v", hwToken)
}

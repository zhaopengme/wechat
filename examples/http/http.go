package main

import (
	"fmt"
	"github.com/zhaopengme/wechat/message"
	"net/http"

	"github.com/zhaopengme/wechat"
)

func hello(rw http.ResponseWriter, req *http.Request) {

	//配置微信参数
	config := &wechat.Config{
		AppID:          "wxbb07c629e06e6c2c",
		AppSecret:      "18dfd6262488279413d45ef2eb2fb750",
		Token:          "pmewebchat",
		EncodingAESKey: "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG",
	}
	wc := wechat.NewWechat(config)
	// 传入request和responseWriter
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}

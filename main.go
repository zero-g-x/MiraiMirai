package main

import (
	"MiraiMirai/mirai"
	"github.com/Logiase/gomirai/bot"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main(){
	//从键盘监听结束信号(win下：ctrl+c
	interrupt:=make(chan os.Signal,1)
	signal.Notify(interrupt,os.Interrupt)

	//初始化bot
	mirai.Client = bot.NewClient("default",mirai.URL,mirai.Authkey)
	mirai.Client.Logger.Level=logrus.TraceLevel
	key,err:=mirai.Client.Auth()
	if err!=nil{
		mirai.Client.Logger.Fatal(err)
	}

	mirai.MyBot,err=mirai.Client.Verify(mirai.QQ,key)
	if err!=nil{
		mirai.Client.Logger.Fatal(err)
	}
	defer func(){
		_=mirai.Client.Release(mirai.QQ)

	}()

	//启动
	go mirai.BotReminder()

	<-interrupt
}
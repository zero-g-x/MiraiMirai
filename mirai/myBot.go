package mirai

import (
	"github.com/Logiase/gomirai/bot"
	"github.com/Logiase/gomirai/message"
	"time"
)

const (
	//targetGroup = 592120911//目标群组
	targetFriend = 1343244602
	QQ = 2974953317//bot账号
	URL = "http://127.0.0.1:8080"
	Authkey = "1234567890"
)
var Client *bot.Client=new(bot.Client)
var MyBot *bot.Bot=new(bot.Bot)

var LastTime =""
func BotReminder(){
	for true{
		now:=time.Now()
		nextMinute := time.Date(now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),0,
			0,now.Location()).Add(time.Minute)
		delta := nextMinute.Unix()-now.Unix()//时间间隔
		time.Sleep(time.Duration(delta)*time.Second)//间隔1min

		newFeed := GetNewFeed()
		//发现更新
		newTime:=newFeed.Entries[0].Updated
		if LastTime!=newTime{
			LastTime=newTime
			title:=newFeed.Entries[0].Title
			author:=newFeed.Entries[0].Author.Name
			url:=newFeed.Entries[0].ID

			_,err:=MyBot.SendFriendMessage(targetFriend,0,
				message.PlainMessage("更新："+title+"\n"+"作者："+author+"\n"+url))
			if err!=nil{
				Client.Logger.Error(err)
			}
		}
		/*_,_=MyBot.SendFriendMessage(targetFriend,0,
			message.PlainMessage("hello"))
*/
	}


}
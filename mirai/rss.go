package mirai

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//根据rss的xml结构定义结构体

type AuthorInfo struct{
	XMLName xml.Name `xml:"author"`
	Name string `xml:"name"`
}

type Entry struct{
	XMLName xml.Name`xml:"entry"`
	Title string `xml:"title"`
	ID string `xml:"id"`
	Updated string `xml:"updated"`
	Author AuthorInfo `xml:"author"`
}

type Feed struct{
	XMLName xml.Name `xml:"feed"`
	Title string `xml:"title"`
	Subtitle string `xml:"subtitle"`
	ID string `xml:"id"`
	Updated string `xml:"updated"`
	Entries [100]Entry `xml:"entry"`
}

//获取站点的xml
func GetXml(url string)[]byte{
	client := &http.Client{}
	req,_:=http.NewRequest("GET",url,nil)

	res,err:=client.Do(req)
	if err!=nil{
		fmt.Println("http get error",err)
	}
	defer res.Body.Close()

	var body []byte
	body,err=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("read err",err)
	}
	return body
}

//获得最近更新的博客内容
func GetNewFeed()Feed{
	var newFeed Feed
	rssURL := "https://0xffff.one/atom"
	xmlContent := GetXml(rssURL)
	err := xml.Unmarshal(xmlContent,&newFeed)
	if err!=nil{
		fmt.Println(err)
	}

	return newFeed
}
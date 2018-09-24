package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GET(url string, cookie string) string {
	fmt.Println("url:", url)
	client := &http.Client{Timeout: time.Second * 1}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//Add 头协议
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Cookie", cookie)
	// reqest.Header.Add("Referer", url)
	// reqest.Header.Add("Host", "info.win0168.com")

	// reqest.Header.Add("Cookie", "UM_distinctid=16574bdb87b90d-09607a7cf6df6a-34677908-13c680-16574bdb87c53f; Bet007EuropeIndex_Cookie=0^1^1^1; win007BfCookie=null; bfWin007FirstMatchTime=2018,8,7,08,00,0")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36")
	response, err := client.Do(reqest) //提交

	if nil != err {
		fmt.Errorf("Get Error%v", err)
	}
	// defer response.Body.Close()
	// cookies := response.Cookies() //遍历cookies
	// for _, cookie := range cookies {
	// 	fmt.Println("cookie:", cookie)
	// }

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
		// fmt.Println("read body", err1)
	}
	return string(body) //网页源码

}

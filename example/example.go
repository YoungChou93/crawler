package main

import (

	"time"
	"github.com/YoungZhou93/crawler/common/crawler_resquest"
	"github.com/YoungZhou93/crawler/controller"
)

func main() {
	crawler := controller.NewController(3)
	request := crawler_request.NewCrawlerRequest("http://photo.poco.cn/jiqiao/", "Get")
	request.SetLabel("")
	crawler.SetRequest(request)
	crawler.Stop()
	crawler.Run()
	crawler.Restart()
	for i:=0;i<10;i++ {
		time.Sleep(500 * time.Millisecond)
	}
	crawler.Stop()

}
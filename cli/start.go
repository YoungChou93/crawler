package cli

import (
	"github.com/urfave/cli"
	"log"
	"github.com/YoungChou93/crawler/controller"
	"github.com/YoungChou93/crawler/common/crawler_resquest"
	"fmt"
	"time"
	"github.com/YoungChou93/crawler/crawler"
)
var (
	aCrawler *crawler.Crawler
)
func start(c *cli.Context){
	if len(c.Args()) != 0 {
		log.Fatalf("the `create` command takes no arguments. See '%s create --help'.", c.App.Name)
	}
	aCrawler = crawler.NewCrawler()
	controller := controller.NewController(3);
	request := crawler_request.NewCrawlerRequest("http://photo.poco.cn/jiqiao/", "Get")
	request.SetLabel("")
	controller.SetRequest(request)
	go func() {
		controller.Run()
	}()
	for {
		time.Sleep(1000*1000*1000)
		fmt.Println(controller.GetSchedulerLength())
	}

}

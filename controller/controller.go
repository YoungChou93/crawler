package controller

import (
	"github.com/YoungChou93/crawler/common/crawler_resquest"
	"github.com/YoungChou93/crawler/downloader"
	"github.com/YoungChou93/crawler/scheduler"
	"github.com/YoungChou93/crawler/page_analyzer"
	"github.com/YoungChou93/crawler/outputer"
	"github.com/YoungChou93/crawler/common/thread_manage"
	"time"
)

type Controller struct {

	//请求队列组件
	scheduler scheduler.Scheduler
	//下载组件
	downloader downloader.Downloader
	//页面分析组件
	pageAnalyzer page_analyzer.PageAnalyzer
	//输出组件
	outputer outputer.Outputer

	//线程管理组件
	threadManager thread_manage.ThreadManager
	threadTotal uint

	//程序状态

}

func NewDefaultController()*Controller{
	return NewController(1)
}

func NewController(threadTotal uint)*Controller{

	scheduler:=scheduler.NewListScheduler()
	downloader:=downloader.NewHttpDownloader()
	pageAnalyzer:=page_analyzer.NewDefaultAnalyer()
	outputer:=outputer.NewOutputerPrint()
	threadManager:=thread_manage.NewChanThreadManager(threadTotal)
	return &Controller{scheduler,downloader,pageAnalyzer,outputer,threadManager,threadTotal}

}


func (this *Controller) SetUrl(url string) {
	request := crawler_request.NewCrawlerRequest(url, "GET")
	this.SetRequest(request)
}

func (this *Controller) SetUrls(urls []string){
	for _,url:= range urls{
		request := crawler_request.NewCrawlerRequest(url, "GET")
		this.SetRequest(request)
	}
}

func (this *Controller) SetRequest(request *crawler_request.CrawlerRequest) {
	this.scheduler.PushBack(request)
}

func (this *Controller) SetRequests(requests []*crawler_request.CrawlerRequest) {
	for _,request:= range requests{
		this.scheduler.PushBack(request)
	}
}

func (this *Controller)SetThreadManager(threadManager thread_manage.ThreadManager)*Controller{
	this.threadManager=threadManager
	return this
}

func (this *Controller)SetScheduler(scheduler scheduler.Scheduler)*Controller{
	this.scheduler=scheduler
	return this
}

func (this *Controller)SetDownloader(downloader downloader.Downloader)*Controller{
	this.downloader=downloader
	return this
}

func (this *Controller)SetPageAnalyzer( pageAnalyzer page_analyzer.PageAnalyzer)*Controller{
	this.pageAnalyzer=pageAnalyzer
	return this
}

func (this *Controller)SetOutputer(outputer outputer.Outputer)*Controller{
	this.outputer=outputer
	return this
}

func(this *Controller)GetSchedulerLength()int{
	return this.scheduler.Length()
}
func(this *Controller)Run(){
	for {

		//从队列中取出一个请求
		request:=this.scheduler.PopFront()

		//time.Sleep()

		if request==nil && this.threadManager.Left()==this.threadManager.Total(){
			this.pageAnalyzer.AnalyzeFinsh()
			break
		}else if request==nil{
			time.Sleep(500 * time.Millisecond)
			continue
		}
		//可用线程数减一
                this.threadManager.Get()

		go func(request *crawler_request.CrawlerRequest) {
			defer this.threadManager.Free()
			page:=this.downloader.Download(request)
			this.pageAnalyzer.Analyze(page)
			this.SetRequests(page.GetTargetRequests())
			this.outputer.Output(page)
		}(request)

	}
}

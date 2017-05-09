package scheduler

import "github.com/YoungChou93/crawler/common/crawler_resquest"

type Scheduler interface {
	PushBack(request * crawler_request.CrawlerRequest)
	PopFront()(request * crawler_request.CrawlerRequest)
	Length()(int)
}

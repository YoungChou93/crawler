package downloader

import (

	"github.com/YoungChou93/crawler/common/crawler_resquest"
	"github.com/YoungChou93/crawler/common/page"
)

type Downloader interface {
	Download(request *crawler_request.CrawlerRequest)*page.Page
}
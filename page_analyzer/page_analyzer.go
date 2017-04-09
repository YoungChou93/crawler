package page_analyzer

import (
	"github.com/YoungZhou93/crawler/common/page"
)

type PageAnalyzer interface {
	Analyze(page *page.Page)
	AnalyzeFinsh()
}


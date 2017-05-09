package page_analyzer

import (
	"github.com/YoungChou93/crawler/common/page"
)

type PageAnalyzer interface {
	Analyze(page *page.Page)
	AnalyzeFinsh()
}


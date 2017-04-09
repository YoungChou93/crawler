package outputer

import (

)
import "github.com/YoungZhou93/crawler/common/page"

type Outputer interface {

	Output(page *page.Page)
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tool"
)

type ToolSourceSearch struct {
	tool.ToolSource
	request.PageInfo
}

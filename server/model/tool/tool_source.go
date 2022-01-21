// 自动生成模板ToolSource
package tool

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ToolSource 结构体
// 如果含有time.Time 请自行import time包
type ToolSource struct {
	global.GVA_MODEL
	VoucherIndex  string `json:"voucherIndex" form:"voucherIndex" gorm:"column:voucher_index;comment:;size:50;"`
	VoucherSource string `json:"voucherSource" form:"voucherSource" gorm:"column:voucher_source;comment:;size:50;"`
}

// TableName ToolSource 表名
func (ToolSource) TableName() string {
	return "tool_source"
}

package tool

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tool"
	toolReq "github.com/flipped-aurora/gin-vue-admin/server/model/tool/request"
)

type ToolSourceService struct {
}

// CreateToolSource 创建ToolSource记录
// Author [ZHY](https://github.com/Beian27)
func (toolSourceService *ToolSourceService) CreateToolSource(toolSource tool.ToolSource) (err error) {
	err = global.GVA_DB.Create(&toolSource).Error
	return err
}

// DeleteToolSource 删除ToolSource记录
// Author [ZHY](https://github.com/Beian27)
func (toolSourceService *ToolSourceService) DeleteToolSource(toolSource tool.ToolSource) (err error) {
	err = global.GVA_DB.Delete(&toolSource).Error
	return err
}

// DeleteToolSourceByIds 批量删除ToolSource记录
// Author [ZHY](https://github.com/Beian27)
func (toolSourceService *ToolSourceService) DeleteToolSourceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]tool.ToolSource{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateToolSource 更新ToolSource记录
// Author [ZHY](https://github.com/Beian27)
func (toolSourceService *ToolSourceService) UpdateToolSource(toolSource tool.ToolSource) (err error) {
	err = global.GVA_DB.Save(&toolSource).Error
	return err
}

// GetToolSource 根据id获取ToolSource记录
// Author [ZHY](https://github.com/Beian27)
func (toolSourceService *ToolSourceService) GetToolSource(id uint) (err error, toolSource tool.ToolSource) {
	err = global.GVA_DB.Where("id = ?", id).First(&toolSource).Error
	return
}

// GetToolSourceInfoList 分页获取ToolSource记录
// Author [ZHY](https://github.com/Beian27)
func (toolSourceService *ToolSourceService) GetToolSourceInfoList(info toolReq.ToolSourceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tool.ToolSource{})
	var toolSources []tool.ToolSource
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&toolSources).Error
	return err, toolSources, total
}

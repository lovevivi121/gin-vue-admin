// 自动生成模板Banner
package fish

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Banner 结构体
type Banner struct {
      global.GVA_MODEL
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:;"`
      Path  string `json:"path" form:"path" gorm:"column:path;comment:;"`
}


// TableName Banner 表名
func (Banner) TableName() string {
  return "banner"
}


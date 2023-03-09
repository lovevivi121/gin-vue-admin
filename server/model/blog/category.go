// 自动生成模板Category
package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Category 结构体
type Category struct {
      global.GVA_MODEL
      CategoryName  string `json:"categoryName" form:"categoryName" gorm:"column:category_name;comment:分类名称;size:100;"`
}


// TableName Category 表名
func (Category) TableName() string {
  return "category"
}


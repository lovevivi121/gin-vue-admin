// 自动生成模板News
package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// News 结构体
type News struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;size:255;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;size:255;"`
      ReadCount  *int `json:"readCount" form:"readCount" gorm:"column:read_count;comment:;size:19;"`
      CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:;size:19;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:19;"`
}


// TableName News 表名
func (News) TableName() string {
  return "news"
}


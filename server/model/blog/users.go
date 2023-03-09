// 自动生成模板User
package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// User 结构体
type User struct {
      global.GVA_MODEL
      Sno  string `json:"sno" form:"sno" gorm:"column:sno;comment:;size:255;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;size:255;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:;size:255;"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:255;"`
      Telephone  string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:;size:255;"`
}


// TableName User 表名
func (User) TableName() string {
  return "user"
}


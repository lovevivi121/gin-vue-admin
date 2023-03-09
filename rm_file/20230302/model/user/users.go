// 自动生成模板Users
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Users 结构体
type Users struct {
      global.GVA_MODEL
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;size:255;"`
      Sno  string `json:"sno" form:"sno" gorm:"column:sno;comment:;size:255;"`
      Telephone  string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:;size:255;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:19;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:;size:255;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName Users 表名
func (Users) TableName() string {
  return "users"
}


package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
    "gorm.io/gorm"
)

type UserService struct {
}

// CreateUser 创建User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user blog.User) (err error) {
	err = global.GVA_DB.Create(&user).Error
	return err
}

// DeleteUser 删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)DeleteUser(user blog.User) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&blog.User{}).Where("id = ?", user.ID).Update("deleted_by", user.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&user).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteUserByIds 批量删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)DeleteUserByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&blog.User{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&blog.User{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateUser 更新User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)UpdateUser(user blog.User) (err error) {
	err = global.GVA_DB.Save(&user).Error
	return err
}

// GetUser 根据id获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)GetUser(id uint) (user blog.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&user).Error
	return
}

// GetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)GetUserInfoList(info blogReq.UserSearch) (list []blog.User, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&blog.User{})
    var users []blog.User
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&users).Error
	return  users, total, err
}

package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
    "gorm.io/gorm"
)

type UsersService struct {
}

// CreateUsers 创建Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (usersService *UsersService) CreateUsers(users user.Users) (err error) {
	err = global.GVA_DB.Create(&users).Error
	return err
}

// DeleteUsers 删除Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (usersService *UsersService)DeleteUsers(users user.Users) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&user.Users{}).Where("id = ?", users.ID).Update("deleted_by", users.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&users).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteUsersByIds 批量删除Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (usersService *UsersService)DeleteUsersByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&user.Users{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&user.Users{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateUsers 更新Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (usersService *UsersService)UpdateUsers(users user.Users) (err error) {
	err = global.GVA_DB.Save(&users).Error
	return err
}

// GetUsers 根据id获取Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (usersService *UsersService)GetUsers(id uint) (users user.Users, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&users).Error
	return
}

// GetUsersInfoList 分页获取Users记录
// Author [piexlmax](https://github.com/piexlmax)
func (usersService *UsersService)GetUsersInfoList(info userReq.UsersSearch) (list []user.Users, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&user.Users{})
    var userss []user.Users
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&userss).Error
	return  userss, total, err
}

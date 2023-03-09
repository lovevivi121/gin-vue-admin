package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
)

type CategoryService struct {
}

// CreateCategory 创建Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService) CreateCategory(category blog.Category) (err error) {
	err = global.GVA_DB.Create(&category).Error
	return err
}

// DeleteCategory 删除Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)DeleteCategory(category blog.Category) (err error) {
	err = global.GVA_DB.Delete(&category).Error
	return err
}

// DeleteCategoryByIds 批量删除Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)DeleteCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]blog.Category{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCategory 更新Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)UpdateCategory(category blog.Category) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

// GetCategory 根据id获取Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)GetCategory(id uint) (category blog.Category, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

// GetCategoryInfoList 分页获取Category记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *CategoryService)GetCategoryInfoList(info blogReq.CategorySearch) (list []blog.Category, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&blog.Category{})
    var categorys []blog.Category
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return  categorys, total, err
}

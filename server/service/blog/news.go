package blog

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/blog"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    blogReq "github.com/flipped-aurora/gin-vue-admin/server/model/blog/request"
)

type NewsService struct {
}

// CreateNews 创建News记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsService *NewsService) CreateNews(news blog.News) (err error) {
	err = global.GVA_DB.Create(&news).Error
	return err
}

// DeleteNews 删除News记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsService *NewsService)DeleteNews(news blog.News) (err error) {
	err = global.GVA_DB.Delete(&news).Error
	return err
}

// DeleteNewsByIds 批量删除News记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsService *NewsService)DeleteNewsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]blog.News{},"id in ?",ids.Ids).Error
	return err
}

// UpdateNews 更新News记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsService *NewsService)UpdateNews(news blog.News) (err error) {
	err = global.GVA_DB.Save(&news).Error
	return err
}

// GetNews 根据id获取News记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsService *NewsService)GetNews(id uint) (news blog.News, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&news).Error
	return
}

// GetNewsInfoList 分页获取News记录
// Author [piexlmax](https://github.com/piexlmax)
func (newsService *NewsService)GetNewsInfoList(info blogReq.NewsSearch) (list []blog.News, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&blog.News{})
    var newss []blog.News
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&newss).Error
	return  newss, total, err
}

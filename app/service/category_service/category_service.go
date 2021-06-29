package category_service

import (
	"code.aliyun.com/netroby/gosense/pkg/common"
	"code.aliyun.com/netroby/gosense/app/orm/model"
	"errors"
	"gorm.io/gorm"
)

func GetCategoryName(cateId int64) string {
	var cate model.Category
	result := common.NewDb.First(&cate, cateId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "默认"
	}
	return cate.Name
}

func GetCategories() []model.Category {
	var categories []model.Category

	result := common.NewDb.Find(&categories)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return categories
}

func GetCategoriesAsMap() map[int64]string {
	categories := GetCategories()
	cm := make(map[int64]string)
	for _, iv := range categories {
		cm[iv.ID] = iv.Name
	}
	return cm
}
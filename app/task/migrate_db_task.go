package task

import (
	"github.com/cnmade/bsmi-kb/app/orm/model"
	"github.com/cnmade/bsmi-kb/pkg/common"
)

func MigrateDbTask() {

	go func() {
		var al []model.Article
		common.NewDb.Limit(5000).Find(&al)
		for _, v := range al {
			common.PDB.Create(&v)
		}
	}()

	go func() {
		var al []model.ArticleHistory
		common.NewDb.Limit(5000).Find(&al)
		for _, v := range al {
			common.PDB.Create(&v)
		}
	}()

	go func() {
		var al []model.Category
		common.NewDb.Limit(5000).Find(&al)
		for _, v := range al {
			common.PDB.Create(&v)
		}
	}()

	go func() {
		var al []model.Tag
		common.NewDb.Limit(5000).Find(&al)
		for _, v := range al {
			common.PDB.Create(&v)
		}
	}()

}
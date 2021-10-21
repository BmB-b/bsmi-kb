package common

import (
	"github.com/cnmade/bsmi-kb/app/orm/model"
	"github.com/cnmade/bsmi-kb/pkg/common/vo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GetPDB(config *vo.AppConfig) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		//DSN: "user=postgres password=postgres dbname=gosense host=127.0.0.1 port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		DSN:                  config.Dbdsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}),
		&gorm.Config{

			Logger: newLogger,
		})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.Category{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.Tag{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.Article{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.ArticleHistory{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.TwoAuth{})

	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.FailBan{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

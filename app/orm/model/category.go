package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

func (Category) TableName() string {
	return "bk_category"

}

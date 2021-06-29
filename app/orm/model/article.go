package model
import  "gorm.io/datatypes"
type Article struct {
	Aid int64
	CateId int64 `gorm:"default:0"`
	Title string
	Content string
	PublishTime string
	PublishStatus int
	Views int64
	TagIds datatypes.JSON
}


func (Article) TableName()  string {
	return "gs_article"
	
}
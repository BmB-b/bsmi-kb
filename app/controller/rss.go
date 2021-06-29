package controller

import (
	"code.aliyun.com/netroby/gosense/app/orm/model"
	"code.aliyun.com/netroby/gosense/pkg/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"time"
)

// RSS controll group
type RSS struct {
}

func (rss *RSS) Alter(c *gin.Context) {
	c.Redirect(301, "/rss")
}

// Out Render and output RSS
//TODO 增加缓存
func (rss *RSS) Out(c *gin.Context) {
	var blogItems []model.Article
	result := common.NewDb.Limit(20).Order("aid desc").Find(&blogItems)
	if result.Error != nil {
		common.LogError(result.Error)
		return
	}
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "HardCoder",
		Link:        &feeds.Link{Href: "https://www.netroby.com"},
		Description: "Opensource , linux, golang",
		Created:     now,
	}
	feed.Items = make([]*feeds.Item, 0)
	for _, blog := range blogItems {

		itemTime, _ := time.Parse("2006-01-02 15:04:05", blog.PublishTime)
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       blog.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("https://www.netroby.com/view/%d", blog.Aid)},
			Description: blog.Content,
			Created:     itemTime,
		})
	}
	c.XML(http.StatusOK, (&feeds.Atom{feed}).FeedXml())
}

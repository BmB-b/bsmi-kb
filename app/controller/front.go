package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cnmade/bsmi-kb/app/orm/model"
	"github.com/cnmade/bsmi-kb/app/service/category_service"
	"github.com/cnmade/bsmi-kb/app/service/tag_service"
	"github.com/cnmade/bsmi-kb/app/vo"
	"github.com/cnmade/bsmi-kb/pkg/common"
	vo2 "github.com/cnmade/bsmi-kb/pkg/common/vo"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func UserPermissionCheck(c *gin.Context) (err error, username interface{}, isAdmin interface{}) {
	session := sessions.Default(c)
	username = session.Get("username")
	isAdmin = session.Get("isAdmin")
	common.Sugar.Infof("username was: %s", username.(string))
	if username == nil {

		return errors.New("需要登录"), nil, nil
	}
	return nil, username, isAdmin
}

func UserPermissionCheckDefaultAllow(c *gin.Context) (err error, username interface{}, isAdmin interface{}) {
	session := sessions.Default(c)
	username = session.Get("username")
	if username == nil {
		username = ""
	}
	isAdmin = session.Get("isAdmin")
	if isAdmin == nil {
		isAdmin = ""
	}
	common.Sugar.Infof("username was: %s", username.(string))
	common.Sugar.Infof("isAdmin was: %s", isAdmin.(string))
	return nil, username, isAdmin
}

type FrontController struct {
}

func (fc *FrontController) DemoPongoCtr(c *gin.Context) {
	c.HTML(200, "demo-pongo.html",
		common.Pongo2ContextWithVersion(pongo2.Context{"hello": "world"}),
	)
}

func (fc *FrontController) AboutCtr(c *gin.Context) {
	var Config = common.GetConfig()

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)

	c.HTML(200, "about.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        Config.Site_name,
			"siteDescription": common.Config.Site_description,
			"username":        username.(string),
			"isAdmin":         isAdmin.(string),
			"mActive":         "mActive",
		}))
	/*common.OutPutHtml(c, views.About(map[string]string{
		"siteName":        Config.Site_name,
		"siteDescription": Config.Site_description,
		"username":        username.(string),
	}))*/
	return
}
func (fc *FrontController) PingCtr(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func (fc *FrontController) ChargeCtr(c *gin.Context) {
	var Config = common.GetConfig()

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)

	c.HTML(200, "charge.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        Config.Site_name,
			"siteDescription": common.Config.Site_description,
			"username":        username.(string),
			"isAdmin":         isAdmin.(string),
		}))
}

func (fc *FrontController) ListTagCtr(c *gin.Context) {

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)
	c.Header("Cache-Control", "no-cache")
	var tags []model.Tag

	_ = common.NewDb.
		Order("total_nums desc").
		Find(&tags)
	//基础tag字体大小
	baseFontSize := 64

	var tagWithFontSize []vo.Tag_with_font_size_vo

	for _, tmpTag := range tags {
		if baseFontSize > 14 {
			baseFontSize = baseFontSize - 4
		}
		if baseFontSize < 14 {
			baseFontSize = 14
		}
		var tmpTagWithFontSize vo.Tag_with_font_size_vo
		tmpTagWithFontSize.TagId = tmpTag.ID
		tmpTagWithFontSize.Name = tmpTag.Name
		tmpTagWithFontSize.FontSize = baseFontSize
		tmpTagWithFontSize.TotalNums = tmpTag.TotalNums
		tagWithFontSize = append(tagWithFontSize, tmpTagWithFontSize)

		common.Sugar.Infof("font size now: %+v", baseFontSize)
	}

	c.HTML(200, "list-tag.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        common.Config.Site_name,
			"siteDescription": common.Config.Site_description,
			"tags":            tagWithFontSize,
			"username":        username.(string),
			"isAdmin":         isAdmin.(string),
			"tagActive":       "mActive",
		}))
	return
}

func (fc *FrontController) HomeCtr(c *gin.Context) {

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)
	if common.Config.PrivateMode == 1 {

		if username == "" {
			c.Redirect(301, "/admin/login")
			return
		}
	}
	c.Header("Cache-Control", "no-cache")

	cateId := c.DefaultQuery("cateId", "")

	tagId := c.DefaultQuery("tagId", "")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		common.Sugar.Error(err)
	}
	page -= 1
	if page < 0 {
		page = 0
	}

	prev_page := page
	if prev_page < 1 {
		prev_page = 1
	}
	next_page := page + 2

	rpp := 20
	offset := page * rpp
	var blogDataList []model.Article

	var result *gorm.DB

	if tagId != "" {
		tagIdNum, _ := strconv.Atoi(tagId)
		result = common.NewDb.
			Raw("select *  from  bk_article, json_each(bk_article.tag_ids) where json_each.value = ? LIMIT ? OFFSET ? ", tagIdNum, rpp, offset).
			Find(&blogDataList)
	} else {

		if cateId == "" {
			result = common.NewDb.
				Limit(rpp).
				Offset(offset).
				Order("aid desc").
				Find(&blogDataList)
		} else {
			realCateId, err := strconv.Atoi(cateId)
			if err != nil {
				common.Sugar.Error(err)
			}
			result = common.NewDb.
				Limit(rpp).
				Offset(offset).
				Order("aid desc").
				Where("cate_id = ?", realCateId).
				Find(&blogDataList)
		}
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		common.ShowMessage(c, &vo2.Msg{
			Msg: "文章不存在",
		})
		return
	}

	c.HTML(200, "index.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        common.Config.Site_name,
			"siteDescription": common.Config.Site_description,
			"username":        username.(string),
			"isAdmin":         isAdmin.(string),
			"blogDataList":    blogDataList,
			"getCateFromMap":  getFuncGetCateFromMap(),
			"categories":      category_service.GetCategories(),
			"prevPage":        fmt.Sprintf("%d", prev_page),
			"nextPage":        fmt.Sprintf("%d", next_page),
			"tagId":           tagId,
			"cateId":          cateId,
			"homeActive":      "mActive",
			"SubCutContent":   common.SubCutContent,
		}))
	return
}

func getFuncGetCateFromMap() func(cateId int64) string {
	cateMap := category_service.GetCategoriesAsMap()

	getCateFromMap := func(cateId int64) string {
		if val, ok := cateMap[cateId]; ok {
			return val
		}
		return "默认"
	}
	return getCateFromMap
}

func ParseTitle(title sql.NullString, username interface{}) string {
	titleStr := ""
	titleRune := []rune(title.String)
	if username == "" {
		strLen := len(titleRune)
		cutLen := math.Floor(float64(strLen / 2))
		maxLen := strLen - int(cutLen)
		if maxLen > strLen {
			maxLen = strLen - 1
		}
		if maxLen <= 0 {
			maxLen = 1
		}
		if maxLen > 6 {
			maxLen = 6
		}
		titleStr = string(titleRune[0:maxLen]) + "********"
	} else {
		return title.String
	}
	return titleStr
}

func (fc *FrontController) SearchCtr(c *gin.Context) {

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)
	c.Header("Cache-Control", "no-cache")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		common.Sugar.Info(err)
	}
	page -= 1
	if page < 0 {
		page = 0
	}

	prev_page := page
	if prev_page < 1 {
		prev_page = 1
	}
	next_page := page + 2
	keyword := c.DefaultQuery("keyword", "")
	common.Sugar.Info(keyword)
	if len(keyword) <= 0 {
		common.ShowMessage(c, &vo2.Msg{
			Msg: "搜索关键字不能为空",
		})
		return
	}
	orig_keyword := keyword
	keyword = strings.Trim(keyword, "%20")
	keyword = strings.TrimSpace(keyword)
	keyword = strings.Replace(keyword, " ", "%", -1)
	keyword = strings.Replace(keyword, "%20", "%", -1)
	common.Sugar.Info(keyword)
	rpp := 20
	offset := page * rpp

	var blogDataList []model.Article

	result := common.NewDb.Where("title LIKE  ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%").
		Limit(rpp).
		Offset(offset).
		Order("aid desc").
		Find(&blogDataList)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		common.ShowMessage(c, &vo2.Msg{
			Msg: "文章不存在",
		})
		return
	}

	c.HTML(200, "search.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        common.Config.Site_name,
			"siteDescription": common.Config.Site_description,
			"blogDataList":    blogDataList,

			"getCateFromMap": getFuncGetCateFromMap(),
			"categories":     category_service.GetCategories(),
			"keyword":        orig_keyword,
			"username":       username.(string),
			"isAdmin":        isAdmin.(string),
			"prevPage":       fmt.Sprintf("%d", prev_page),
			"nextPage":       fmt.Sprintf("%d", next_page),
		}))
	return
}

func (fc *FrontController) ArticleHistoryCtr(c *gin.Context) {

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)
	c.Header("Cache-Control", "no-cache")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		common.Sugar.Info(err)
	}
	page -= 1
	if page < 0 {
		page = 0
	}

	prev_page := page
	if prev_page < 1 {
		prev_page = 1
	}
	next_page := page + 2
	aid := c.DefaultQuery("aid", "")
	common.Sugar.Info(aid)
	if len(aid) <= 0 {
		common.ShowMessage(c, &vo2.Msg{
			Msg: "搜索关键字不能为空",
		})
		return
	}

	rpp := 20
	offset := page * rpp

	var blogDataList []model.ArticleHistory

	result := common.NewDb.Where("aid = ?", aid).
		Limit(rpp).
		Offset(offset).
		Order("id desc").
		Find(&blogDataList)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		common.ShowMessage(c, &vo2.Msg{
			Msg: "文章不存在",
		})
		return
	}

	c.HTML(200, "article_history_list.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        common.Config.Site_name,
			"siteDescription": common.Config.Site_description,
			"blogDataList":    blogDataList,

			"getCateFromMap": getFuncGetCateFromMap(),
			"categories":     category_service.GetCategories(),
			"aid":            aid,
			"username":       username.(string),
			"isAdmin":        isAdmin.(string),
			"prevPage":       fmt.Sprintf("%d", prev_page),
			"nextPage":       fmt.Sprintf("%d", next_page),
		}))
	return
}

func (fc *FrontController) ViewAltCtr(c *gin.Context) {

	id := c.DefaultQuery("id", "0")
	c.Redirect(301, fmt.Sprintf("/view/%s", id))
}

func (fc *FrontController) CountViewCtr(c *gin.Context) {
	c.Header("Cache-Control", "no-cache")

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		common.Sugar.Error("ID不能为空")
		return
	}
	common.LogInfo("更新统计, 文章id: " + fmt.Sprintf("%d", id))
	var blogItem model.Article
	result := common.NewDb.First(&blogItem, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		common.ShowMessage(c, &vo2.Msg{
			Msg: "文章不存在",
		})
		return
	}

	blogItem.Views = blogItem.Views + 1
	common.NewDb.Where("aid = ?", blogItem.Aid).
		Save(blogItem)

	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 UTC")
	c.Header("Cache-Control", "no-cache, no-store, no-transform, must-revalidate, private, max-age=0")
	c.Header("Pragma", "no-cache")
	c.String(http.StatusOK, fmt.Sprintf("document.getElementById('vct').innerHTML=%d", blogItem.Views))
}

func (fc *FrontController) ViewArticleHistoryCtr(c *gin.Context) {

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)
	c.Header("Cache-Control", "no-cache")
	id := c.Param("id")

	var blogItem model.ArticleHistory

	result := common.NewDb.
		Where("id = ?", id).
		First(&blogItem)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		common.ShowMessage(c, &vo2.Msg{
			Msg: "文章不存在",
		})
		return
	}

	var tagIds []int64

	err := json.Unmarshal(blogItem.TagIds, &tagIds)
	if err != nil {
		common.Sugar.Info("解析标签失败")
	}

	c.HTML(200, "view_article_history.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"title":           blogItem.Title,
			"siteName":        common.Config.Site_name,
			"siteDescription": common.SubCutContent(blogItem.Content, 64),

			"getCateFromMap": getFuncGetCateFromMap(),
			"categories":     category_service.GetCategories(),
			"username":       username.(string),
			"isAdmin":        isAdmin.(string),

			"tags": tag_service.BatchGetTagName(tagIds),
			"out": map[string]string{
				"aid":         fmt.Sprintf("%d", blogItem.Aid),
				"cateName":    getFuncGetCateFromMap()(blogItem.CateId),
				"cateId":      fmt.Sprintf("%d", blogItem.CateId),
				"title":       blogItem.Title,
				"content":     blogItem.Content,
				"publishTime": blogItem.PublishTime,
				"updateTime":  blogItem.UpdateTime,
				"username":    username.(string),
			},
		}))
	return

}

func (fc *FrontController) ViewCtr(c *gin.Context) {

	_, username, isAdmin := UserPermissionCheckDefaultAllow(c)
	c.Header("Cache-Control", "no-cache")
	id := c.Param("id")
	if id == "" {
		id = "1"
	}
	var blogItem model.Article

	result := common.NewDb.First(&blogItem, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		jumpUrl := "/admin/login"
		if isAdmin != "" {
			jumpUrl = "/admin/addblog"
		}
		common.ShowMessage(c, &vo2.Msg{
			Msg:     "首页不存在，你需要先创建首页",
			JumpUrl: jumpUrl,
		})
		return
	}

	// 把上级的id找出来

	pidList := []int64{blogItem.Aid}
	if blogItem.PAid > 0 {
		pidList = getPidList(blogItem.PAid, pidList)
	}
	common.Sugar.Infof("pidList: %+v", pidList)
	var tmpNavItemList []vo.Nav_item
	if blogItem.Aid <= 1 {
		tmpNavItemList = getNavItemListForHome()
	} else {

		tmpNavItemList = getNavItemList(blogItem.PAid, blogItem.Aid, pidList)
	}

	common.Sugar.Infof("navItemList: %+v", tmpNavItemList)
	var tagIds []int64

	err := json.Unmarshal(blogItem.TagIds, &tagIds)
	if err != nil {
		common.Sugar.Info("解析标签失败")
	}

	c.HTML(200, "view.html",
		common.Pongo2ContextWithVersion(pongo2.Context{
			"title":           blogItem.Title,
			"siteName":        common.Config.Site_name,
			"siteDescription": common.SubCutContent(blogItem.Content, 64),
			"navItemList":     tmpNavItemList,
			"getCateFromMap":  getFuncGetCateFromMap(),
			"categories":      category_service.GetCategories(),
			"username":        username.(string),
			"isAdmin":         isAdmin.(string),

			"tags": tag_service.BatchGetTagName(tagIds),
			"out": map[string]interface{}{
				"aid":         blogItem.Aid,
				"cateName":    getFuncGetCateFromMap()(blogItem.CateId),
				"cateId":      fmt.Sprintf("%d", blogItem.CateId),
				"title":       blogItem.Title,
				"content":     blogItem.Content,
				"publishTime": blogItem.PublishTime,
				"updateTime":  blogItem.UpdateTime,
				"username":    username.(string),
			},
		}))
	return

}

func getPidList(aid int64, pidList []int64) []int64 {

	var blogItem model.Article

	pid := aid

	for pid > 0 {

		common.NewDb.First(&blogItem, pid)
		pid = blogItem.PAid

		if blogItem.Aid > 0 {
			pidList = append(pidList, blogItem.Aid)
		}
	}
	return pidList
}

func getNavItemList(pid int64, aid int64, pidList []int64) []vo.Nav_item {

	var outList []vo.Nav_item

	//上上级
	//上级
	var firstList []model.Article

	//取第一层级，取本层级

	common.NewDb.Where("p_aid = ? ", 0).
		Order("sort_id asc").
		Find(&firstList)
	if firstList != nil && len(firstList) > 0 {
		//循环查数据
		for _, v := range firstList {
			tmpNav := vo.Nav_item{
				Id:   v.Aid,
				Name: v.Title,
			}

			for _, vp := range pidList {
				//id 在 列表里面，就把他们都查出来
				if v.Aid == vp {
					tmpNav = processChildNav(tmpNav, pidList)
				}
			}

			outList = append(outList, tmpNav)
		}
	}

	return outList
}

func processChildNav(nav vo.Nav_item, pidList []int64) vo.Nav_item {

	childNavList := []vo.Nav_item{}
	//处理子节点

	childItem := []model.Article{}

	common.NewDb.Where("p_aid = ? ", nav.Id).
		Order("sort_id asc").
		Find(&childItem)
	if childItem != nil && len(childItem) > 0 {
		childNavList = []vo.Nav_item{}
		for _, v1 := range childItem {
			tmpNavItem := vo.Nav_item{
				Id:   v1.Aid,
				Name: v1.Title,
			}
			for _, vp := range pidList {
				//id 在 列表里面，就把他们都查出来
				if v1.Aid == vp {
					//	nextPidList := removeElementFromList(pidList, v1.Aid)
					tmpNavItem = processChildNav(tmpNavItem, pidList)

				}
			}
			childNavList = append(childNavList, tmpNavItem)
		}
		nav.Children = childNavList
	}
	return nav
}

func removeElementFromList(pidList []int64, aid int64) []int64 {
	var newPidList []int64
	for _, v := range pidList {
		if v != aid {
			newPidList = append(newPidList, v)
		}
	}
	return newPidList
}

func getNavItemListForHome() []vo.Nav_item {

	var sameLevelItem []model.Article
	var childItem []model.Article

	//取上一层级，取本层级

	common.NewDb.Where("p_aid = ? ", 0).
		Order("sort_id asc").
		Find(&sameLevelItem)

	//同级的
	var childItemList []vo.Nav_item
	var childItemList2 []vo.Nav_item

	if sameLevelItem != nil && len(sameLevelItem) > 0 {
		for _, v := range sameLevelItem {
			tmpNavItem := vo.Nav_item{
				Id:   v.Aid,
				Name: v.Title,
			}

			//取子级数据
			common.NewDb.Where("p_aid = ? ", v.Aid).
				Order("sort_id asc").
				Find(&childItem)

			if childItem != nil && len(childItem) > 0 {
				childItemList2 = []vo.Nav_item{}
				for _, v1 := range childItem {
					tmpNavItemLv3 := vo.Nav_item{
						Id:   v1.Aid,
						Name: v1.Title,
					}
					//取子级数据
					common.NewDb.Where("p_aid = ? ", v1.Aid).
						Order("sort_id asc").
						Find(&childItem)

					if childItem != nil && len(childItem) > 0 {
						childItemList3 := []vo.Nav_item{}
						for _, v1 := range childItem {
							tmpNavItemLv3 := vo.Nav_item{
								Id:   v1.Aid,
								Name: v1.Title,
							}

							childItemList3 = append(childItemList3, tmpNavItemLv3)
						}
						tmpNavItemLv3.Children = childItemList3
					}
					childItemList2 = append(childItemList2, tmpNavItemLv3)
				}
				tmpNavItem.Children = childItemList2
			}
			childItemList = append(childItemList, tmpNavItem)
		}

	}
	return childItemList
}

func (fc *FrontController) checkNeedCharge(c *gin.Context, blog vo.VBlogItem) bool {
	ttformat := "2006-01-02 15:04:05"
	t2, _ := time.Parse(ttformat, blog.Publish_time.String)

	weekBefore := time.Now().Add(-168 * time.Hour)

	if t2.Before(weekBefore) {

		c.Redirect(301, "/charge")
		return true
	}
	return false
}
func (fc *FrontController) checkNeedChargeButNotRedirect(c *gin.Context, tt sql.NullString) bool {
	ttformat := "2006-01-02 15:04:05"
	t2, _ := time.Parse(ttformat, tt.String)
	weekBefore := time.Now().Add(-168 * time.Hour)

	if t2.Before(weekBefore) {
		return true
	}
	return false
}

func (fc *FrontController) LogoutCtr(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Delete("isAdmin")
	err := session.Save()
	if err != nil {
		common.LogError(err)
	}
	c.Redirect(301, "/")
}

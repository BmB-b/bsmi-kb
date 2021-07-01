package main

import (
	. "github.com/cnmade/bsmi-kb/app/controller"
	"github.com/cnmade/bsmi-kb/app/controller/admincontroller"
	"github.com/cnmade/bsmi-kb/pkg/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"

	"github.com/cnmade/pongo2gin"
)

func main() {
	common.InitApp()
	r := gin.New()
	r.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: "views",
		ContentType: "text/html; charset=utf-8",
		AlwaysNoCache: true,
	})

	r.Static("/assets", "./public/assets")
	store := cookie.NewStore([]byte("gssecret"))
	r.Use(sessions.Sessions("mysession", store))
	fc := new(FrontController)
	r.GET("/", fc.HomeCtr)
	r.HEAD("/", fc.HomeCtr)
	r.GET("/list-tag", fc.ListTagCtr)
	r.GET("/demopongo", fc.DemoPongoCtr)
	r.GET("/about", fc.AboutCtr)
	r.GET("/view/:id", fc.ViewCtr)
	r.GET("/view.php", fc.ViewAltCtr)
	r.GET("/ping", fc.PingCtr)
	r.GET("/search", fc.SearchCtr)
	r.GET("/charge", fc.ChargeCtr)
	r.GET("/user/logout", fc.LogoutCtr)
	r.GET("/countview/:id", fc.CountViewCtr)

	admin := r.Group("/admin")
	{
		admin.GET("/", admincontroller.ListBlogCtr)
		admin.GET("/login", admincontroller.LoginCtr)
		admin.POST("/login-process", admincontroller.LoginProcessCtr)
		admin.GET("/logout", admincontroller.LogoutCtr)
		admin.GET("/addblog", admincontroller.AddBlogCtr)
		admin.POST("/save-blog-add", admincontroller.SaveBlogAddCtr)
		admin.GET("/listblog", admincontroller.ListBlogCtr)
		admin.GET("/export", admincontroller.ExportCtr)
		admin.GET("/deleteblog/:id", admincontroller.DeleteBlogCtr)
		admin.POST("/save-blog-edit", admincontroller.SaveBlogEditCtr)
		admin.GET("/editblog/:id", admincontroller.EditBlogCtr)


		admin.GET("/list-cate", admincontroller.ListCateCtr)
		admin.POST("/save-edit-cate", admincontroller.SaveEditCateCtr)
		admin.GET("/edit-cate/:id", admincontroller.EditCateCtr)
		admin.GET("/add-cate", admincontroller.AddCateCtr)
		admin.POST("/save-add-cate", admincontroller.SaveAddCateCtr)

		admin.GET("/list-tag", admincontroller.ListTagCtr)


		admin.GET("/files", admincontroller.Files)
		admin.POST("/fileupload", admincontroller.FileUpload)
	}


	// rss


	rss := new(RSS)
	r.GET("/rss.php", rss.Alter)
	r.GET("/rss", rss.Out)


	a := new(Api)
	api := r.Group("/api")
	{
		api.GET("/", a.Index)
		api.GET("/nav-all", a.NavAll)
		api.GET("/nav-load", a.NavLoad)
		api.GET("view/:id", a.View)
	}
	log.Info().Msg("Server listen on 127.0.0.1:8005")
	err := r.Run("127.0.0.1:8005")
	if err != nil {
		common.LogError(err)
	}
}

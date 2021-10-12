package common

import (
	"database/sql"
	"github.com/cnmade/bsmi-kb/pkg/common/vo"
	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/grokify/html-strip-tags-go"
	"github.com/kataras/hcaptcha"
	"github.com/naoina/toml"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

/**
 * Logging error
 */
func LogError(err error) {
	if err != nil {
		Sugar.Error(tracerr.Sprint(tracerr.Wrap(err)))
	}
}

/**
 * Logging info
 */
func LogInfo(msg string) {
	if msg != "" {
		Sugar.Info(msg)
	}
}

func LogInfoF(msg string, v interface{}) {
	if msg != "" {
		Sugar.Infof(msg, v)
	}
}

/**
 * close rows defer
 */
func CloseRowsDefer(rows *sql.Rows) {
	_ = rows.Close()
}

/*
* ShowMessage with template
 */
func ShowMessage(c *gin.Context, m *vo.Msg) {

	c.HTML(200, "message-traditional.html",
		Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        Config.Site_name,
			"siteDescription": Config.Site_description,
			"message":         m.Msg,
		}))
	return
}

func ShowUMessage(c *gin.Context, m *vo.Umsg) {

	c.HTML(200, "message-traditional.html",
		Pongo2ContextWithVersion(pongo2.Context{
			"siteName":        Config.Site_name,
			"siteDescription": Config.Site_description,
			"message":         m.Msg,
			"url":             m.Url,
		}))
	return
}

func GetMinutes() string {
	return time.Now().Format("200601021504")
}

func GetConfig() *vo.AppConfig {
	_cm := "GetConfig@pkg/common/common"
	//TODO load config from cmd line argument
	f, err := os.Open("./vol/config.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var config vo.AppConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		Sugar.Infof(_cm+" error: %+v", err)
	}
	return &config
}

var (
	Config        *vo.AppConfig
	NewDb         *gorm.DB
	PDB           *gorm.DB
	Logger, _     = zap.NewProduction()
	Sugar         *zap.SugaredLogger
	BsmiKbVersion string
	HCaptchClient *hcaptcha.Client
)

func InitApp() {
	Config = GetConfig()
	//	gin.SetMode(Config.SrvMode)
	gin.SetMode(gin.DebugMode)
	NewDb = GetPDB(Config)
	//PDB = GetPDB(Config)
	defer Logger.Sync()
	Sugar = Logger.Sugar()

	HCaptchClient = hcaptcha.New(Config.HCaptchaSecretKey)
}

func OutPutHtml(c *gin.Context, s string) {
	c.Header("Content-Type", "text/html;charset=UTF-8")
	c.String(200, "%s", s)
	return
}
func OutPutText(c *gin.Context, s string) {
	c.Header("Content-Type", "text/plain;charset=UTF-8")
	c.String(200, "%s", s)
	return
}

/**
 * 截取指定长度的字符串，中文
 */
func SubCutContent(content string, length int) string {
	if len(content) <= length {
		return content
	}

	content = strip.StripTags(content)
	content = strings.TrimSpace(content)
	content = strings.Replace(content, "<!DOCTYPE html>", "", 1)
	content = strings.Replace(content, "&nbsp;", "", 1)

	tmpContent := []rune(content)

	rawLen := len(tmpContent)

	if length > rawLen {
		return content
	}

	return string(tmpContent[0:length])
}

package email_service

import (
	"crypto/tls"
	"github.com/cnmade/bsmi-kb/pkg/common"
	"github.com/go-mail/mail"
)

func SendTwoAuth(ToEmail string, KeyB string) {
	d := mail.NewDialer(
		common.Config.EmailConfig.SmtpHost,
		common.Config.EmailConfig.SmtpPort,
		common.Config.EmailConfig.SmtpUser,
		common.Config.EmailConfig.SmtpPassword,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := mail.NewMessage()
	m.SetHeader("From", "postmaster@sandbox9c08d31eb85441448981f19fec977ee3.mailgun.org")
	m.SetHeader("To", ToEmail)
	m.SetHeader("Subject", "两步登陆验证")
	m.SetBody("text/html", "您的两步登录验证码为：" + KeyB )
	if err := d.DialAndSend(m); err != nil {
		common.Sugar.Error(err)
	}
}

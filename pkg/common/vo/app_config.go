package vo

type AppConfig struct {
	Dbdsn            string
	Admin_user       string
	Admin_password   string
	AdminEmail 		 string
	Site_name        string
	Site_description string
	Site_url         string
	SrvMode          string
	ObjectStorageType uint //1 本地存储 ./vol/oss/  2 aws s3 云存储
	CaptchaEnabled uint
	HCaptchaSiteKey	string
	HCaptchaSecretKey string
	ObjectStorage    struct {
		Aws_access_key_id     string
		Aws_secret_access_key string
		Aws_region            string
		Aws_bucket            string
		Cdn_url               string
	}
	EmailConfig struct {
		SmtpHost string
		SmtpPort int
		SmtpUser string
		SmtpPassword string
	}
}


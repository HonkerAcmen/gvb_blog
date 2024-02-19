package config

type Config struct {
	Mysql      Mysql        `yaml:"mysql" json:"mysql"`
	Logger     Logger       `yaml:"logger" json:"logger"`
	System     System       `yaml:"system" json:"system"`
	SiteInfo   ConfigSite   `yaml:"site_info" json:"site_info"`
	UploadInfo ConfigUpload `ymal:"upload" json:"upload"`
}

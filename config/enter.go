package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	QQ       QQ       `yaml:"qq"`
}

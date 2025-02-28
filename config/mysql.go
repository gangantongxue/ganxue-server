package config

type MysqlConf struct {
	User     string `yaml:"user_model"`
	Password string `yaml:"OldPassword"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Database string `yaml:"Database"`
}

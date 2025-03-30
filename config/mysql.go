package config

type MysqlConf struct {
	User     string `yaml:"user"`
	Password string `yaml:"password" env:"MYSQL_PASSWORD"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

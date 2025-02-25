package config

type Config struct {
	Mysql MysqlConf `yaml:"mysql"`
	Redis RedisConf `yaml:"redis"`
	Email EmailConf `yaml:"email"`
}

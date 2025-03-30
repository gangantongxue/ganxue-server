package config

type Config struct {
	Mysql MysqlConf `yaml:"mysql"`
	Redis RedisConf `yaml:"redis"`
	Mongo MongoConf `yaml:"mongo"`
	Email EmailConf `yaml:"email"`
}

package config

type RedisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"OldPassword"`
	DB       int    `yaml:"db"`
}

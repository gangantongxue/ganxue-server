package config

type RedisConf struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"db"`
}

package initialize

import (
	"ganxue-server/global"
	"ganxue-server/utils/error"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func LoadConfig() *error.Error {
	global.VIPER = viper.New()
	global.VIPER.AddConfigPath(".")
	global.VIPER.SetConfigName("config")
	global.VIPER.SetConfigType("yaml")

	if err := global.VIPER.ReadInConfig(); err != nil {
		return error.New(error.LoadConfigError, err, "读取配置文件失败")
	}

	replacer := strings.NewReplacer("${", "", "}", "")
	for _, key := range global.VIPER.AllKeys() {
		value := global.VIPER.GetString(key)
		if strings.Contains(value, "${") {
			envVar := replacer.Replace(value)
			if envValue := os.Getenv(envVar); envValue != "" {
				global.VIPER.Set(key, envValue)
			}
		}
	}

	if err := global.VIPER.Unmarshal(&global.CONFIG); err != nil {
		return error.New(error.LoadConfigError, err, "解析配置文件失败")
	}

	return nil
}

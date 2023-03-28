package config

import (
	"github.com/spf13/viper"
	"os"
)

var Vipper *viper.Viper

func init() {
	loadConfig()
}

func loadConfig() {
	vp := viper.New()
	vp.AddConfigPath("/app/config")
	//vp.AddConfigPath(getConfigDirPath()) //相对main.go 或 test文件的路径
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Vipper = vp
}

func getConfigDirPath() string {
	configDir := "config"
	// 读取配置文件, 解决跑测试的时候找不到配置文件的问题，最多往上找5层目录
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(configDir); err == nil {
			return configDir
		} else {
			configDir = "../" + configDir
		}
	}

	return configDir
}
